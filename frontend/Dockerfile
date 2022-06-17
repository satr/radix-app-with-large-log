FROM golang:1.18.3-alpine3.16 as builder

COPY ./app /app
WORKDIR /app

RUN apk update && apk upgrade
RUN apk add ca-certificates curl git  && \
    apk add --no-cache gcc musl-dev

RUN go get -u golang.org/x/lint/golint github.com/frapposelli/wwhrd && \
    go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/frontend
RUN addgroup -S -g 1700 radix
RUN adduser -S -u 1000 -G radix radix

FROM scratch
COPY --from=builder /usr/local/bin/frontend /usr/local/bin/frontend
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER 1000
EXPOSE 8082
ENTRYPOINT ["/usr/local/bin/frontend"]