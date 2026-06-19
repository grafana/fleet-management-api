FROM bufbuild/buf:1.71.0@sha256:7f3e3dfb8650f39878625bbc9f2016a51a781693b209165671d5a61d11c74992 AS buf

FROM golang:1.23.5

# Copy the buf binary from the buf image to the final image
COPY --from=buf /usr/local/bin/buf /usr/local/bin/buf

# Install protoc-gen-go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2

# Install protoc-gen-connect-go
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.17.0

# Set the working directory inside the container
WORKDIR /api

# Copy the files from the current directory into the container
COPY . .

ENTRYPOINT [ "buf" ]

CMD ["generate", "api", "-o", "./api"]
