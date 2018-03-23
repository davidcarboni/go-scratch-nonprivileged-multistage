
### Build

FROM golang as builder

# Source
WORKDIR /build
ADD . .

# Compile
RUN CGO_ENABLED=0 go build -o google .

# User definition: only the 'nobody' user
RUN cat /etc/passwd | grep nobody > passwd.nobody


### Run

FROM scratch

# User definition
COPY --from=builder /build/passwd.nobody /etc/passwd

# SSL
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# App
COPY --from=builder /build/google .

# Run
USER nobody
CMD ["./google"]
