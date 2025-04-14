FROM golang:alpine AS builder

LABEL app="go-template"

EXPOSE 9210

WORKDIR /src/traction/go-template

RUN go install github.com/air-verse/air@latest

RUN  cp $GOPATH/bin/air /bin/air

RUN which air

RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest

RUN cp $GOPATH/bin/ginkgo /bin/ginkgo

COPY ./ ./

RUN go mod download


ENTRYPOINT [ "air", "-c", "air.toml" ]

