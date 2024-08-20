FROM golang:1.23 as build

WORKDIR /app

COPY . .

WORKDIR /app

RUN go build .

CMD [ "./goroapi" ]