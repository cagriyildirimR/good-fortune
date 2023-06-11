FROM golang:1.18-alpine as builder

WORKDIR /usr/src/app

COPY . .

RUN adduser -D appuser &&\
    CGO_ENABLED=0 go build -o GoodFortune

FROM scratch

LABEL authors="cagriyildirim"
EXPOSE 8080

COPY --from=builder /usr/src/app/GoodFortune /
COPY --from=builder /etc/passwd /etc/passwd
USER appuser

ENTRYPOINT ["./GoodFortune"]

# Old Dockerfile
#FROM golang:1.18-alpine
#LABEL authors="cagriyildirim"
#
#WORKDIR /usr/src/app
#
#COPY . .
#
#RUN go build
#
#EXPOSE 8080
#
#CMD ["./GoodFortune"]