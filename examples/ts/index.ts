import * as tg from "@twingate/pulumi-twingate"
import * as pulumi from "@pulumi/pulumi"
import * as tls from "@pulumi/tls"

// Create a Twingate remote network
const remoteNetwork = new tg.TwingateRemoteNetwork("test-network-js", { name: "Office JS" })

// Create a Twingate service account
const serviceAccount = new tg.TwingateServiceAccount("ci_cd_account_js", { name: "CI CD Service JS" })

// Create a second Twingate service account
const serviceAccount2 = new tg.TwingateServiceAccount("ci_cd_account_js_2", { name: "CI CD Service JS 2" })

// Create a Twingate service account key
const serviceAccountKey = new tg.TwingateServiceAccountKey("ci_cd_key_js", { name: "CI CD Key JS", serviceAccountId: serviceAccount.id })

// Create a Twingate connecntor
const tggcpConnector = new tg.TwingateConnector("twingateConnectorJS", { remoteNetworkId: remoteNetwork.id });

// Create a Twingate group
const tgGroup = new tg.TwingateGroup("twingateGroup", {
    name: "demo group JS",
});

// Create a second Twingate group
const tgGroup2 = new tg.TwingateGroup("twingateGroup2", {
    name: "demo group JS 2",
});

// To see serviceAccountKeyOut, execute command `pulumi stack output --show-secrets`
export const serviceAccountKeyOut = pulumi.interpolate`${serviceAccountKey.token}`;

// Get group id by name
function getGroupId(groupName: string) {
    const groups: any = tg.getTwingateGroupsOutput({ name: groupName })?.groups ?? []
    return groups[0].id
}

// Create a Twingate Resource and configure resource permission
new tg.TwingateResource("twingate_home_page_js", {
    name: "Twingate Home Page JS test updated",
    address: "www.twingate.com",
    remoteNetworkId: remoteNetwork.id,
    accessGroups: [
        {
            groupId: tgGroup.id,
        },
        {
            groupId: tgGroup2.id,
        }
    ],
    accessServices: [
        {
            serviceAccountId: serviceAccount.id,
        },
        {
            serviceAccountId: serviceAccount2.id,
        }
    ],
    protocols: {
        allowIcmp: true,
        tcp: {
            policy: "RESTRICTED",
            ports: ["80"],
        },
        udp: {
            policy: "ALLOW_ALL",
        },
    }
})

// Example: Create a Resource with JIT (Just-In-Time) Access Policies
// Demonstrates group-specific access policies with different configurations:
// - Group 1: AUTO_LOCK with automatic approval (7 day duration)
// - Group 2: ACCESS_REQUEST with manual approval (2 hour duration)
new tg.TwingateResource("jit_resource_js", {
    name: "JIT Access Resource JS",
    address: "internal-app.example.com",
    remoteNetworkId: remoteNetwork.id,
    accessPolicies: [
        {
            mode: "AUTO_LOCK",
            duration: "7d",
            approvalMode: "MANUAL",
        }
    ],
    accessGroups: [
        {
            groupId: tgGroup.id,
            // Auto-lock: Access automatically expires after duration
            accessPolicies: [
                {
                    mode: "AUTO_LOCK",
                    approvalMode: "AUTOMATIC",
                    duration: "7d",
                }
            ],
        },
        {
            groupId: tgGroup2.id,
            // Access request: Requires manual approval before granting access
            accessPolicies: [
                {
                    mode: "ACCESS_REQUEST",
                    approvalMode: "MANUAL",
                    duration: "2h",
                }
            ],
        }
    ],
    protocols: {
        tcp: {
            policy: "RESTRICTED",
            ports: ["443", "8080"],
        },
    }
})

// Get Twingate connector and filter results
const result: Promise<tg.GetTwingateConnectorsResult> = tg.getTwingateConnectors({ nameContains: "twingate" });

result.then((value) => {
    const connectors = value.connectors;
    console.log(connectors);
});

// Create a Twingate DNS Filtering Profile
const exampleProfile = new tg.TwingateDNSFilteringProfile("exampleProfileJS", {
    name: "JS Pulumi DNS Filtering Profile",
    priority: 2,
    fallbackMethod: "AUTO",
    groups: [tgGroup.id],
    allowedDomains: {
        isAuthoritative: false,
        domains: [
            "twingate.com",
            "zoom.us"
        ]
    },
    deniedDomains: {
        isAuthoritative: true,
        domains: [
            "evil.example"
        ]
    },
    contentCategories: {
        blockAdultContent: true
    },
    securityCategories: {
        blockDnsRebinding: false,
        blockNewlyRegisteredDomains: false
    },
    privacyCategories: {
        blockDisguisedTrackers: true
    }
} as tg.TwingateDNSFilteringProfileArgs);

