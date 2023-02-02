# This is a multi-stage Dockerfile. The first part executes a build in a Golang
# container, and the second retrieves the binary from the build container and
# inserts it into a "scratch" image.

# Part 1: Create a layer for Go module dependencies
#
FROM golang:1.18 as builder

COPY . /mwe

WORKDIR /mwe

RUN GOOS=linux go build -a -installsuffix cgo -o mwe

# Part 3: Build the Mediawiki Exporter image proper
#
FROM ubuntu:22.04 as image

RUN apt update                                              \
  && apt-get -y --force-yes install --no-install-recommends \
  ca-certificates                                           \
  && apt-get clean                                          \
  && apt-get autoclean                                      \
  && apt-get autoremove                                     \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=builder /mwe /bin

EXPOSE 8000

CMD ["mwe"]