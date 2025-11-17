import pulumi
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

# Get Twingate connector and filter results
result = tg.get_twingate_connectors(name_contains="t")

# Access the connectors property of the returned result
connectors = result.connectors

# Print the list of connectors
for connector in connectors:
    print(f"Connector ID: {connector.remote_network_id}, Name: {connector.name}")

# Create a Twingate DNS Filtering Profile
example_profile = tg.TwingateDNSFilteringProfile("exampleProfile",
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
