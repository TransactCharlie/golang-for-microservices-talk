# Intermediary Build Container
FROM golang:1.9

RUN mkdir -p /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/principals-server
WORKDIR /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/principals-server
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o principals-server .


# Final (static from scratch container)
FROM scratch
COPY --from=0 /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/principals-server/principals-server /principals-server
ENTRYPOINT ["/principals-server"]
