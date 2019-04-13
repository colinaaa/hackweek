# build stage
FROM golang:latest AS build
LABEL maintainer="colinwang"

WORKDIR /app

# set go build and go module env for go 1.11
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags 'extldflags="-static"' .
RUN chmod +x ./hackweek
RUN /sbin/ip route|awk '/default/ { print  $3,"\tdockerhost" }' >> hosts

EXPOSE 8080
CMD [ "./hackweek" ]


# cert stage
FROM alpine:3.8 AS certs

RUN apk --no-cache add ca-certificates


# production stage
FROM scratch AS prod

COPY --from=build /app/hackweek .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/hosts /etc/hosts

CMD [ "./hackweek" ]