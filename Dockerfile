
FROM alpine:latest

RUN addgroup -g 1000 -S appgroup && \
    adduser -u 1000 -S appuser -G appgroup

WORKDIR /home/appuser

COPY git-id /usr/local/bin/git-id

RUN chown appuser:appgroup /usr/local/bin/git-id && \
    chmod +x /usr/local/bin/git-id

USER appuser

ENTRYPOINT ["git-id"]
