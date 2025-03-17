# --- STAGE 1: Build the Go binary ---
FROM golang:1.20-alpine AS builder

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum first, to cache module downloads
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the project
# The output binary will be named 'ye' (adapt to your liking)
RUN go build -o ye .

# --- STAGE 2: Create a small final image ---
FROM alpine:latest

# Create and set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from builder stage
COPY --from=builder /app/ye .

# If this is a CLI tool that doesn't serve on a port, you don't need EXPOSE
# If needed, e.g., if your Go app listens on port 8080, you might do:
# EXPOSE 8080

# The default command will run the 'ye' binary
CMD ["./ye"]

