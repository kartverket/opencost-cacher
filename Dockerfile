FROM scratch

COPY opencost-cacher /app/opencost-cacher

ENTRYPOINT ["/app/opencost-cacher"]
