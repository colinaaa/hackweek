FROM alpine:3.8
LABEL maintainer="colinwang"

RUN apk --no-cache add tzdata ca-certificates
# set time
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

WORKDIR /app

COPY ./hackweek ./hackweek

CMD [ "./hackweek" ]
