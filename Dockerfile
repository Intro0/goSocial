# Build a statically linked Linux binary.
FROM golang:1.27.1-alpine AS builder

WORKDIR /src

# Download dependencies separately so Docker can reuse this layer when only
# application code changes.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=development
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w -X main.version=${VERSION}" -o /out/api ./cmd/api

# scratch contains no shell or package manager: only the API and the CA bundle
# needed for outbound HTTPS requests such as SendGrid.
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /out/api /api

USER 65532:65532
EXPOSE 8080

ENTRYPOINT ["/api"]
