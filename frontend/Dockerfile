FROM golang:1.18.3-alpine3.16 as builder

WORKDIR ./app
COPY ./app ./app

RUN apk update && apk upgrade
RUN apk add ca-certificates curl git  && \
    apk add --no-cache gcc musl-dev && \
    go get -u golang.org/x/lint/golint github.com/frapposelli/wwhrd

RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -a -installsuffix cgo -o /usr/local/bin/frontend

FROM scratch
COPY --from=build /usr/local/bin/frontend /usr/local/bin/frontend
RUN addgroup -S -g 1700 radix
RUN adduser -S -u 1000 -G radix radix

USER 1000
EXPOSE 8082
ENTRYPOINT ["/usr/local/bin/frontend"]