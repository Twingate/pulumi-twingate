import pulumi
import pulumi_tls as tls
import pulumi_twingate as tg

# Create a Twingate remote network
remote_network = tg.TwingateRemoteNetwork("test_network_py", name="Office PY")

# Create a Twingate service account
service_account = tg.TwingateServiceAccount("ci_cd_account_py", name="CI CD Service PY")
service_account2 = tg.TwingateServiceAccount("ci_cd_account_py_2", name="CI CD Service PY 2")

# Create a Twingate service key
service_account_key = tg.TwingateServiceAccountKey("ci_cd_key_py", name="CI CD Key PY", service_account_id=service_account.id)

# Create a Twingate connector
tggcp_connector = tg.TwingateConnector("twingateConnectorPY", remote_network_id=remote_network.id)

# Create Twingate groups
tg_group = tg.TwingateGroup("twingateGroup", name="demo group PY")
tg_group2 = tg.TwingateGroup("twingateGroup2", name="demo group PY 2")


# To see service_account_key, execute command `pulumi stack output --show-secrets`
pulumi.export("service_account_key", service_account_key.token)


# Get group id by name
def get_group_id(group_name):
    group = tg.get_twingate_groups_output(name=group_name).groups[0]
    return group.id


# Create a Twingate Resource and configure resource permission
twingate_resource = tg.TwingateResource(
    "twingate_home_page_py",
    name="Twingate Home Page PY test updated",
    address="www.twingate.com",
    remote_network_id=remote_network.id,
    access_groups=[
        {
            "groupId": tg_group.id,
        },
        {
            "groupId": tg_group2.id,
        }
    ],
    access_services=[
        {
            "serviceAccountId": service_account.id,
        },
        {
            "serviceAccountId": service_account2.id,
        }
    ],
    protocols={
        "allowIcmp": True,
        "tcp": {
            "policy": "RESTRICTED",
            "ports": ["80"],
        },
        "udp": {
            "policy": "ALLOW_ALL",
        },
    }
)

# Example: Create a Resource with JIT (Just-In-Time) Access Policies
# Demonstrates group-specific access policies with different configurations:
# - Group 1: AUTO_LOCK with automatic approval (7 day duration)
# - Group 2: ACCESS_REQUEST with manual approval (2 hour duration)
jit_resource = tg.TwingateResource(
    "jit_resource_py",
    name="JIT Access Resource PY",
    address="internal-app.example.com",
    remote_network_id=remote_network.id,
    access_policies=[
        tg.TwingateResourceAccessPolicyArgs(
            mode="AUTO_LOCK",
            duration="7d",
            approval_mode="MANUAL",
        )
    ],
    access_groups=[
        tg.TwingateResourceAccessGroupArgs(
            group_id=tg_group.id,
            # Auto-lock: Access automatically expires after duration
            access_policies=[
                tg.TwingateResourceAccessGroupAccessPolicyArgs(
                    mode="AUTO_LOCK",
                    approval_mode="AUTOMATIC",
                    duration="7d",
                )
            ],
        ),
        tg.TwingateResourceAccessGroupArgs(
            group_id=tg_group2.id,
            # Access request: Requires manual approval before granting access
            access_policies=[
                tg.TwingateResourceAccessGroupAccessPolicyArgs(
                    mode="ACCESS_REQUEST",
                    approval_mode="MANUAL",
                    duration="2h",
                )
            ],
        )
    ],
    protocols={
        "tcp": {
            "policy": "RESTRICTED",
            "ports": ["443", "8080"],
        },
    }
)

# Get Twingate connector and filter results
result = tg.get_twingate_connectors(name_contains="t")

# Access the connectors property of the returned result
connectors = result.connectors

# Print the list of connectors
for connector in connectors:
    print(f"Connector ID: {connector.remote_network_id}, Name: {connector.name}")

# Create a Twingate DNS Filtering Profile
example_profile = tg.TwingateDNSFilteringProfile("exampleProfilePY",
    name="PY Pulumi DNS Filtering Profile",
    priority=2,
    fallback_method="AUTO",
    groups=[
        tg_group.id
    ],
    allowed_domains=tg.TwingateDNSFilteringProfileAllowedDomainsArgs(
        is_authoritative=False,
        domains=[
            "twingate.com",
            "zoom.us"
        ]
    ),
    denied_domains=tg.TwingateDNSFilteringProfileDeniedDomainsArgs(
        is_authoritative=True,
        domains=[
            "evil.example"
        ]
    ),
    content_categories=tg.TwingateDNSFilteringProfileContentCategoriesArgs(
        block_adult_content=True
    ),
    security_categories=tg.TwingateDNSFilteringProfileSecurityCategoriesArgs(
        block_dns_rebinding=False,
        block_newly_registered_domains=False
    ),
    privacy_categories=tg.TwingateDNSFilteringProfilePrivacyCategoriesArgs(
        block_disguised_trackers=True
    )
)

