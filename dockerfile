# build stage
FROM golang:1.11 AS build
LABEL maintainer="colinwang"

WORKDIR /app

# set go build and go module env for go 1.11
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLE=0 GOOS=linux go build .

CMD [ "./hackweek" ]


# cert stage
FROM alpine:3.8 AS certs

RUN apk --no-cache add tzdata  ca-certificates && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone


# production stage
FROM scratch AS prod

COPY --from=build /app /app
COPY --from=certs /etc/localtime /etc/localtime
COPY --from=certs /etc/timezone /etc/timezone
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD [ "./hackweek" ]