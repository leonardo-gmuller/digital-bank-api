FROM golang:latest

ARG DB_HOST=postgresdb
ARG DB_DRIVER=postgres
ARG DB_USER=postgres
ARG DB_PASSWORD=postgres
ARG DB_NAME=digital_bank
ARG DB_PORT=5432

ENV DB_HOST postgresdb
ENV DB_DRIVER postgres
ENV DB_USER postgres
ENV DB_PASSWORD postgres
ENV DB_NAME digital_bank
ENV DB_PORT 5432

ENV PORT 8000

ENV JWT_SECRET digitalbank

ENV ENVIRONMENT production
ENV DEVELOPMENT false

ENV APP_NAME digital-bank-api
ENV APP_ID 1
ENV APP_GRACEFUL_SHUTDOWN_TIMEOUT 20s

ENV SERVER_SWAGGER_HOST 0.0.0.0:8000
ENV SERVER_ADDRESS 0.0.0.0:8000
ENV SERVER_READ_TIMEOUT 5s
ENV SERVER_WRITE_TIMEOUT 60s

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./src ./src

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -C src -o /digital-bank-api

EXPOSE $PORT

# Run
CMD ["/digital-bank-api"]