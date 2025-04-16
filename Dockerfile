FROM golang:1.24-alpine AS builder
RUN apk update && apk add --no-cache ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt
COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .
RUN go build -o bin/application .

FROM alpine:3.21 AS runner
WORKDIR /opt

RUN apk update && apk add --no-cache libc6-compat curl && rm -rf /var/cache/apk/*
COPY ./healthcheck.sh .
COPY --from=builder /opt/bin/application ./

RUN addgroup -S apprunner && adduser -S apprunner -G apprunner
RUN chown apprunner:apprunner ./application && chown apprunner:apprunner ./healthcheck.sh
RUN chmod +x ./application && chmod +x ./healthcheck.sh
USER apprunner

CMD ["./application"]
HEALTHCHECK --interval=15s --timeout=30s --start-period=5s --retries=3 CMD ["./healthcheck.sh"]