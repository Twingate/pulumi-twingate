import * as tg from "@twingate/pulumi-twingate"
import * as pulumi from "@pulumi/pulumi"

// Create a Twingate remote network
const remoteNetwork = new tg.TwingateRemoteNetwork("test-network-js", { name: "Office JS" })

// Create a Twingate service account
const serviceAccount = new tg.TwingateServiceAccount("ci_cd_account_js", { name: "CI CD Service JS" })

// Create a Twingate service account key
const serviceAccountKey = new tg.TwingateServiceAccountKey("ci_cd_key_js", { name: "CI CD Key JS", serviceAccountId: serviceAccount.id })

// Create a Twingate connecntor
const tggcpConnector = new tg.TwingateConnector("twingateConnectorJS", { remoteNetworkId: remoteNetwork.id });

// Create a Twingate group
const tggroup = new tg.TwingateGroup("twingateGroup", {
    name: "demo group JS",
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
    name: "Twingate Home Page JS",
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
