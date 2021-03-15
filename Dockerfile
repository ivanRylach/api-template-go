FROM golang:1.16-alpine3.13 as go_builder
RUN apk add --no-cache make
WORKDIR /app
ADD . .
RUN cd api && make build

FROM alpine:3.13
WORKDIR /app
COPY --from=go_builder /app/api/bin/api-service ./app-service
EXPOSE 8080
CMD ./app-service

