# Digest: sha256:0d3653dd6f35159ec6e3d10263a42372f6f194c3dea0b35235d72aabde86486e
FROM golang:1.22.5-alpine3.20

# Prepare working directory
WORKDIR /app

# Copy Go files
COPY go.mod go.sum ./
COPY *.go ./

# Pull Go project dependencies
RUN go mod tidy

# Build Go project
RUN go build -o /echo-server

# Get optional port override
ENV PORT=36110

# Run project
CMD ["/echo-server"]
