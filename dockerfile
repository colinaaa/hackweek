FROM golang:1.11
LABEL maintainer="colinwang"

WORKDIR /app

# set go build and go module env for go 1.11
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on

COPY . .

RUN go build

CMD [ "./hackweek" ]
