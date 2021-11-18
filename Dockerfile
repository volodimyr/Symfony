FROM golang:1.17.3

COPY . /go/src/Symphony
RUN cd /go/src/Symphony && mkdir -p target && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o target/ -v .

FROM alpine:latest
COPY --from=0 /go/src/Symphony/target/Symphony /usr/bin/
ENTRYPOINT [ "/usr/bin/Symphony" ]