FROM golang:1.12-alpine as builder

RUN apk add --no-cache git

WORKDIR /gomod

COPY . .

RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=builder  /gomod/go-echo-server /go-echo-server

ENTRYPOINT [ "/go-echo-server" ]
