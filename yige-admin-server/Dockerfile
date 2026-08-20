FROM golang:1.22-alpine AS builder

ARG GOPROXY=https://goproxy.cn,direct
ENV GOPROXY=${GOPROXY}
ARG GOSUMDB=sum.golang.google.cn
ENV GOSUMDB=${GOSUMDB}
WORKDIR /src
RUN apk add --no-cache ca-certificates tzdata
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/yige-admin-server ./cmd/server

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata wget \
    && addgroup -S app \
    && adduser -S -G app app \
    && mkdir -p /var/log/yige-admin \
    && chown -R app:app /var/log/yige-admin
WORKDIR /app
COPY --from=builder --chown=app:app /out/yige-admin-server ./yige-admin-server
USER app
EXPOSE 8081
ENTRYPOINT ["./yige-admin-server"]
