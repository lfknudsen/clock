all: clock stopwatch timer

clock: cmd/clock/main.go
	go build -C cmd/clock -o ../../
	install
	
stopwatch: cmd/stopwatch/main.go
	go build -C cmd/stopwatch -o ../../
	install

timer: cmd/timer/main.go
	go build -C cmd/timer -o ../../
	install
