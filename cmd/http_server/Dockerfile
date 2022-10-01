FROM golang:1.18-alpine

WORKDIR /app

RUN ls

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o svr cmd/http_server/*.go

EXPOSE 8080
CMD [ "./svr" ]