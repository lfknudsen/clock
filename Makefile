all: clock stopwatch

clock: cmd/clock/main.go
	go build -C cmd/clock -o ../../clock.exe
	
stopwatch: cmd/stopwatch/main.go
	go build -C cmd/stopwatch -o ../../stopwatch.exe