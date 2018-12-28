# Compiler image
FROM didstopia/base:go-alpine-3.5 AS go-builder

# Copy the project 
COPY . /tmp/tinyserver/
WORKDIR /tmp/tinyserver/

# Install dependencies
RUN make deps

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/tinyserver



# Runtime image
FROM scratch

# Copy certificates
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the built binary
COPY --from=go-builder /go/bin/tinyserver /go/bin/tinyserver

# Expose environment variables


# Run the binary
ENTRYPOINT ["/go/bin/tinyserver"]
