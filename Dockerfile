FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /proxy-server

EXPOSE 8080

CMD [ "/proxy-server" ]