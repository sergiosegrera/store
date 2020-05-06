FROM golang:latest

WORKDIR /service
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o service.o ./main.go" --command="./service.o"
