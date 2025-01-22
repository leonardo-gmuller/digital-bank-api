FROM golang:1.22.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY app ./app

RUN GOOS=linux go build -C="app/cmd/server" -ldflags="-w -s" -o digitalbankapi .

FROM alpine:3.18

ARG DB_HOST=postgresdb
ARG DB_DRIVER=postgres
ARG DB_USER=postgres
ARG DB_PASSWORD=postgres
ARG DB_NAME=digital_bank
ARG DB_PORT=5432

ENV DB_HOST=${DB_HOST}
ENV DB_DRIVER=${DB_DRIVER}
ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV DB_PORT=${DB_PORT}

ENV PORT=8000
ENV JWT_SECRET=digitalbank
ENV ENVIRONMENT=production
ENV DEVELOPMENT=false

ENV APP_NAME=digital-bank-api
ENV APP_ID=1
ENV APP_GRACEFUL_SHUTDOWN_TIMEOUT=20s

ENV SERVER_SWAGGER_HOST=0.0.0.0:8000
ENV SERVER_ADDRESS=0.0.0.0:8000
ENV SERVER_READ_TIMEOUT=5s
ENV SERVER_WRITE_TIMEOUT=60s

ENV JWT_SECRET=secret
ENV JWT_SECRET_ADMIN=secret
ENV JWT_EXPIRES_IN=30000

RUN apk add --no-cache bash curl ca-certificates \
    && apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=builder /app/app/cmd/server/digitalbankapi .

EXPOSE 8000

CMD ["./digitalbankapi"]