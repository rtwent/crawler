Compiling:
    Linux:
        GOOS=linux GOARCH=amd64 go build -o crawler
    Windows
        GOOS=windows GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean