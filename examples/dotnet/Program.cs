using System;
using Pulumi;
using Pulumi.Twingate;
using Pulumi.Twingate.Inputs;
using Pulumi.Twingate.Outputs;

await Deployment.RunAsync(() =>
{
    // Create a Twingate remote network
    var remoteNetwork = new TwingateRemoteNetwork("test_network_cs", new TwingateRemoteNetworkArgs
    {
        Name = "Office CS",
    });

    // Create a Twingate service account
    var serviceAccount = new TwingateServiceAccount("ci_cd_account_cs", new TwingateServiceAccountArgs
    {
        Name = "CI CD Service CS",
    });

    // Create a Twingate service key
    var serviceAccountKey = new TwingateServiceAccountKey("ci_cd_key_cs", new TwingateServiceAccountKeyArgs
    {
        Name = "CI CD Key CS",
        ServiceAccountId = serviceAccount.Id,
    });

    // Create a Twingate connector
    var tggcpConnector = new TwingateConnector("twingateConnectorCS", new TwingateConnectorArgs
    {
        RemoteNetworkId = remoteNetwork.Id,
    });

    // Create a Twingate group
    var tgGroup = new TwingateGroup("twingateGroup", new TwingateGroupArgs
    {
        Name = "demo group CS",
    });

    // Create a Twingate Resource and configure resource permission
    var twingateResource = new TwingateResource("twingate_home_page_cs", new TwingateResourceArgs
    {
        Name = "Twingate Home Page CS",
        Address = "www.twingate.com",
        RemoteNetworkId = remoteNetwork.Id,
        AccessGroups = new TwingateResourceAccessGroupArgs
        {
            GroupId = tgGroup.Id,
        },
        AccessServices = new TwingateResourceAccessServiceArgs
        {
            ServiceAccountId = serviceAccount.Id,
        },
        Protocols = new TwingateResourceProtocolsArgs
        {
            AllowIcmp = true,
            Tcp = new TwingateResourceProtocolsTcpArgs
            {
                Policy = "RESTRICTED",
                Ports = { "80" },
            },
            Udp = new TwingateResourceProtocolsUdpArgs
            {
                Policy = "ALLOW_ALL",
            },
        },
    });

    // Specify the additional criteria to filter the Twingate groups
    var groupsArgs = new GetTwingateGroupsArgs
    {
        NameContains = "Everyone",
        IsActive = true,
    };

    // Invoke the GetTwingateGroups function asynchronously with the specified criteria
    var groupsResultTask = GetTwingateGroups.InvokeAsync(groupsArgs);

    // Wait for the task to complete and get the result synchronously
    var groupsResult = groupsResultTask.Result;

    // Access the properties of the groups result
    foreach (var group in groupsResult.Groups)
    {
        Console.WriteLine($"Group ID: {group.Id}");
        Console.WriteLine($"Group Name: {group.Name}");
        Console.WriteLine($"Group Active: {group.IsActive}");
        Console.WriteLine($"Group Security Policy ID: {group.SecurityPolicyId}");
        Console.WriteLine($"Group Type: {group.Type}");
        Console.WriteLine();
    }

    // Specify the additional criteria to filter the Twingate connectors
    var connectorArgs = new GetTwingateConnectorsArgs
    {
        NameContains = "t",
    };

    // Invoke the GetTwingateGroups function asynchronously with the specified criteria
    var connectorsResultTask = GetTwingateConnectors.InvokeAsync(connectorArgs);

    // Wait for the task to complete and get the result synchronously
    var connectorResult = connectorsResultTask.Result;

    // Access the properties of the connector result
    foreach (var connector in connectorResult.Connectors)
    {
        Console.WriteLine($"Connector Name: {connector.Name}");
        Console.WriteLine($"Connector Remote Network ID: {connector.RemoteNetworkId}");
        Console.WriteLine($"Connector Status Updates: {connector.StatusUpdatesEnabled}");
        Console.WriteLine();
    }
});
