FROM golang:1.14

WORKDIR /go/src/app

ENV IEX_CLOUD_KEY="your-iex-api-key"

COPY . .

RUN go mod download && go install ./... && go build -o main 

EXPOSE 8082

CMD ["./main"]
