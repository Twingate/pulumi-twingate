package main

import (
	"github.com/Twingate/pulumi-twingate/sdk/v3/go/twingate"
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

		// Example: Create a Resource with JIT (Just-In-Time) Access Policy
		// Apply the policy directly to the access group for it to take effect
		_, err = twingate.NewTwingateResource(ctx, "jit_resource_go", &twingate.TwingateResourceArgs{
			Name:            pulumi.StringPtr("JIT Access Resource Go"),
			Address:         pulumi.String("internal-app.example.com"),
			RemoteNetworkId: remoteNetwork.ID(),
			AccessGroups: &twingate.TwingateResourceAccessGroupArray{
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group.ID(),
					// Access policy must be set on the group level when using access_groups
					AccessPolicies: &twingate.TwingateResourceAccessGroupAccessPolicyArray{
						&twingate.TwingateResourceAccessGroupAccessPolicyArgs{
							Mode:         pulumi.String("AUTO_LOCK"), // Automatically lock access after duration
							ApprovalMode: pulumi.String("AUTOMATIC"), // No manual approval required
							Duration:     pulumi.String("24h"),       // Access granted for 24 hours
						},
					},
				},
			},
			Protocols: &twingate.TwingateResourceProtocolsArgs{
				AllowIcmp: pulumi.BoolPtr(true),
				Tcp: &twingate.TwingateResourceProtocolsTcpArgs{
					Policy: pulumi.StringPtr("ALLOW_ALL"),
				},
			},
		})
		if err != nil {
			return err
		}

		// Example: Create a Resource with group-specific Access Policies
		_, err = twingate.NewTwingateResource(ctx, "group_policy_resource_go", &twingate.TwingateResourceArgs{
			Name:            pulumi.StringPtr("Group-Specific Policy Resource Go"),
			Address:         pulumi.String("sensitive-app.example.com"),
			RemoteNetworkId: remoteNetwork.ID(),
			AccessGroups: &twingate.TwingateResourceAccessGroupArray{
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group.ID(),
					// This group gets auto-lock access
					AccessPolicies: &twingate.TwingateResourceAccessGroupAccessPolicyArray{
						&twingate.TwingateResourceAccessGroupAccessPolicyArgs{
							Mode:         pulumi.String("AUTO_LOCK"),
							ApprovalMode: pulumi.String("AUTOMATIC"),
							Duration:     pulumi.String("8h"),
						},
					},
				},
				&twingate.TwingateResourceAccessGroupArgs{
					GroupId: group2.ID(),
					// This group requires manual access requests
					AccessPolicies: &twingate.TwingateResourceAccessGroupAccessPolicyArray{
						&twingate.TwingateResourceAccessGroupAccessPolicyArgs{
							Mode:         pulumi.String("ACCESS_REQUEST"),
							ApprovalMode: pulumi.String("MANUAL"), // Requires manual approval
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
		_, err = twingate.NewTwingateDNSFilteringProfile(ctx, "dns_profile", &twingate.TwingateDNSFilteringProfileArgs{
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

		return nil
	})
}