# ---------------------------------------------------------------------------
# Twingate Gateway, CAs, and SSH Resource
# ---------------------------------------------------------------------------
# Generate an RSA CA key + self-signed X.509 certificate at deploy time using
# the Pulumi tls provider. For prod, replace with your own PKI material.
x509_key = tls.PrivateKey("x509_key_py", algorithm="RSA", rsa_bits=4096)
x509_cert = tls.SelfSignedCert(
    "x509_cert_py",
    private_key_pem=x509_key.private_key_pem,
    is_ca_certificate=True,
    validity_period_hours=8760,  # 1 year
    allowed_uses=[
        "key_encipherment",
        "digital_signature",
        "cert_signing",
        "crl_signing",
    ],
    subject={
        "common_name": "pulumi-twingate-tls-ca-py",
        "organization": "Twingate Example",
    },
)

tls_ca = tg.TwingateX509CertificateAuthority(
    "tls_ca_py",
    name="Pulumi TLS CA PY",
    certificate=x509_cert.cert_pem,
)

# Generate an ed25519 SSH CA key — its public_key_openssh is the OpenSSH-format key.
ssh_key = tls.PrivateKey("ssh_key_py", algorithm="ED25519")
ssh_ca = tg.TwingateSSHCertificateAuthority(
    "ssh_ca_py",
    name="Pulumi SSH CA PY",
    public_key=ssh_key.public_key_openssh,
)

# Gateway sits in the remote network and brokers traffic for downstream resources.
gateway = tg.TwingateGateway(
    "gateway_py",
    address="10.0.0.1:8443",
    remote_network_id=remote_network.id,
    x509_ca_id=tls_ca.id,
    ssh_ca_id=ssh_ca.id,
)

# SSH Resource reachable through the Gateway.
ssh_resource = tg.TwingateSSHResource(
    "ssh_resource_py",
    name="Bastion SSH PY",
    address="bastion.internal.example.com",
    remote_network_id=remote_network.id,
    gateway_id=gateway.id,
    username="ubuntu",
    access_groups=[
        {
            "group_id": tg_group.id,
        },
    ],
    protocols={
        "tcp": {
            "policy": "RESTRICTED",
            "ports": ["22"],
        },
    },
)

# Rendered gateway config (e.g. for the gateway runtime to consume).
gateway_config = tg.TwingateGatewayConfig(
    "gateway_config_py",
    port=8443,
    metrics_port=9090,
    ssh={
        "gateway": {
            "username": "gateway",
            "key_type": "ed25519",
            "user_cert_ttl": "5m",
            "host_cert_ttl": "24h",
        },
        "ca": {
            "private_key_file": "/etc/gateway/ssh-ca.key",
        },
        "resources": [
            {
                "name": ssh_resource.name,
                "address": ssh_resource.address,
                "username": "ubuntu",
            },
        ],
    },
    tls={
        "certificate_file": "/etc/gateway/tls.crt",
        "private_key_file": "/etc/gateway/tls.key",
    },
)

pulumi.export("gateway_id", gateway.id)
pulumi.export("gateway_config_content", gateway_config.content)

# ---------------------------------------------------------------------------
# Twingate User
# ---------------------------------------------------------------------------
example_user = tg.TwingateUser(
    "example_user_py",
    email="example.user.py@example.com",
    first_name="Example",
    last_name="User",
    role="MEMBER",
    is_active=True,
    send_invite=False,
)

# ---------------------------------------------------------------------------
# Connector tokens — bake into a deployment (EC2 user_data, Helm values, ...)
# ---------------------------------------------------------------------------
connector_tokens = tg.TwingateConnectorTokens(
    "connector_tokens_py",
    connector_id=tggcp_connector.id,
)
pulumi.export("connector_access_token", pulumi.Output.secret(connector_tokens.access_token))
pulumi.export("connector_refresh_token", pulumi.Output.secret(connector_tokens.refresh_token))

# ---------------------------------------------------------------------------
# Look up a Security Policy by name and apply it to a group access block.
# ---------------------------------------------------------------------------
default_policy = tg.get_twingate_security_policy(name="Default Policy")

policy_resource = tg.TwingateResource(
    "policy_resource_py",
    name="Policy-bound Resource PY",
    address="policy-app.example.com",
    remote_network_id=remote_network.id,
    access_groups=[
        {
            "group_id": tg_group.id,
            "security_policy_id": default_policy.id,
        }
    ],
)
