all: clock.exe stopwatch.exe timer.exe
	install clock.exe stopwatch.exe timer.exe

clock.exe: cmd/clock/main.go
	go build -C cmd/clock -o ../../clock.exe
	
stopwatch.exe: cmd/stopwatch/main.go
	go build -C cmd/stopwatch -o ../../stopwatch.exe

timer.exe: cmd/timer/main.go
	go build -C cmd/timer -o ../../timer.exe

gnu:

clock: cmd/clock/main.go
	go build -C cmd/clock -o ../../clock

stopwatch: cmd/stopwatch/main.go
	go build -C cmd/stopwatch -o ../../stopwatch

timer: cmd/timer/main.go
	go build -C cmd/timer -o ../../timer
