FROM golang:1.18-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update 
RUN apk add --no-cache git 
RUN apk add --no-cache ca-certificates tzdata 
RUN apk add --no-cache --update gcc musl-dev 
RUN update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


WORKDIR /app

# use modules
COPY go.* ./

RUN go mod download
RUN go mod verify

COPY cmd ./cmd
COPY pkg ./pkg

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -linkmode external -extldflags "-static"' -a \
    -o /app/server cmd/server/*.go

RUN ls -l /app

############################
# STEP 2 build a small image
############################
FROM alpine:3.11


WORKDIR /app

COPY --from=builder /app/server /app/server

COPY ./docker-config.yaml /app/config.yaml
COPY ./api.yaml /app/api.yaml
COPY ./swagger-ui /app/swagger-ui

EXPOSE 8080

ENTRYPOINT ["./server", "-config", "/app/config.yaml"]