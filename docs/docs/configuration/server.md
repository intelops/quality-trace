# Configuring the Qualitytrace Server

Qualitytrace requires a very minimal configuration to be launched, needing just the connection information to connect with the PostgreSQL database which is installed as part of the server install. There are a couple of ways to provide this database connection information.

For Docker-based installs, the server configuration file is placed in the `./quality-trace/quality-trace.yaml` file by default when you run the `quality-trace server install` command and select the `Using Docker Compose` option. The configuration file is mounted to `/app/config.yaml` within the Qualitytrace Docker container. When Qualitytrace is run with a `docker compose -f quality-trace/docker-compose.yaml  up -d` command, the server will use the contents of this file to connect to the Postgres database. All other configuration data is stored in the Postgres instance.

This is an example of a `quality-trace.yaml` file:

```yaml
postgres:
  host: postgres
  user: postgres
  password: postgres
  port: 5432
  dbname: postgres
  params: sslmode=disable

server: 
  httpPort: 11633
  pathPrefix: /
```

Alternatively, we support setting a series of environment variables which can contain the connection information for the Postgres instance. If these environment variables are set, they will be used by the Qualitytrace server to connect to the database.

The list of environment variables and example values is:
- `QUALITYTRACE_POSTGRES_HOST: "postgres"`
- `QUALITYTRACE_POSTGRES_PORT: "5432"`
- `QUALITYTRACE_POSTGRES_DBNAME: "postgres"`
- `QUALITYTRACE_POSTGRES_USER: "postgres"`
- `QUALITYTRACE_POSTGRES_PASSWORD: "postgres"`
- `QUALITYTRACE_POSTGRES_PARAMS: ""`

You can also change the defaults for the Qualitytrace server, altering the port and path you access the dashboard from. For a Docker-based install done locally, this URL is typically `http://localhost:11633/`. You can alter it by setting either of the environment variables or using the `server` object in the server configuration file:

- `QUALITYTRACE_SERVER_HTTPPORT=11633`
- `QUALITYTRACE_SERVER_PATHPREFIX="/"`


You can also intitalize the server with a number of resources the first time it is launched by using [provisioning](./provisioning).

