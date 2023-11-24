FROM golang AS builder

WORKDIR /apps
COPY . .
RUN CGO_ENABLED=0 go build -o app main.go

FROM alpine
WORKDIR /apps
COPY --from=builder /apps/app .
ENTRYPOINT [ "./app" ]