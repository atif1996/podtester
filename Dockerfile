# build stage
FROM golang:1.13-alpine3.10 AS builder
ENV GOPATH=""
COPY .  /build
WORKDIR /build 
RUN apk add --no-cache make
RUN go get -u golang.org/x/lint/golint
RUN PATH="~/go/bin:$PATH"
RUN make binary 

#second stage
FROM atif1996/alpine-awscli-kubectl:1.1.0
WORKDIR /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/podtester podtester

ENTRYPOINT ["/podtester"]
