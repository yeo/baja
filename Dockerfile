FROM golang:1.12-stretch as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /baja .


FROM scratch

COPY --from=builder /baja /usr/bin/baja

CMD ["/usr/bin/baja"]
