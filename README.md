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

```bash
docker compose up -d
```

Create a local `.envrc` file:

```bash
export ADDR=":3000"
export EXTERNAL_URL="localhost:3000"
export DB_ADDR="postgres://admin:adminpassword@localhost/social?sslmode=disable"
export FRONTEND_URL=""
export FROM_EMAIL=""
export SENDGRID_API_KEY=""
```

Then run the migrations and start the API:

```bash
make migrate-up
go run ./cmd/api
```

Open Swagger UI at [http://localhost:3000/v1/swagger/index.html](http://localhost:3000/v1/swagger/index.html).

## Email setup

Registration sends an invitation email through SendGrid. Set `FROM_EMAIL` to a verified SendGrid sender address and add a `SENDGRID_API_KEY` before using the registration endpoint. In non-production environments, SendGrid sandbox mode prevents delivery.

## Useful commands

```bash
make migrate-up   # apply pending migrations
make seed         # create local sample data
make gen-docs     # regenerate Swagger files
go test ./...     # run tests
```
