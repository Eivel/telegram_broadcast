FROM golang:alpine3.13 as golang_builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk --no-cache add ca-certificates
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main main.go


FROM scratch
COPY --from=golang_builder /build/main /app/
COPY --from=golang_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=golang_builder /tmp /tmp
WORKDIR /app
CMD ["./main"]
