# Grafana Fleet Management API

This repository contains the Protobuf definitions and [connect-go](https://github.com/connectrpc/connect-go) generated code for the Grafana Cloud [Fleet Management APIs](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/).

The generated Go clients can be used to interact with the APIs.

## API stability

The stability of an API is signaled by the version suffix of its Protobuf
package. Stability levels are defined as follows:

- **Alpha (`vNalphaM`, e.g. `pipeline.v2alpha1`)**: undergoes rapid iteration.
  Breaking changes are made freely, and there is no expectation of stability.
- **Beta (`vNbetaM`, e.g. `pipeline.v2beta1`)**: complete and ready to be
  declared stable, subject to public testing. A beta API is stable but may still
  change, including backwards-incompatible changes after a deprecation period.
- **Stable (`vN`, e.g. `pipeline.v1`)**: backwards-incompatible changes are not
  made; a breaking change requires a new major version.
  - In very rare cases, such as security concerns or regulatory compliance, breaking changes may be made to a stable API, possibly without a deprecation window.

An API graduates by being renamed in place (`v2alpha1` → `v2beta1` → `v2`); only
one version of a given line exists at a time.

## Updating the Protobuf definitions

Make the changes to the relevant `/api/*/*.proto` files.

To build the container that has the required dependencies, run:

```bash
make build-container
```

To generate the code with the new Protobuf definitions, run the following command to regenerate the code in a Docker container:

```bash
make buf-generate
```
