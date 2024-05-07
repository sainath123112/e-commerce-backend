FROM golang:1.22

WORKDIR /app

COPY . .

EXPOSE 8081

RUN go build -o /main

CMD [ "/main" ]