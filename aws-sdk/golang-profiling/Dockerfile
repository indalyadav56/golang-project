FROM golang:1.24:alpine as builder

WORKDIR /app

COPY  . .

RUN go build -o main .


EXPOSE 8080

CMD [ "main" ]