// ---------------------------------------------------------------------------
// Twingate Gateway, CAs, and SSH Resource
// ---------------------------------------------------------------------------
// Generate an RSA CA key + self-signed X.509 certificate at deploy time using
// the Pulumi tls provider. For prod, replace with your own PKI material.
const x509Key = new tls.PrivateKey("x509_key_js", {
    algorithm: "RSA",
    rsaBits: 4096,
});
const x509Cert = new tls.SelfSignedCert("x509_cert_js", {
    privateKeyPem: x509Key.privateKeyPem,
    isCaCertificate: true,
    validityPeriodHours: 8760, // 1 year
    allowedUses: [
        "key_encipherment",
        "digital_signature",
        "cert_signing",
        "crl_signing",
    ],
    subject: {
        commonName: "pulumi-twingate-tls-ca-js",
        organization: "Twingate Example",
    },
});

const tlsCa = new tg.TwingateX509CertificateAuthority("tls_ca_js", {
    name: "Pulumi TLS CA JS",
    certificate: x509Cert.certPem,
});

// Generate an ed25519 SSH CA key — its publicKeyOpenssh is the OpenSSH-format key.
const sshKey = new tls.PrivateKey("ssh_key_js", { algorithm: "ED25519" });
const sshCa = new tg.TwingateSSHCertificateAuthority("ssh_ca_js", {
    name: "Pulumi SSH CA JS",
    publicKey: sshKey.publicKeyOpenssh,
});

// Gateway sits in the remote network and brokers traffic for downstream resources.
const gateway = new tg.TwingateGateway("gateway_js", {
    address: "10.0.0.1:8443",
    remoteNetworkId: remoteNetwork.id,
    x509CaId: tlsCa.id,
    sshCaId: sshCa.id,
});

// SSH Resource reachable through the Gateway.
const sshResource = new tg.TwingateSSHResource("ssh_resource_js", {
    name: "Bastion SSH JS",
    address: "bastion.internal.example.com",
    remoteNetworkId: remoteNetwork.id,
    gatewayId: gateway.id,
    username: "ubuntu",
    accessGroups: [
        {
            groupId: tgGroup.id,
        },
    ],
    protocols: {
        tcp: {
            policy: "RESTRICTED",
            ports: ["22"],
        },
    },
});

// Rendered gateway config (e.g. for the gateway runtime to consume).
const gatewayConfig = new tg.TwingateGatewayConfig("gateway_config_js", {
    port: 8443,
    metricsPort: 9090,
    ssh: {
        gateway: {
            username: "gateway",
            keyType: "ed25519",
            userCertTtl: "5m",
            hostCertTtl: "24h",
        },
        ca: {
            privateKeyFile: "/etc/gateway/ssh-ca.key",
        },
        resources: [
            {
                name: sshResource.name,
                address: sshResource.address,
                username: "ubuntu",
            },
        ],
    },
    tls: {
        certificateFile: "/etc/gateway/tls.crt",
        privateKeyFile: "/etc/gateway/tls.key",
    },
});

export const gatewayId = gateway.id;
export const gatewayConfigContent = gatewayConfig.content;

// ---------------------------------------------------------------------------
// Twingate User
// ---------------------------------------------------------------------------
const exampleUser = new tg.TwingateUser("example_user_js", {
    email: "example.user.js@example.com",
    firstName: "Example",
    lastName: "User",
    role: "MEMBER",
    isActive: true,
    sendInvite: false,
});

// ---------------------------------------------------------------------------
// Connector tokens — bake into a deployment (EC2 user_data, Helm values, ...)
// ---------------------------------------------------------------------------
const connectorTokens = new tg.TwingateConnectorTokens("connector_tokens_js", {
    connectorId: tggcpConnector.id,
});
export const connectorAccessToken = pulumi.secret(connectorTokens.accessToken);
export const connectorRefreshToken = pulumi.secret(connectorTokens.refreshToken);

// ---------------------------------------------------------------------------
// Look up a Security Policy by name and apply it to a group access block.
// ---------------------------------------------------------------------------
const defaultPolicy = tg.getTwingateSecurityPolicyOutput({ name: "Default Policy" });

new tg.TwingateResource("policy_resource_js", {
    name: "Policy-bound Resource JS",
    address: "policy-app.example.com",
    remoteNetworkId: remoteNetwork.id,
    accessGroups: [
        {
            groupId: tgGroup.id,
            securityPolicyId: defaultPolicy.apply(p => p.id ?? ""),
        },
    ],
});
