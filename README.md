# Twingate Resource Provider

The Twingate Resource Provider lets you manage [Twingate](https://www.twingate.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @twingate/pulumi-twingate
```

or `yarn`:

```bash
yarn add @twingate/pulumi-twingate
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-twingate
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-twingate/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Twingate.Twingate
```

## Configuration

The following configuration points are available for the `twingate` provider:

- `twingate:apiToken` - The access key for API operations. You can retrieve this from the Twingate Admin Console
  ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
  TWINGATE_API_TOKEN environment variable.
- `twingate:network` - Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
  `autoco.twingate.com`, where `autoco` is your network ID. Alternatively, this can be specified using the TWINGATE_NETWORK
  environment variable.
- `twingate:url` - The default is 'twingate.com'. This is optional and shouldn't be changed under normal circumstances.


## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/twingate/api-docs/).
