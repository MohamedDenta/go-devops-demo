# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY api ./api
RUN go build -o server ./api/

# ---- Run stage ----
FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=builder /app/server /server

EXPOSE 9090
USER nonroot:nonroot
CMD ["/server"]
