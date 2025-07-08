
FROM alpine:latest

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /home/appuser

COPY gid /usr/local/bin/gid

RUN chown appuser:appgroup /usr/local/bin/gid && \
    chmod +x /usr/local/bin/gid

USER appuser

ENTRYPOINT ["gid"]

CMD ["--help"]
