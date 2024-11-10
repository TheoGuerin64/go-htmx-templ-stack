FROM golang:1.23.2 AS dev
COPY --from=ghcr.io/a-h/templ /ko-app/templ /usr/local/bin/templ
RUN go install github.com/air-verse/air@latest
WORKDIR /app/
COPY ./go.mod ./go.sum ./
RUN go mod download
CMD ["air"]

FROM debian:12-slim AS templ-builder
COPY --from=ghcr.io/a-h/templ /ko-app/templ /usr/local/bin/templ
WORKDIR /app/
COPY ./go.mod ./internal/templates/ ./
RUN templ generate

FROM golang:1.23.2 AS go-builder
WORKDIR /app/
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY --from=templ-builder /app/*.go ./internal/templates/
COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
RUN go build ./cmd/main.go

FROM debian:12-slim AS prod
WORKDIR /app/
COPY --from=go-builder /app/main ./
COPY ./static/ ./static/
CMD ["/app/main"]
