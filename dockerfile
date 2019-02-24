FROM golang:latest
LABEL maintainer="colinwang"

# set time
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY ./hackweek ./hackweek

CMD [ "./hackweek" ]
