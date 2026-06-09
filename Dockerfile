
FROM golang:1.25-alpine AS builder

WORKDIR /src


COPY go.mod go.sum ./
RUN go mod download


COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/travelsphere .


FROM alpine:3.20

WORKDIR /app


RUN adduser -D -u 10001 appuser

COPY --from=builder /out/travelsphere ./travelsphere
COPY conf ./conf
COPY views ./views
COPY static ./static

USER appuser

EXPOSE 8080

ENTRYPOINT ["./travelsphere"]
