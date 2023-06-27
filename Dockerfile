FROM golang:1.20-alpine as build
COPY . /app
WORKDIR /app
RUN go build -o user-api
RUN ls -Alp

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/user-api /app/user-api
EXPOSE 8080
ENTRYPOINT ["/app/user-api"]
