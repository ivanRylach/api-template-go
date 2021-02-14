FROM golang:1.15-alpine3.13 as go_builder
RUN apk add --no-cache make
WORKDIR /app
ADD . .
RUN make build

FROM alpine:3.13
WORKDIR /app
COPY --from=go_builder /app/bin/api-service ./app-service
EXPOSE 8080
CMD ./app-service

