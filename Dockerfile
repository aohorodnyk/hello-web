FROM golang:1.14-alpine3.11 AS build

WORKDIR /app
COPY main.go /app/main.go
RUN go build -o /app/main /app/main.go


FROM alpine:3.11 AS prod

WORKDIR /app
COPY --from=build /app/main /app/main

CMD ["/app/main"]
