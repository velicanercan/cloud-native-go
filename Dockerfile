FROM golang:alpine

WORKDIR /cloud-native-go
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api \
    && go build -o ./bin/migrate ./cmd/migrate

CMD ["/cloud-native-go/bin/api"]
EXPOSE 8080