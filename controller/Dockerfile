## compile controller services
FROM gcr.io/linkerd-io/go-deps:925bb845 as golang
WORKDIR /linkerd-build
COPY controller/gen controller/gen
COPY pkg pkg
COPY controller controller

# use `install` so that we produce multiple binaries
RUN CGO_ENABLED=0 GOOS=linux go install -mod=readonly ./pkg/...
RUN CGO_ENABLED=0 GOOS=linux go install -mod=readonly ./controller/cmd/...

## package runtime
FROM scratch
ENV PATH=$PATH:/go/bin
COPY LICENSE /linkerd/LICENSE
COPY --from=golang /go/bin /go/bin

ARG LINKERD_VERSION
ENV LINKERD_CONTAINER_VERSION_OVERRIDE=${LINKERD_VERSION}
