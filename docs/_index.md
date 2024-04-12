---
title: Twingate
meta_desc: Provides an overview of the Twingate Provider for Pulumi.
layout: package
---

The Twingate provider for Pulumi can be used to provision any of the cloud resources available in [Twingate](https://www.twingate.com/).
The Twingate provider must be configured with credentials to deploy and update resources in Twingate.

## Example

{{< chooser language "typescript,python" >}}
{{% choosable language typescript %}}

```typescript
import * as tg from "@twingate/pulumi-twingate"
import * as pulumi from "@pulumi/pulumi"

// Create a Twingate remote network
const remoteNetwork = new tg.TwingateRemoteNetwork("test-network", { name: "Office" })

// Create a Twingate service account
const serviceAccount = new tg.TwingateServiceAccount("ci_cd_account", { name: "CI CD Service" })

// Create a Twingate service account key
const serviceAccountKey = new tg.TwingateServiceAccountKey("ci_cd_key", { name: "CI CD Key", serviceAccountId: serviceAccount.id })

// Create a Twingate connecntor
const tggcpConnector = new tg.TwingateConnector("twingateConnector", { remoteNetworkId: remoteNetwork.id });

// Create a Twingate group
const tggroup = new tg.TwingateGroup("twingateGroup", {
    name: "demo group",
});

// To see serviceAccountKeyOut, execute command `pulumi stack output --show-secrets`
export const serviceAccountKeyOut = pulumi.interpolate`${serviceAccountKey.token}`;

// Get group id by name
function getGroupId(groupName: string) {
    const groups: any = tg.getTwingateGroupsOutput({ name: groupName })?.groups ?? []
    return groups[0].id
}

// Create a Twingate Resource and configure resource permission
new tg.TwingateResource("twingate_home_page", {
    name: "Twingate Home Page",
    address: "www.twingate.com",
    remoteNetworkId: remoteNetwork.id,
    accessGroups: [
        {
            groupId: getGroupId("Everyone"),
        }
    ],
    accessServices: [
        {
            serviceAccountId: serviceAccount.id,
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

// Get Twingate connector and filter results
const result: Promise<tg.GetTwingateConnectorsResult> = tg.getTwingateConnectors({ nameContains: "twingate" });

result.then((value) => {
    const connectors = value.connectors;
    console.log(connectors);
});

```

{{% /choosable %}}
{{% choosable language python %}}

```python
import pulumi
import pulumi_twingate as tg

# Create a Twingate remote network
remote_network = tg.TwingateRemoteNetwork("test_network", name="Office")

# Create a Twingate service account
service_account = tg.TwingateServiceAccount("ci_cd_account", name="CI CD Service")

# Create a Twingate service key
service_account_key = tg.TwingateServiceAccountKey("ci_cd_key", name="CI CD Key", service_account_id=service_account.id)

# Create a Twingate connector
tggcp_connector = tg.TwingateConnector("twingateConnector", remote_network_id=remote_network.id)

# Create a Twingate group
tg_group = tg.TwingateGroup("twingateGroup", name="demo group")


# To see service_account_key, execute command `pulumi stack output --show-secrets`
pulumi.export("service_account_key", service_account_key.token)


# Get group id by name
def get_group_id(group_name):
    group = tg.get_twingate_groups_output(name=group_name).groups[0]
    return group.id


# Create a Twingate Resource and configure resource permission
twingate_resource = tg.TwingateResource(
    "twingate_home_page",
    name="Twingate Home Page",
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
```

{{% /choosable %}}
{{< /chooser >}}
