FROM golang:1.22.2 AS builder

WORKDIR /app

# Set Go proxy and module settings
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV GO111MODULE=on

# Copy go.mod and go.sum (if it exists)
COPY go.mod ./
COPY go.sum* ./

# Download dependencies explicitly
RUN go mod download && go mod verify

# Copy the rest of the source code
COPY . .

# Tidy up the go.mod and go.sum files
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.19

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./main"]