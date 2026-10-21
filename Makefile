all: cmd/clock/main.go cmd/stopwatch/main.go
	go build -C cmd/clock -o ../../clock.exe
	go build -C cmd/stopwatch -o ../../stopwatch.exe