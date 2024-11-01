FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download && go mod verify 

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go

CMD ["app"]
