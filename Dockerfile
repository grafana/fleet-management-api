FROM bufbuild/buf:1.61.0@sha256:3b7fb27f8ff153037b3f781924c6aa1327352f14ee547cf45fd75f6d994fe0eb AS buf

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
