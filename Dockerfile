
### Build

FROM golang:onbuild as builder

# Source
WORKDIR /app
ADD . /app/

# Compile
ENV CGO_ENABLED=0 
ENV GOOS=linux 
RUN go build -a -installsuffix cgo -o main .

# User definition: only the 'nobody' user
RUN cat /etc/passwd | grep nobody > passwd.nobody


### Run

FROM scratch

# User definition
COPY --from=builder /app/passwd.nobody /etc/passwd

# App
COPY --from=builder /app/main .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run
USER nobody
CMD ["./main"]
