FROM golang:1.18-bullseye as build
ARG SERVICE

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY pkg pkg
COPY internal/${SERVICE} internal/${SERVICE}
COPY cmd/${SERVICE} cmd/${SERVICE}

RUN go build -o /go/bin/goserve ./cmd/...

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build /go/bin/goserve goserve

USER nonroot:nonroot

EXPOSE 8080

ENV TZ=UTC

ENTRYPOINT ["./goserve"]
