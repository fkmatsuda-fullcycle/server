FROM fkmatsuda/golang-builder:1.18.0-1-alpine AS builder
RUN mkdir -p /src/server
COPY . /src/server/
WORKDIR /src/server
RUN go build -buildvcs=false -o /usr/bin/server .

FROM alpine:3.15
COPY --from=builder /usr/bin/server /usr/bin/server
CMD ["/usr/bin/server"]
