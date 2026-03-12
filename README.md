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

## Development

### Building the Provider Locally

To build and test the provider from source:

#### Prerequisites

- **Go 1.24+**: Required to build the provider
- **Node.js 22+**: Required for the Node.js/TypeScript SDK
- **Pulumi CLI**: Install from [pulumi.com](https://www.pulumi.com/docs/install/)

#### Build Steps

1. **Build the provider and SDKs:**

   ```bash
   # For full development build (all SDKs)
   make development

   # For just the provider and Node.js SDK
   make provider build_nodejs
   ```

2. **Install the local provider plugin:**

   After building, you must install the locally built provider as a Pulumi plugin:

   ```bash
   pulumi plugin install resource twingate v4.1.0-alpha.1772811417+dirty \
     --file bin/pulumi-resource-twingate
   ```

   **Note:** The version string will match what's in your build (includes `+dirty` if you have uncommitted changes).

#### Common Errors

**Error: "404 HTTP error fetching plugin"**

If you see this error when running `pulumi up` or `pulumi preview`:

```
error: Could not automatically download and install resource plugin 'pulumi-resource-twingate'
at version v4.1.0-alpha.1772811417+dirty: 404 HTTP error fetching plugin from
https://api.github.com/repos/Twingate/pulumi-twingate/releases/tags/...
```

**Solution:** This happens because you're using a local development build. Install the plugin manually:

```bash
pulumi plugin install resource twingate v4.1.0-alpha.1772811417+dirty \
  --file bin/pulumi-resource-twingate
```

Replace the version string with whatever version your build generated (check the error message for the exact version).

**Verify installation:**

```bash
pulumi plugin ls | grep twingate
```

You should see your local version listed.

### Testing GitHub Actions Workflows Locally

You can test GitHub Actions workflows locally using [`act`](https://github.com/nektos/act).

#### Install act

On macOS with Homebrew:

```bash
brew install act
```

On Linux:

```bash
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash
```

For other platforms, see the [act installation guide](https://github.com/nektos/act#installation).

#### Run workflows locally

To test a specific workflow job, use:

```bash
# List all available jobs
act --list

# Test the lint workflow
act pull_request -j lint

```

**Note:** The first time you run `act`, it will ask you to choose a Docker image size. Select "Medium" for most workflows.
