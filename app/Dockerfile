FROM golang:alpine as build

WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . ./
RUN go build -o build -tags jsoniter

FROM scratch

EXPOSE 8080

COPY --from=build /opt/app/build /
COPY --from=build /opt/app/templates /templates/

ENTRYPOINT ["/build"]
