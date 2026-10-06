# goSocial

A Go API for a small social app built with PostgreSQL, Chi, Swagger, Zap, bcrypt, and SendGrid.

## What it does

- Create, read, update, and delete posts
- Add comments and follow or unfollow users
- Browse a paginated, searchable feed
- Register users through expiring invitation tokens
- Serve OpenAPI documentation through Swagger UI

## Run it locally

You need Go, Docker, and the [`migrate`](https://github.com/golang-migrate/migrate) CLI.

Build the local Redis image once from the companion repository:

```bash
git clone https://github.com/Intro0/redis-from-scratch-go ../redis-from-scratch-go
docker build -t redis-from-scratch-go:local ../redis-from-scratch-go
```

Then start PostgreSQL and Redis:

```bash
docker compose up -d
```

To use the official Redis image instead of the local custom server:

```bash
REDIS_IMAGE=redis:7-alpine docker compose up -d
```

Create a local `.envrc` file:

```bash
export ADDR=":3000"
export EXTERNAL_URL="localhost:3000"
export DB_ADDR="postgres://admin:adminpassword@localhost/social?sslmode=disable"
export FRONTEND_URL=""
export FROM_EMAIL=""
export SENDGRID_API_KEY=""
export REDIS_ADDR="localhost:6380"
export REDIS_PASSWORD=""
export REDIS_DB="0"
export REDIS_ENABLED="false"
```

Set `REDIS_ENABLED="true"` to use Redis for user-profile caching.

Then run the migrations and start the API:

```bash
make migrate-up
go run ./cmd/api
```

Open Swagger UI at [http://localhost:3000/v1/swagger/index.html](http://localhost:3000/v1/swagger/index.html).

## Email setup

Registration sends an invitation email through SendGrid. Set `FROM_EMAIL` to a verified SendGrid sender address and add a `SENDGRID_API_KEY` before using the registration endpoint. In non-production environments, SendGrid sandbox mode prevents delivery.

## Redis cache

Docker Compose uses the local [`redis-from-scratch-go`](https://github.com/Intro0/redis-from-scratch-go) image by default. Set `REDIS_ENABLED="true"` to enable user-profile caching. The server supports the required `PING`, `GET`, and `SETEX` commands. Cache invalidation will require its `DEL` command first. Use database `0` with no password until that server supports `SELECT` and `AUTH`. You can use the official Redis image through `REDIS_IMAGE` instead.

## Useful commands

```bash
make migrate-up   # apply pending migrations
make seed         # create local sample data
make gen-docs     # regenerate Swagger files
go test ./...     # run tests
```
