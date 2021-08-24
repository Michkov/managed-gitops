# managed-gitops

## Overview

See [Google Drive](https://drive.google.com/drive/u/0/folders/1p_yIOJ1WLu-lqz-BVDn076l1K1pEOc1d)


### Components

There are 3 separated, tighly-coupled components contained within this repository:
- **Frontend**: [PatternFly](https://www.patternfly.org/)/React based based on PatternFly React Seed
- **Backend**: Simple Go module for listening for REST requests
- **Cluster Agent**: A Kubernetes controller/operator that lives on the cluster and handles cluster operations.
    - Based on operator-sdk v1.11


### Development Environment

This repository contains scripts which may be used to setup/run the development environment:
- **`create-dev-env.sh`**: Start up Postgresql in a docker container, with the database initialized
- **`(delete/stop)-dev-env.sh`**: Stop or delete the database.
- **`db-schema.sql`**: The database schema used by the components
- **`psql.sh`**: Allows you to interact with the DB from the command line. Requires `psql` CLI util to be installed on your local machine (for example, by installing PostgreSQL)

