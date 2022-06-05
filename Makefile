
.PHONY=update install linux

install:
	go install
update:
	go mod tidy

linux:
	env GOOS=linux GOARCH=amd64 go build -o gom_amd64

