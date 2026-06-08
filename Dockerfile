# ---- Build stage ----
FROM golang:1.25-alpine AS builder

WORKDIR /src

# Cache dependencies first.
COPY go.mod go.sum ./
RUN go mod download

# Build the static binary.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/travelsphere .

# ---- Runtime stage ----
FROM alpine:3.20

# Beego serves templates/static relative to the working directory.
WORKDIR /app

# Non-root user for safety.
RUN adduser -D -u 10001 appuser

# App binary plus the runtime assets Beego expects to find on disk.
COPY --from=builder /out/travelsphere ./travelsphere
COPY conf ./conf
COPY views ./views
COPY static ./static

USER appuser

EXPOSE 8080

ENTRYPOINT ["./travelsphere"]
