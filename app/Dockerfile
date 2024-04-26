FROM golang

EXPOSE 8080

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . ./
RUN go build -o build

CMD ["./build"]
