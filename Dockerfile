FROM golang:1.18

RUN apt-get update && apt-get install ffmpeg -y

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /vast

EXPOSE 8080

CMD [ "/vast" ]