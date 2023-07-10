FROM golang:1.20-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o user-api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/user-api /app/user-api
EXPOSE 8080
ENTRYPOINT ["/app/user-api"]
