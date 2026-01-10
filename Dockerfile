# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY api ./api
RUN go build -o server ./api/main.go

# ---- Run stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /app/server /server

EXPOSE 8888
USER nonroot:nonroot
CMD ["/server"]
