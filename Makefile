init:
	go mod download
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o atomicals-cli main.go