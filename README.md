NATS publisher/subscriber implementation 
=============

[![Build Status](https://travis-ci.org/ProtocolONE/nats.svg?branch=master)](https://travis-ci.org/ProtocolONE/nats) [![codecov](https://codecov.io/gh/ProtocolONE/nats/branch/master/graph/badge.svg)](https://codecov.io/gh/ProtocolONE/nats)

## Environment variables:

| Name              | Required | Default                  | Description                        |
|:------------------|:--------:|:-------------------------|:-----------------------------------|
| NATS_SERVER_URLS  | -        | 127.0.0.1:4222           | List of NATS server urls           |
| NATS_CLUSTER_ID   | -        | test-cluster             | Identity of cluster to connect     |
| NATS_CLIENT_ID    | -        | billing-server-publisher | Client connection ID               |
| NATS_CLIENT_NAME  | -        | NATS publisher           | Client connection name             |
| NATS_ASYNC        | -        | false                    | Use async mode to publish messages | 
| NATS_USER         | -        | ""                       | User name for connection           |
| NATS_PASSWORD     | -        | ""                       | Password for connection            |

## Contributing
We feel that a welcoming community is important and we ask that you follow PaySuper's [Open Source Code of Conduct](https://github.com/paysuper/code-of-conduct/blob/master/README.md) in all interactions with the community.

PaySuper welcomes contributions from anyone and everyone. Please refer to each project's style and contribution guidelines for submitting patches and additions. In general, we follow the "fork-and-pull" Git workflow.

The master branch of this repository contains the latest stable release of this component.