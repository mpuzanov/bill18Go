FROM golang:1.13-alpine AS builder
ENV APP_NAME bill18go
RUN apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates \
    && adduser -D -g appuser appuser
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/bill18go ./cmd/${APP_NAME}

FROM alpine:1.13
ENV APP_NAME bill18go
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"
WORKDIR /app/${APP_NAME}
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /app/${APP_NAME}/bin/${APP_NAME} ./bin/
COPY --from=builder /app/${APP_NAME}/configs/config-prod.yaml ./configs/
COPY --from=builder /app/${APP_NAME}/public ./public/
RUN chown -R appuser .
USER appuser
EXPOSE 8090
ENTRYPOINT ["./bin/bill18go"]
CMD ["-c","./configs/config-prod.yaml"]
