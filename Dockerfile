FROM golang:alpine3.12

RUN mkdir /app

ADD . /app

RUN apk add nmap --no-cache && rm -f /var/cache/apk/*

WORKDIR /app

COPY  go.mod go.sum ./
RUN go mod download 

COPY  . .

RUN chmod +x /app/waitfor.sh
RUN go build main.go

EXPOSE 8080

CMD ["./main"]