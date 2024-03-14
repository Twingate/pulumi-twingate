---
title: Twingate Installation & Configuration
meta_desc: Information on how to install the Twingate provider.
layout: installation
---

## Installation

The Pulumi Twingate provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/twingate`](https://www.npmjs.com/package/@pulumiverse/twingate)
* Python: [`pulumiverse_twingate`](https://pypi.org/project/pulumiverse_twingate/)
* Go: [`github.com/pulumiverse/pulumi-twingate/sdk/go/twingate`](https://pkg.go.dev/github.com/pulumiverse/pulumi-twingate/sdk/go/twingate)
* .NET: [`Pulumiverse.Twingate`](https://www.nuget.org/packages/Pulumiverse.Twingate)


## Configuration

> Note:  
> Replace the following **sample content**, with the configuration options
> of the wrapped Terraform provider and remove this note.

The following configuration points are available for the `twingate` provider:

- `twingate:apiKey` (environment: `twingate_API_KEY`) - the API key for `twingate`
- `twingate:region` (environment: `twingate_REGION`) - the region in which to deploy resources

### Provider Binary

The Twingate provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource twingate <version>
```

Replace the version string `<version>` with your desired version.
