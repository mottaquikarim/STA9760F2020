FROM golang:1.13 as builder

WORKDIR /app
COPY . .

# Auto restart
RUN go get github.com/githubnemo/CompileDaemon

ENV STA9760F2020_DEBUG 1

# This is your actual go file
RUN go build -o build/site_builder cmd/site_builder/main.go

EXPOSE 80

ENTRYPOINT CompileDaemon --build="go build -o build/site_builder cmd/site_builder/main.go" --command=./build/site_builder --pattern=. --exclude-dir="./docs" --exclude-dir="./build"
