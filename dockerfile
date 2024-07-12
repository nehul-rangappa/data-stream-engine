FROM golang:1.22.2-alpine

WORKDIR /data-stream-engine

ENV KAFKA_CONSUMER_TOPIC user-login
ENV KAFKA_PRODUCER_TOPIC user-analysis

# Alternative approach to avoid copying all files would be to only copy the executable and run it in the container
COPY . .

RUN go mod download

RUN go build -o /stream-engine

CMD [ "/stream-engine" ]