FROM golang:1.23.4-alpine3.19 AS builder

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -trimpath -o=bin/main ./cmd

FROM gcr.io/distroless/static-debian11

ENV TZ=Asia/Tokyo

COPY --chown=nonroot:nonroot --from=builder /backend/bin/main .
USER nonroot

ENTRYPOINT [ "./main" ]