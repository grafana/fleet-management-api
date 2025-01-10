# Grafana Fleet Management API

This repository contains the Protobuf definitions and [connect-go](https://github.com/connectrpc/connect-go) generated code for the Grafana Cloud [Fleet Management APIs](https://grafana.com/docs/grafana-cloud/send-data/fleet-management/api-reference/).

The generated Go clients can be used to interact with the APIs.

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
