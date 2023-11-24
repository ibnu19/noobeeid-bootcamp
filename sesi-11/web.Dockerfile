FROM golang AS builder

WORKDIR /apps
COPY . .
RUN CGO_ENABLED=0 go build -o web main.go   

FROM alpine
WORKDIR /apps
COPY --from=builder /apps/web .
ENTRYPOINT [ "./web" ]