##
# BUILD CONTAINER
##

FROM golang:1.12 as builder

WORKDIR /go/src/github.com/mvisonneau/mmds

COPY Makefile .
RUN \
make setup

COPY . .
RUN \
make deps ;\
make build-docker

##
# RELEASE CONTAINER
##

FROM scratch

WORKDIR /

COPY --from=builder /go/src/github.com/mvisonneau/mmds/mmds /

ENTRYPOINT ["/mmds"]
CMD [""]
