init:
	go mod init
start:
	go run main.go
use:
	curl 127.0.0.1:8080/
	curl 127.0.0.1:8080/metrics
