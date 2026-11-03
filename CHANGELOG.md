CHANGELOG
=========

## HEAD (Unreleased)

### Added

* `tlsMode` on the `upstream` and `downstream` blocks of `TwingateWebAppResource`.
  The upstream provider added it in v5.2.0, but the schema was not regenerated
  afterwards, so it never reached the SDKs.

### Fixed

* The HCL examples in the [Registry docs](https://www.pulumi.com/registry/packages/twingate/)
  opened with a `pulumi {}` block that the Pulumi HCL language plugin rejects.
  They now use `terraform {}` (#650).
* Identifiers in generated descriptions no longer lose their underscores to
  Markdown emphasis: `VERIFY_FULL`, `VERIFY_CA`, `GOOGLE_CLOUD`, `ON_PREMISE`
  and `twingate_service_account_key` were rendered with `*` in place of `_`.

### terraform-provider-twingate v5.0.0 (breaking)

Upgrades the bridged upstream provider from v4.3.2 to v5.0.0. The SDK module
moves to `github.com/Twingate/pulumi-twingate/sdk/v5`.

**Removed**
* `TwingateGatewayConfig` — the upstream `twingate_gateway_config` resource no
  longer exists.
* `username` and `protocols` on `TwingateSSHResource`.
* `protocols` on `TwingateKubernetesResource`.

  `protocols` is unaffected on `TwingateResource`, which still supports it.

**Added**
* `TwingateWebAppResource` — new upstream `twingate_web_app_resource`. Requires
  `address`, `gatewayId`, `remoteNetworkId`, and `upstream`/`downstream` port
  blocks.

---

pulumi-package-publisher
* Use [Pulumi Package Publisher Action](https://github.com/pulumi/pulumi-package-publisher) in release.yml template
* fix dotnetversion typo in release template #94
* fix `/usr/share/dotnet/host/xfr does not exist` issue when running `make build_dotnet`

---
