####################
# Build in alpine
####################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata apache2 && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/sedeo/backend-interview/
COPY ../.. .
#RUN GIT_COMMIT=$(git rev-parse --short HEAD)

RUN go mod download
RUN go mod verify

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/service cmd/service/main.go 

####################
# Package in scratch
####################

FROM alpine
# Could be from scratch

# Import from builder.
COPY --from=builder /etc/apache2/mime.types /etc/mime.types
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd 
# Copy our static executable
COPY --from=builder /go/bin/service /go/bin/service 
# Use an unprivileged user.
#USER appuser

EXPOSE 9111

ENTRYPOINT ["/go/bin/service"]