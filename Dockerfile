FROM golang:1.23 as builder

FROM scratch

COPY opencost-cacher /app/opencost-cacher
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app/opencost-cacher"]
