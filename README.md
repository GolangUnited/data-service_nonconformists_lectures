# golang-united-lectures
This repository contains code of gRPC service for lectures.

## Requirements
- Docker installed
- PostgreSQL database installed locally or in the docker container

## How to run the project
1. Set environment variables with correct values:

```bash
export LECTURES_DB_HOST=
export LECTURES_DB_PORT=
export LECTURES_DB_USER=
export LECTURES_DB_PASSWORD=
export LECTURES_DB_DATABASE=
```

2. To start gRPC server:

```bash
make start
```

3. To stop gRPC server:

```bash
make stop
```
