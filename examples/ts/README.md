# Typescript Example
This example demonstrate how Twingate remote network, service account, service key and resources can be created and configured in Typescript.

## Resources demonstrated
* `TwingateRemoteNetwork`, `TwingateConnector`, `TwingateConnectorTokens`
* `TwingateServiceAccount`, `TwingateServiceAccountKey`
* `TwingateGroup`, `TwingateUser`
* `TwingateResource` (including JIT access policies and a security-policy-bound access group)
* `TwingateDNSFilteringProfile`
* `TwingateX509CertificateAuthority`, `TwingateSSHCertificateAuthority`
* `TwingateGateway`, `TwingateSSHResource`, `TwingateGatewayConfig`

## Data sources demonstrated
* `getTwingateConnectors` — filter connectors by name substring
* `getTwingateGroupsOutput` — look up groups by name
* `getTwingateSecurityPolicyOutput` — resolve a Security Policy by name and bind it to an access group

## Notes
The X.509 CA certificate and SSH CA public key are generated at deploy time using
the [`@pulumi/tls`](https://www.pulumi.com/registry/packages/tls/) provider
(`tls.PrivateKey` + `tls.SelfSignedCert`). These are **self-signed** — fine for
example/dev tenants; swap in your real PKI material for production.

The connector access/refresh tokens are surfaced as stack outputs (run
`pulumi stack output --show-secrets` to view them) so they can feed into a
downstream connector deployment.

## Pre-requisite
* NodeJS
* Pulumi

## How to Use
* Clone the repository
* `cd /path/to/repo/examples/ts`
* `npm install pulumi`
* Setup Pulumi-Twingate Provider, see configuration section [here](../../README.md)
* `pulumi up`

Note: `pulumi up` should automatically download the required typescript dependency and Pulumi Plugins.