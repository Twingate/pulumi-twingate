// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package twingate

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/Twingate/terraform-provider-twingate/v3/twingate"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	bridgev3 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/Twingate/pulumi-twingate/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "twingate"
	// modules:
	mainMod = "index" // the twingate module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-twingate/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() bridgev3.ProviderInfo {
	// Create a Pulumi provider mapping

	provider := twingate.New("Pulumi", version.Version)()

	prov := bridgev3.ProviderInfo{
		// Instantiate the Terraform provider
		P:                       tfbridge.ShimProvider(provider),
		TFProviderModuleVersion: "v3",
		Name:                    "twingate",
		Version:                 version.Version,
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Twingate",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Twingate",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/Twingate/pulumi-twingate/main/docs/logo.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/Twingate/pulumi-twingate",
		Description:       "A Pulumi package for creating and managing Twingate cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "twingate", "category/infrastructure"},
		License:    "Apache-2.0",
		Homepage:   "https://www.twingate.com",
		Repository: "https://github.com/Twingate/pulumi-twingate",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "Twingate",
		MetadataInfo: bridgev3.NewProviderMetadata(metadata),
		Config:       map[string]*bridgev3.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*bridgev3.ResourceInfo{
			"twingate_connector":             {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateConnector")},
			"twingate_connector_tokens":      {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateConnectorTokens")},
			"twingate_dns_filtering_profile": {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateDNSFilteringProfile")},
			"twingate_group":                 {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateGroup")},
			"twingate_remote_network":        {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateRemoteNetwork")},
			"twingate_resource":              {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateResource")},
			"twingate_service_account":       {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateServiceAccount")},
			"twingate_service_account_key":   {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateServiceAccountKey")},
			"twingate_user":                  {Tok: bridgev3.MakeResource(mainPkg, mainMod, "TwingateUser")},
		},
		DataSources: map[string]*bridgev3.DataSourceInfo{
			"twingate_connector":             {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateConnector")},
			"twingate_connectors":            {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateConnectors")},
			"twingate_dns_filtering_profile": {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateDNSFilteringProfile")},
			"twingate_group":                 {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateGroup")},
			"twingate_groups":                {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateGroups")},
			"twingate_remote_network":        {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateRemoteNetwork")},
			"twingate_remote_networks":       {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateRemoteNetworks")},
			"twingate_resource":              {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateResource")},
			"twingate_resources":             {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateResources")},
			"twingate_security_policy":       {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateSecurityPolicy")},
			"twingate_security_policies":     {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateSecurityPolicies")},
			"twingate_service_accounts":      {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateServiceAccounts")},
			"twingate_user":                  {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateUser")},
			"twingate_users":                 {Tok: bridgev3.MakeDataSource(mainPkg, mainMod, "getTwingateUsers")},
		},
		JavaScript: &bridgev3.JavaScriptInfo{
			PackageName: "@twingate/pulumi-twingate",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &bridgev3.PythonInfo{
			PackageName: "pulumi_twingate",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &bridgev3.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/Twingate/pulumi-%[1]s/sdk/", mainPkg),
				bridgev3.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &bridgev3.CSharpInfo{
			RootNamespace: "Twingate",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	// prov.MustComputeTokens(tokens.SingleModule("twingate_", mainMod, tokens.MakeStandard(mainPkg)))

	// TODO: that func call panics
	// panic: Set not supported - is it possible to treat this as immutable?
	// probably there is issue with pulumi-terraform-bridge/pf library compatibility
	// as problem with library method, and not with terraform provider itself
	//
	/// prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}
