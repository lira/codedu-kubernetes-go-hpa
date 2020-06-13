FROM golang:alpine as builder
RUN apk update
COPY src/hpa .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hpa .

FROM scratch
COPY --from=builder /go/hpa /usr/bin/hpa
EXPOSE 8000
CMD ["hpa"]