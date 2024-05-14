FROM golang:1.22

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./


RUN go build -v -o ComputerShop

ENTRYPOINT ["./ComputerShop"]
