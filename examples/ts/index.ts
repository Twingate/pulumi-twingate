import * as tg from "@twingate/pulumi-twingate"
import * as pulumi from "@pulumi/pulumi"

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

// Example: Create a Resource with JIT (Just-In-Time) Access Policy
// Apply the policy directly to the access group for it to take effect
new tg.TwingateResource("jit_resource_js", {
    name: "JIT Access Resource JS",
    address: "internal-app.example.com",
    remoteNetworkId: remoteNetwork.id,
    accessGroups: [
        {
            groupId: tgGroup.id,
            // Access policy must be set on the group level when using accessGroups
            accessPolicies: [
                {
                    mode: "AUTO_LOCK",           // Automatically lock access after duration
                    approvalMode: "AUTOMATIC",    // No manual approval required
                    duration: "24h",              // Access granted for 24 hours
                }
            ],
        }
    ],
    protocols: {
        allowIcmp: true,
        tcp: {
            policy: "ALLOW_ALL",
        },
    }
})

// Example: Create a Resource with group-specific Access Policies
new tg.TwingateResource("group_policy_resource_js", {
    name: "Group-Specific Policy Resource JS",
    address: "sensitive-app.example.com",
    remoteNetworkId: remoteNetwork.id,
    accessGroups: [
        {
            groupId: tgGroup.id,
            // This group gets auto-lock access
            accessPolicies: [
                {
                    mode: "AUTO_LOCK",
                    approvalMode: "AUTOMATIC",
                    duration: "8h",
                }
            ],
        },
        {
            groupId: tgGroup2.id,
            // This group requires manual access requests
            accessPolicies: [
                {
                    mode: "ACCESS_REQUEST",
                    approvalMode: "MANUAL",      // Requires manual approval
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
const exampleProfile = new tg.TwingateDNSFilteringProfile("exampleProfile", {
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
