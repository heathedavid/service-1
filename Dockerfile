# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/heathedavid/service-1
ADD ../../namsra/flag /go/src/github.com/namsral/flag

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/heathedavid/service-1

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/service-1

# Document that the service listens on port 8080.
EXPOSE 8080
