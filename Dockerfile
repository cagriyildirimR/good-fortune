FROM golang:1.20
LABEL authors="cagriyildirim"

WORKDIR /usr/src/app

COPY . .

RUN go build

EXPOSE 8080

CMD ["./GoodFortune"]