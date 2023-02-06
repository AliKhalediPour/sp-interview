FROM golang:1.18-alpine as builder

WORKDIR /go/src

RUN set -xe \
	&& apk add --no-cache gcc libc-dev git curl

COPY go.mod go.sum ./

# download the packages
RUN GOPROXY=https://goproxy.io,direct go mod download

COPY . .

# Build the applications as a staticly one, to allow it to run on alpine version
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o sp-consumer ./cmd

FROM alpine:3.15.0

WORKDIR /srv

COPY --from=builder /go/src/ .

EXPOSE 5002

CMD [ "/srv/sp-consumer" ]