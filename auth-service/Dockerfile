# Use a lightweight Alpine base
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the pre-built auth-service binary into the image
COPY auth-service .

# Ensure the binary is executable
RUN chmod +x auth-service

EXPOSE 8080

CMD ["./auth-service"]