# .NET Example
This example demonstrates how Twingate remote network, service account, service key and resources can be created and configured in C#.

## Resources demonstrated
* `TwingateRemoteNetwork`, `TwingateConnector`, `TwingateConnectorTokens`
* `TwingateServiceAccount`, `TwingateServiceAccountKey`
* `TwingateGroup`, `TwingateUser`
* `TwingateResource` (including JIT access policies and a security-policy-bound access group)
* `TwingateDNSFilteringProfile`
* `TwingateX509CertificateAuthority`, `TwingateSSHCertificateAuthority`
* `TwingateGateway`, `TwingateSSHResource`, `TwingateGatewayConfig`

## Data sources demonstrated
* `GetTwingateConnectors` — filter connectors by name substring
* `GetTwingateGroups` — look up groups by name
* `GetTwingateSecurityPolicy` — resolve a Security Policy by name and bind it to an access group

## Notes
The X.509 CA certificate and SSH CA public key are generated at deploy time using
the [`Pulumi.Tls`](https://www.pulumi.com/registry/packages/tls/) provider
(`PrivateKey` + `SelfSignedCert`). These are **self-signed** — fine for
example/dev tenants; swap in your real PKI material for production.

The connector access/refresh tokens are surfaced as stack outputs (run
`pulumi stack output --show-secrets` to view them) so they can feed into a
downstream connector deployment.

## Pre-requisite
* .NET SDK
* Pulumi

## How to Use
* Clone the repository
* `cd /path/to/repo/examples/dotnet`
* `dotnet restore`
* Setup Pulumi-Twingate Provider, see configuration section [here](../../README.md)
* `pulumi up`

Note: `pulumi up` should automatically download the required .NET packages and Pulumi plugins.
