all: cmd/clock/main.go
	go build -C cmd/clock -o ../../clock.exe