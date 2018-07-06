FROM golang:1.10.1
RUN go get -d -v github.com/gree-gorey/bash-exporter/cmd/bash-exporter
WORKDIR /go/src/github.com/gree-gorey/bash-exporter/cmd/bash-exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bash-exporter .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/github.com/gree-gorey/bash-exporter/cmd/bash-exporter/bash-exporter .
COPY ./example/run.sh /usr/local/bash-exporter/run.sh
CMD ["./bash-exporter"]
