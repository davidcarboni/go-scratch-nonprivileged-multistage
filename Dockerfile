
### Build

FROM golang:onbuild as builder

# Source
WORKDIR /build
ADD . .

# Compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# User definition: only the 'nobody' user
RUN cat /etc/passwd | grep nobody > passwd.nobody


### Run

FROM scratch

# User definition
COPY --from=builder /build/passwd.nobody /etc/passwd

# SSL
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# App
COPY --from=builder /build/main .

# Run
USER nobody
CMD ["./main"]
