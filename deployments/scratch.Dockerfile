############################
FROM golang:1.17-alpine AS builder
ENV APP_NAME bill18go
RUN apk add --no-cache tzdata \
    && apk add -U --no-cache ca-certificates \
    && adduser -D -g appuser appuser
WORKDIR /opt/${APP_NAME}
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/${APP_NAME} ./cmd/${APP_NAME}
############################
FROM scratch
ENV APP_NAME bill18go
LABEL name=${APP_NAME} maintainer="Mikhail Puzanov <mpuzanov@mail.ru>" version="1"
WORKDIR /opt/${APP_NAME}
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /opt/${APP_NAME}/bin/${APP_NAME} ./bin/
COPY --from=builder /opt/${APP_NAME}/configs/config-prod.yaml ./configs/
COPY --from=builder /opt/${APP_NAME}/public/ ./public/

EXPOSE 8090
ENTRYPOINT ["./bin/bill18go"]
CMD ["-c","./configs/config-prod.yaml"]

# компиляция образа
#docker build -t puzanovma/bill18go -f scratch.Dockerfile .
#docker run --rm -it -e "PORT=8090" -v $$(pwd)/logs:/app/logs:rw -p 8090:8090 puzanovma/bill18go