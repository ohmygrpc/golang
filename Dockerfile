FROM golang:1.16.3 as build

ENV ORG_NAME=ohmygrpc
ENV SERVICE_NAME=golang

ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /${SERVICE_NAME}/bin
COPY ./bin ./

RUN if [ "$BUILDPLATFORM" = "linux/amd64" ]; then mv ${SERVICE_NAME}.linux.amd64 ${SERVICE_NAME} ; fi
RUN if [ "$BUILDPLATFORM" = "linux/arm64" ]; then mv ${SERVICE_NAME}.linux.arm64 ${SERVICE_NAME} ; fi


FROM --platform=$BUILDPLATFORM gcr.io/distroless/base

ENV ORG_NAME=ohmygrpc
ENV SERVICE_NAME=golang

ARG TARGETPLATFORM
ARG BUILDPLATFORM

COPY --from=build /${SERVICE_NAME}/bin/${SERVICE_NAME} /app/${SERVICE_NAME}
ENTRYPOINT ["app/${SERVICE_NAME}"]
