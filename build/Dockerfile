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