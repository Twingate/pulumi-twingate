using System;
using System.Collections.Generic;
using Pulumi;
using Pulumi.Tls;
using Pulumi.Tls.Inputs;
using Twingate.Twingate;
using Twingate.Twingate.Inputs;
using Twingate.Twingate.Outputs;

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

    // Create a second Twingate service account
    var serviceAccount2 = new TwingateServiceAccount("ci_cd_account_cs_2", new TwingateServiceAccountArgs
    {
        Name = "CI CD Service CS 2",
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

    // Create a second Twingate group
    var tgGroup2 = new TwingateGroup("twingateGroup2", new TwingateGroupArgs
    {
        Name = "demo group CS 2",
    });

    // Create a Twingate Resource and configure resource permission
    var twingateResource = new TwingateResource("twingate_home_page_cs", new TwingateResourceArgs
    {
        Name = "Twingate Home Page CS test updated",
        Address = "www.twingate.com",
        RemoteNetworkId = remoteNetwork.Id,
        AccessGroups = new[]
        {
            new TwingateResourceAccessGroupArgs
            {
                GroupId = tgGroup.Id,
            },
            new TwingateResourceAccessGroupArgs
            {
                GroupId = tgGroup2.Id,
            },
        },
        AccessServices = new[]
        {
            new TwingateResourceAccessServiceArgs
            {
                ServiceAccountId = serviceAccount.Id,
            },
            new TwingateResourceAccessServiceArgs
            {
                ServiceAccountId = serviceAccount2.Id,
            },
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

    // Example: Create a Resource with JIT (Just-In-Time) Access Policies
    // Demonstrates group-specific access policies with different configurations:
    // - Group 1: AUTO_LOCK with automatic approval (7 day duration)
    // - Group 2: ACCESS_REQUEST with manual approval (2 hour duration)
    var jitResource = new TwingateResource("jit_resource_cs", new TwingateResourceArgs
    {
        Name = "JIT Access Resource CS",
        Address = "internal-app.example.com",
        RemoteNetworkId = remoteNetwork.Id,
        AccessPolicies = new[]
        {
            new TwingateResourceAccessPolicyArgs
            {
                Mode = "AUTO_LOCK",
                Duration = "7d",
                ApprovalMode = "MANUAL",
            },
        },
        AccessGroups = new[]
        {
            new TwingateResourceAccessGroupArgs
            {
                GroupId = tgGroup.Id,
                // Auto-lock: Access automatically expires after duration
                AccessPolicies = new[]
                {
                    new TwingateResourceAccessGroupAccessPolicyArgs
                    {
                        Mode = "AUTO_LOCK",
                        ApprovalMode = "AUTOMATIC",
                        Duration = "7d",
                    },
                },
            },
            new TwingateResourceAccessGroupArgs
            {
                GroupId = tgGroup2.Id,
                // Access request: Requires manual approval before granting access
                AccessPolicies = new[]
                {
                    new TwingateResourceAccessGroupAccessPolicyArgs
                    {
                        Mode = "ACCESS_REQUEST",
                        ApprovalMode = "MANUAL",
                        Duration = "2h",
                    },
                },
            },
        },
        Protocols = new TwingateResourceProtocolsArgs
        {
            Tcp = new TwingateResourceProtocolsTcpArgs
            {
                Policy = "RESTRICTED",
                Ports = { "443", "8080" },
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

    // Create a Twingate DNS Filtering Profile
    var exampleProfile = new TwingateDNSFilteringProfile("exampleProfile", new TwingateDNSFilteringProfileArgs
    {
        Name = "CS Pulumi DNS Filtering Profile",
        Priority = 2,
        FallbackMethod = "AUTO",
        Groups =
            {
                tgGroup.Id
            },
        AllowedDomains = new TwingateDNSFilteringProfileAllowedDomainsArgs
        {
            IsAuthoritative = false,
            Domains =
                {
                    "twingate.com",
                    "zoom.us"
                }
        },
        DeniedDomains = new TwingateDNSFilteringProfileDeniedDomainsArgs
        {
            IsAuthoritative = true,
            Domains =
                {
                    "evil.example"
                }
        },
        ContentCategories = new TwingateDNSFilteringProfileContentCategoriesArgs
        {
            BlockAdultContent = true
        },
        SecurityCategories = new TwingateDNSFilteringProfileSecurityCategoriesArgs
        {
            BlockDnsRebinding = false,
            BlockNewlyRegisteredDomains = false
        },
        PrivacyCategories = new TwingateDNSFilteringProfilePrivacyCategoriesArgs
        {
            BlockDisguisedTrackers = true
        }
    });

    // -----------------------------------------------------------------------
    // Twingate Gateway, CAs, and SSH Resource
    // -----------------------------------------------------------------------
    // Generate an RSA CA key + self-signed X.509 certificate at deploy time using
    // the Pulumi tls provider. For prod, replace with your own PKI material.
    var x509Key = new PrivateKey("x509_key_cs", new PrivateKeyArgs
    {
        Algorithm = "RSA",
        RsaBits = 4096,
    });
    var x509Cert = new SelfSignedCert("x509_cert_cs", new SelfSignedCertArgs
    {
        PrivateKeyPem = x509Key.PrivateKeyPem,
        IsCaCertificate = true,
        ValidityPeriodHours = 8760, // 1 year
        AllowedUses =
        {
            "key_encipherment",
            "digital_signature",
            "cert_signing",
            "crl_signing",
        },
        Subject = new SelfSignedCertSubjectArgs
        {
            CommonName = "pulumi-twingate-tls-ca-cs",
            Organization = "Twingate Example",
        },
    });

    var tlsCa = new TwingateX509CertificateAuthority("tls_ca_cs", new TwingateX509CertificateAuthorityArgs
    {
        Name = "Pulumi TLS CA CS",
        Certificate = x509Cert.CertPem,
    });

    // Generate an ed25519 SSH CA key — its PublicKeyOpenssh is the OpenSSH-format key.
    var sshKey = new PrivateKey("ssh_key_cs", new PrivateKeyArgs
    {
        Algorithm = "ED25519",
    });
    var sshCa = new TwingateSSHCertificateAuthority("ssh_ca_cs", new TwingateSSHCertificateAuthorityArgs
    {
        Name = "Pulumi SSH CA CS",
        PublicKey = sshKey.PublicKeyOpenssh,
    });

    // Gateway sits in the remote network and brokers traffic for downstream resources.
    var gateway = new TwingateGateway("gateway_cs", new TwingateGatewayArgs
    {
        Address = "10.0.0.1:8443",
        RemoteNetworkId = remoteNetwork.Id,
        X509CaId = tlsCa.Id,
        SshCaId = sshCa.Id,
    });

    // SSH Resource reachable through the Gateway.
    var sshResource = new TwingateSSHResource("ssh_resource_cs", new TwingateSSHResourceArgs
    {
        Name = "Bastion SSH CS",
        Address = "bastion.internal.example.com",
        RemoteNetworkId = remoteNetwork.Id,
        GatewayId = gateway.Id,
        Username = "ubuntu",
        AccessGroups = new[]
        {
            new TwingateSSHResourceAccessGroupArgs
            {
                GroupId = tgGroup.Id,
            },
        },
        Protocols = new TwingateSSHResourceProtocolsArgs
        {
            Tcp = new TwingateSSHResourceProtocolsTcpArgs
            {
                Policy = "RESTRICTED",
                Ports = { "22" },
            },
        },
    });

    // Rendered gateway config (e.g. for the gateway runtime to consume).
    var gatewayConfig = new TwingateGatewayConfig("gateway_config_cs", new TwingateGatewayConfigArgs
    {
        Port = 8443,
        MetricsPort = 9090,
        Ssh = new TwingateGatewayConfigSshArgs
        {
            Gateway = new TwingateGatewayConfigSshGatewayArgs
            {
                Username = "gateway",
                KeyType = "ed25519",
                UserCertTtl = "5m",
                HostCertTtl = "24h",
            },
            Ca = new TwingateGatewayConfigSshCaArgs
            {
                PrivateKeyFile = "/etc/gateway/ssh-ca.key",
            },
            Resources =
            {
                new TwingateGatewayConfigSshResourceArgs
                {
                    Name = sshResource.Name,
                    Address = sshResource.Address,
                    Username = "ubuntu",
                },
            },
        },
        Tls = new TwingateGatewayConfigTlsArgs
        {
            CertificateFile = "/etc/gateway/tls.crt",
            PrivateKeyFile = "/etc/gateway/tls.key",
        },
    });

    // -----------------------------------------------------------------------
    // Twingate User
    // -----------------------------------------------------------------------
    var exampleUser = new TwingateUser("example_user_cs", new TwingateUserArgs
    {
        Email = "example.user.cs@example.com",
        FirstName = "Example",
        LastName = "User",
        Role = "MEMBER",
        IsActive = true,
        SendInvite = false,
    });

    // -----------------------------------------------------------------------
    // Connector tokens — bake into a deployment (EC2 user_data, Helm, ...)
    // -----------------------------------------------------------------------
    var connectorTokens = new TwingateConnectorTokens("connector_tokens_cs", new TwingateConnectorTokensArgs
    {
        ConnectorId = tggcpConnector.Id,
    });

    // -----------------------------------------------------------------------
    // Look up a Security Policy by name and apply it to a group access block.
    // -----------------------------------------------------------------------
    var defaultPolicy = GetTwingateSecurityPolicy.Invoke(new GetTwingateSecurityPolicyInvokeArgs
    {
        Name = "Default Policy",
    });

    var policyResource = new TwingateResource("policy_resource_cs", new TwingateResourceArgs
    {
        Name = "Policy-bound Resource CS",
        Address = "policy-app.example.com",
        RemoteNetworkId = remoteNetwork.Id,
        AccessGroups = new[]
        {
            new TwingateResourceAccessGroupArgs
            {
                GroupId = tgGroup.Id,
                SecurityPolicyId = defaultPolicy.Apply(p => p.Id),
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["gateway_id"] = gateway.Id,
        ["gateway_config_content"] = gatewayConfig.Content,
        ["connector_access_token"] = Output.CreateSecret(connectorTokens.AccessToken),
        ["connector_refresh_token"] = Output.CreateSecret(connectorTokens.RefreshToken),
    };
});
