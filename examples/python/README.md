# Python Example
This example demonstrate how Twingate remote network, service account, service key and resources can be created and configured in Python.

## Resources demonstrated
* `TwingateRemoteNetwork`, `TwingateConnector`, `TwingateConnectorTokens`
* `TwingateServiceAccount`, `TwingateServiceAccountKey`
* `TwingateGroup`, `TwingateUser`
* `TwingateResource` (including JIT access policies and a security-policy-bound access group)
* `TwingateDNSFilteringProfile`
* `TwingateX509CertificateAuthority`, `TwingateSSHCertificateAuthority`
* `TwingateGateway`, `TwingateSSHResource`, `TwingateWebAppResource`

## Data sources demonstrated
* `get_twingate_connectors` — filter connectors by name substring
* `get_twingate_groups_output` — look up groups by name
* `get_twingate_security_policy` — resolve a Security Policy by name and bind it to an access group

## Notes
The X.509 CA certificate and SSH CA public key are generated at deploy time using
the [`pulumi_tls`](https://www.pulumi.com/registry/packages/tls/) provider
(`tls.PrivateKey` + `tls.SelfSignedCert`). These are **self-signed** — fine for
example/dev tenants; swap in your real PKI material for production.

The connector access/refresh tokens are surfaced as stack outputs (run
`pulumi stack output --show-secrets` to view them) so they can feed into a
downstream connector deployment.

## Pre-requisite
* Python and PIP
* Pulumi

## How to Use
* Clone the repository
* `cd /path/to/repo/examples/python`
* `virtualenv venv`
* `source venv/bin/activate`
* `./venv/bin/pip install -r requirements.txt`
* Setup Pulumi-Twingate Provider, see configuration section [here](../../README.md)
* `pulumi up`

Note: `pulumi up` should automatically download the required Python dependency and Pulumi Plugins.