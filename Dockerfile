FROM golang:1.4.2

ENV CONN_HOST 127.0.0.1
ENV CONN_PORT 6379

ADD src /go/src
RUN mkdir /go/src/redis-bench
RUN mv /go/src/go-redis-benchmark.go /go/src/redis-bench/main.go

# installing app
RUN go install redis-bench

# remove sources and obj files
RUN rm -rf src/ /go/pkg /go/src
#EXPOSE 8095

# fire app
CMD /go/bin/redis-bench