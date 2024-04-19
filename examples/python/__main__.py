import pulumi
import pulumi_twingate as tg

# Create a Twingate remote network
remote_network = tg.TwingateRemoteNetwork("test_network_py", name="Office PY")

# Create a Twingate service account
service_account = tg.TwingateServiceAccount("ci_cd_account_py", name="CI CD Service PY")

# Create a Twingate service key
service_account_key = tg.TwingateServiceAccountKey("ci_cd_key_py", name="CI CD Key PY", service_account_id=service_account.id)

# Create a Twingate connector
tggcp_connector = tg.TwingateConnector("twingateConnectorPY", remote_network_id=remote_network.id)

# Create a Twingate group
tg_group = tg.TwingateGroup("twingateGroup", name="demo group PY")


# To see service_account_key, execute command `pulumi stack output --show-secrets`
pulumi.export("service_account_key", service_account_key.token)


# Get group id by name
def get_group_id(group_name):
    group = tg.get_twingate_groups_output(name=group_name).groups[0]
    return group.id


# Create a Twingate Resource and configure resource permission
twingate_resource = tg.TwingateResource(
    "twingate_home_page_py",
    name="Twingate Home Page PY",
    address="www.twingate.com",
    remote_network_id=remote_network.id,
    access_groups=[
        {
            "groupId": get_group_id("Everyone"),
        }
    ],
    access_services=[
        {
            "service_account_ids": service_account.id,
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
