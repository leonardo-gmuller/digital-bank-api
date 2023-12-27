FROM golang:latest

ARG DB_HOST=postgresdb
ARG DB_DRIVER=postgres
ARG DB_USER=root
ARG DB_PASSWORD=root
ARG DB_NAME=postgres
ARG DB_PORT=5432

ENV DB_HOST ${DB_HOST}
ENV DB_DRIVER ${DB_DRIVER}
ENV DB_USER ${DB_USER}
ENV DB_PASSWORD ${DB_PASSWORD}
ENV DB_NAME ${DB_NAME}
ENV DB_PORT ${DB_PORT}
ENV PORT 8000

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -C src -o /digital-bank-api

EXPOSE $PORT

# Run
CMD ["/digital-bank-api"]