FROM golang:1.16.3 as builder

ENV ORG_NAME=ohmygrpc
ENV SERVICE_NAME=golang

ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /${SERVICE_NAME}
COPY . .

WORKDIR /${SERVICE_NAME}/cmd

RUN CGO_ENABLED=0 \
    GOOS=$(echo "$TARGETPLATFORM" | cut -d '/' -f1) \
    GOARCH=$(echo "$TARGETPLATFORM" | cut -d '/' -f2) \
    go build -a -installsuffix cgo -ldflags="-w -s" -o /${SERVICE_NAME}/bin/${SERVICE_NAME}


FROM --platform=$TARGETPLATFORM gcr.io/distroless/base

ENV ORG_NAME=ohmygrpc
ENV SERVICE_NAME=echo

ARG TARGETPLATFORM
ARG BUILDPLATFORM

COPY --from=builder /${SERVICE_NAME}/bin/${SERVICE_NAME} /app/${SERVICE_NAME}
ENTRYPOINT ["app/${SERVICE_NAME}"]
