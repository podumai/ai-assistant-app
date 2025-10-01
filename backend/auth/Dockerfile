ARG ALPINE_IMAGE_VERSION=3.22
ARG SERVICE_PORT=8080
ARG GRPC_PORT=8081

FROM golang:1.25.1-alpine${ALPINE_IMAGE_VERSION} AS build

COPY . /project
WORKDIR /project
RUN CGO_ENABLED=0 go build cmd/main.go

FROM alpine:${ALPINE_IMAGE_VERSION}

COPY --from=build /project/main /app/
COPY --from=build /project/conf.yaml /etc/auth_service/
EXPOSE ${SERVICE_PORT}
EXPOSE ${GRPC_PORT}
ENTRYPOINT ["/app/main"]
