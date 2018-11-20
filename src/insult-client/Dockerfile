# Intermediary Build Container
FROM golang:1.9

RUN mkdir -p /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/insult-client
WORKDIR /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/insult-client
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o insult-client .


# Final (static from scratch container)
FROM scratch

ARG VCS_REF
ARG BUILD_DATE

MAINTAINER TransactCharlie
LABEL org.label-schema.name="insult-client" \
      org.label-schema.description="Asks to be insulted" \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/TransactCharlie/golang-for-microservices-talk"

COPY --from=0 /go/src/github.com/transactcharlie/golang-for-microservices-talk/src/insult-client/insult-client /insult-client
ENTRYPOINT ["/insult-client"]
