FROM golang:1.22

WORKDIR /app

COPY . .

EXPOSE 8082

RUN go build -o /main

CMD [ "/main" ]