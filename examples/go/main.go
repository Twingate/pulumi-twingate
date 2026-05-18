package main

import (
	"github.com/Twingate/pulumi-twingate/sdk/v4/go/twingate"
	"github.com/pulumi/pulumi-tls/sdk/v5/go/tls"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Twingate remote network
		remoteNetwork, err := twingate.NewTwingateRemoteNetwork(ctx, "test_network_go", &twingate.TwingateRemoteNetworkArgs{
			Name: pulumi.StringPtr("Office Go"),
		})
		if err != nil {
			return err
		}

		// Create a Twingate service account
		serviceAccount, err := twingate.NewTwingateServiceAccount(ctx, "ci_cd_account_go", &twingate.TwingateServiceAccountArgs{
			Name: pulumi.StringPtr("CI CD Service Go"),
		})
		if err != nil {
			return err
		}

		// Create a second Twingate service account
		serviceAccount2, err := twingate.NewTwingateServiceAccount(ctx, "ci_cd_account_go_2", &twingate.TwingateServiceAccountArgs{
			Name: pulumi.StringPtr("CI CD Service Go 2"),
		})
		if err != nil {
			return err
		}

		// Create a Twingate service key
		serviceAccountKey, err := twingate.NewTwingateServiceAccountKey(ctx, "ci_cd_key_go", &twingate.TwingateServiceAccountKeyArgs{
			Name:             pulumi.StringPtr("CI CD key Go"),
			ServiceAccountId: serviceAccount.ID(),
		})
		if err != nil {
			return err
		}

		// To see service_account_key, execute command `pulumi stack output --show-secrets`
		ctx.Export("service_account_key", pulumi.ToSecret(serviceAccountKey.Token))

		// Create a Twingate connector
		connector, err := twingate.NewTwingateConnector(ctx, "test_connector_go", &twingate.TwingateConnectorArgs{
			RemoteNetworkId: remoteNetwork.ID(),
		})
		if err != nil {
			return err
		}

		ctx.Export("connector_id", connector.ID())

		// Create a Twingate group
		group, err := twingate.NewTwingateGroup(ctx, "twingate_group_go", &twingate.TwingateGroupArgs{
			Name: pulumi.StringPtr("Demo Group Go"),
		})
		if err != nil {
			return err
		}

		// Create a second Twingate group
		group2, err := twingate.NewTwingateGroup(ctx, "twingate_group_go_2", &twingate.TwingateGroupArgs{
			Name: pulumi.StringPtr("Demo Group Go 2"),
		})
		if err != nil {
			return err
		}

		// Create a Twingate Resource and configure resource permission
		resource, err := twingate.NewTwingateResource(ctx, "twingate_home_page_go", &twingate.TwingateResourceArgs{
			Name:            pulumi.StringPtr("Twingate Home Page Go Test"),
			Address:         pulumi.String("www.twingate.com"),
			RemoteNetworkId: remoteNetwork.ID(),
			AccessGroups: &twingate.TwingateResourceAccessGroupArray{
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group.ID(),
				},
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group2.ID(),
				},
			},
			AccessServices: &twingate.TwingateResourceAccessServiceArray{
				&twingate.TwingateResourceAccessServiceArgs{
					ServiceAccountId: serviceAccount.ID(),
				},
				&twingate.TwingateResourceAccessServiceArgs{
					ServiceAccountId: serviceAccount2.ID(),
				},
			},
			Protocols: &twingate.TwingateResourceProtocolsArgs{
				AllowIcmp: pulumi.BoolPtr(true),
				Tcp: &twingate.TwingateResourceProtocolsTcpArgs{
					Policy: pulumi.StringPtr("RESTRICTED"),
					Ports: pulumi.StringArray{
						pulumi.String("80"),
					},
				},
				Udp: &twingate.TwingateResourceProtocolsUdpArgs{
					Policy: pulumi.StringPtr("ALLOW_ALL"),
				},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("resource_id", resource.ID())

		// Example: Create a Resource with JIT (Just-In-Time) Access Policies
		// Demonstrates group-specific access policies with different configurations:
		// - Group 1: AUTO_LOCK with automatic approval (7 day duration)
		// - Group 2: ACCESS_REQUEST with manual approval (2 hour duration)
		_, err = twingate.NewTwingateResource(ctx, "jit_resource_go", &twingate.TwingateResourceArgs{
			Name:            pulumi.StringPtr("JIT Access Resource Go"),
			Address:         pulumi.String("internal-app.example.com"),
			RemoteNetworkId: remoteNetwork.ID(),
			AccessPolicies: &twingate.TwingateResourceAccessPolicyArray{
				&twingate.TwingateResourceAccessPolicyArgs{
					Mode:         pulumi.String("AUTO_LOCK"),
					Duration:     pulumi.String("7d"),
					ApprovalMode: pulumi.String("MANUAL"),
				},
			},
			AccessGroups: &twingate.TwingateResourceAccessGroupArray{
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group.ID(),
					// Auto-lock: Access automatically expires after duration
					AccessPolicies: &twingate.TwingateResourceAccessGroupAccessPolicyArray{
						&twingate.TwingateResourceAccessGroupAccessPolicyArgs{
							Mode:         pulumi.String("AUTO_LOCK"),
							ApprovalMode: pulumi.String("AUTOMATIC"),
							Duration:     pulumi.String("7d"),
						},
					},
				},
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group2.ID(),
					// Access request: Requires manual approval before granting access
					AccessPolicies: &twingate.TwingateResourceAccessGroupAccessPolicyArray{
						&twingate.TwingateResourceAccessGroupAccessPolicyArgs{
							Mode:         pulumi.String("ACCESS_REQUEST"),
							ApprovalMode: pulumi.String("MANUAL"),
							Duration:     pulumi.String("2h"),
						},
					},
				},
			},
			Protocols: &twingate.TwingateResourceProtocolsArgs{
				Tcp: &twingate.TwingateResourceProtocolsTcpArgs{
					Policy: pulumi.StringPtr("RESTRICTED"),
					Ports: pulumi.StringArray{
						pulumi.String("443"),
						pulumi.String("8080"),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Create a Twingate DNS Filtering Profile
		_, err = twingate.NewTwingateDNSFilteringProfile(ctx, "exampleProfileGo", &twingate.TwingateDNSFilteringProfileArgs{
			Name:           pulumi.String("Go Pulumi DNS Filtering Profile"),
			Priority:       pulumi.Float64(2),
			FallbackMethod: pulumi.String("AUTO"),
			Groups:         pulumi.StringArray{group.ID()},

			AllowedDomains: &twingate.TwingateDNSFilteringProfileAllowedDomainsArgs{
				IsAuthoritative: pulumi.Bool(false),
				Domains: pulumi.StringArray{
					pulumi.String("twingate.com"),
					pulumi.String("zoom.us"),
				},
			},

			DeniedDomains: &twingate.TwingateDNSFilteringProfileDeniedDomainsArgs{
				IsAuthoritative: pulumi.Bool(true),
				Domains: pulumi.StringArray{
					pulumi.String("evil.example"),
				},
			},

			ContentCategories: &twingate.TwingateDNSFilteringProfileContentCategoriesArgs{
				BlockAdultContent: pulumi.Bool(true),
			},

			SecurityCategories: &twingate.TwingateDNSFilteringProfileSecurityCategoriesArgs{
				BlockDnsRebinding:           pulumi.Bool(false),
				BlockNewlyRegisteredDomains: pulumi.Bool(false),
			},

			PrivacyCategories: &twingate.TwingateDNSFilteringProfilePrivacyCategoriesArgs{
				BlockDisguisedTrackers: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}

		// ---------------------------------------------------------------------
		// Twingate Gateway, CAs, and SSH Resource
		// ---------------------------------------------------------------------
		// Generate an RSA CA key + self-signed X.509 certificate at deploy time using
		// the Pulumi tls provider. For prod, replace with your own PKI material.
		x509Key, err := tls.NewPrivateKey(ctx, "x509_key_go", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("RSA"),
			RsaBits:   pulumi.IntPtr(4096),
		})
		if err != nil {
			return err
		}
		x509Cert, err := tls.NewSelfSignedCert(ctx, "x509_cert_go", &tls.SelfSignedCertArgs{
			PrivateKeyPem:       x509Key.PrivateKeyPem,
			IsCaCertificate:     pulumi.BoolPtr(true),
			ValidityPeriodHours: pulumi.Int(8760), // 1 year
			AllowedUses: pulumi.StringArray{
				pulumi.String("key_encipherment"),
				pulumi.String("digital_signature"),
				pulumi.String("cert_signing"),
				pulumi.String("crl_signing"),
			},
			Subject: &tls.SelfSignedCertSubjectArgs{
				CommonName:   pulumi.StringPtr("pulumi-twingate-tls-ca-go"),
				Organization: pulumi.StringPtr("Twingate Example"),
			},
		})
		if err != nil {
			return err
		}

		tlsCa, err := twingate.NewTwingateX509CertificateAuthority(ctx, "tls_ca_go", &twingate.TwingateX509CertificateAuthorityArgs{
			Name:        pulumi.StringPtr("Pulumi TLS CA Go"),
			Certificate: x509Cert.CertPem,
		})
		if err != nil {
			return err
		}

		// Generate an ed25519 SSH CA key — its PublicKeyOpenssh is the OpenSSH-format key.
		sshKey, err := tls.NewPrivateKey(ctx, "ssh_key_go", &tls.PrivateKeyArgs{
			Algorithm: pulumi.String("ED25519"),
		})
		if err != nil {
			return err
		}
		sshCa, err := twingate.NewTwingateSSHCertificateAuthority(ctx, "ssh_ca_go", &twingate.TwingateSSHCertificateAuthorityArgs{
			Name:      pulumi.StringPtr("Pulumi SSH CA Go"),
			PublicKey: sshKey.PublicKeyOpenssh,
		})
		if err != nil {
			return err
		}

		// Gateway sits in the remote network and brokers traffic for downstream resources.
		gateway, err := twingate.NewTwingateGateway(ctx, "gateway_go", &twingate.TwingateGatewayArgs{
			Address:         pulumi.String("10.0.0.1:8443"),
			RemoteNetworkId: remoteNetwork.ID(),
			X509CaId:        tlsCa.ID(),
			SshCaId:         sshCa.ID(),
		})
		if err != nil {
			return err
		}

		// SSH Resource reachable through the Gateway.
		sshResource, err := twingate.NewTwingateSSHResource(ctx, "ssh_resource_go", &twingate.TwingateSSHResourceArgs{
			Name:            pulumi.StringPtr("Bastion SSH Go"),
			Address:         pulumi.String("bastion.internal.example.com"),
			RemoteNetworkId: gateway.RemoteNetworkId,
			GatewayId:       gateway.ID(),
			Username:        pulumi.StringPtr("ubuntu"),
			AccessGroups: &twingate.TwingateSSHResourceAccessGroupArray{
				&twingate.TwingateSSHResourceAccessGroupArgs{
					GroupId: group.ID(),
				},
			},
			Protocols: &twingate.TwingateSSHResourceProtocolsArgs{
				Tcp: &twingate.TwingateSSHResourceProtocolsTcpArgs{
					Policy: pulumi.StringPtr("RESTRICTED"),
					Ports: pulumi.StringArray{
						pulumi.String("22"),
					},
				},
			},
		})
		if err != nil {
			return err
		}

		// Rendered gateway config (e.g. for the gateway runtime to consume).
		gatewayConfig, err := twingate.NewTwingateGatewayConfig(ctx, "gateway_config_go", &twingate.TwingateGatewayConfigArgs{
			Port:        pulumi.IntPtr(8443),
			MetricsPort: pulumi.IntPtr(9090),
			Ssh: &twingate.TwingateGatewayConfigSshArgs{
				Gateway: &twingate.TwingateGatewayConfigSshGatewayArgs{
					Username:    pulumi.StringPtr("gateway"),
					KeyType:     pulumi.StringPtr("ed25519"),
					UserCertTtl: pulumi.StringPtr("5m"),
					HostCertTtl: pulumi.StringPtr("24h"),
				},
				Ca: &twingate.TwingateGatewayConfigSshCaArgs{
					PrivateKeyFile: pulumi.StringPtr("/etc/gateway/ssh-ca.key"),
				},
				Resources: &twingate.TwingateGatewayConfigSshResourceArray{
					&twingate.TwingateGatewayConfigSshResourceArgs{
						Name:     sshResource.Name,
						Address:  sshResource.Address,
						Username: pulumi.String("ubuntu"),
					},
				},
			},
			Tls: &twingate.TwingateGatewayConfigTlsArgs{
				CertificateFile: pulumi.StringPtr("/etc/gateway/tls.crt"),
				PrivateKeyFile:  pulumi.StringPtr("/etc/gateway/tls.key"),
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("gateway_id", gateway.ID())
		ctx.Export("gateway_config_content", gatewayConfig.Content)

		// ---------------------------------------------------------------------
		// Twingate User
		// ---------------------------------------------------------------------
		_, err = twingate.NewTwingateUser(ctx, "example_user_go", &twingate.TwingateUserArgs{
			Email:      pulumi.String("example.user.go@example.com"),
			FirstName:  pulumi.StringPtr("Example"),
			LastName:   pulumi.StringPtr("User"),
			Role:       pulumi.StringPtr("MEMBER"),
			IsActive:   pulumi.BoolPtr(true),
			SendInvite: pulumi.BoolPtr(false),
		})
		if err != nil {
			return err
		}

		// ---------------------------------------------------------------------
		// Connector tokens — bake into a deployment (EC2 user_data, Helm, ...)
		// ---------------------------------------------------------------------
		connectorTokens, err := twingate.NewTwingateConnectorTokens(ctx, "connector_tokens_go", &twingate.TwingateConnectorTokensArgs{
			ConnectorId: connector.ID(),
		})
		if err != nil {
			return err
		}
		ctx.Export("connector_access_token", pulumi.ToSecret(connectorTokens.AccessToken))
		ctx.Export("connector_refresh_token", pulumi.ToSecret(connectorTokens.RefreshToken))

		// ---------------------------------------------------------------------
		// Look up a Security Policy by name and apply it to a group access block.
		// ---------------------------------------------------------------------
		defaultPolicy := twingate.GetTwingateSecurityPolicyOutput(ctx, twingate.GetTwingateSecurityPolicyOutputArgs{
			Name: pulumi.String("Default Policy"),
		})
		_, err = twingate.NewTwingateResource(ctx, "policy_resource_go", &twingate.TwingateResourceArgs{
			Name:            pulumi.StringPtr("Policy-bound Resource Go"),
			Address:         pulumi.String("policy-app.example.com"),
			RemoteNetworkId: remoteNetwork.ID(),
			AccessGroups: &twingate.TwingateResourceAccessGroupArray{
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId:          group.ID(),
					SecurityPolicyId: defaultPolicy.Id(),
				},
			},
		})
		if err != nil {
			return err
		}

		// ---------------------------------------------------------------------
		// Data-source examples (mirroring TS/Python/.NET)
		// ---------------------------------------------------------------------
		connectorsResult, err := twingate.GetTwingateConnectors(ctx, &twingate.GetTwingateConnectorsArgs{
			NameContains: pulumi.StringRef("t"),
		}, nil)
		if err != nil {
			return err
		}
		for _, c := range connectorsResult.Connectors {
			ctx.Log.Info("Connector "+c.Name+" -> "+c.RemoteNetworkId, nil)
		}

		groupsResult, err := twingate.GetTwingateGroups(ctx, &twingate.GetTwingateGroupsArgs{
			NameContains: pulumi.StringRef("demo"),
		}, nil)
		if err != nil {
			return err
		}
		for _, g := range groupsResult.Groups {
			ctx.Log.Info("Group "+g.Name+" (id="+g.Id+")", nil)
		}

		return nil
	})
}
