FROM golang:1.14.0 AS builder

ENV GOPATH=/go

COPY . /go/src/dop
WORKDIR /go/src/dop

RUN go mod download -x \
    && go build -a -installsuffix cgo main.go

EXPOSE 8080
CMD ["/go/src/dop/main"]

#FROM alpine:3.11.3
#RUN adduser -D -h /app app
#WORKDIR /app
#COPY --from=builder /go/src/dop/dop .
#COPY --from=builder /go/src/dop/conf ./conf
#COPY --from=builder /go/src/dop/static ./static
#COPY --from=builder /go/src/dop/views ./views
#EXPOSE 8080
#CMD ["/app/dop"]
