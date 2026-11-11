# Clock, Timer, Stopwatch

Three tiny time-related command-line programmes.

## Usage


To print the current date and time:
```
clock
```

To start a timer:
```
timer <duration string>
```

##### Example: 

```
timer 12m30s
```


To start a stopwatch:

```
stopwatch
```

## Compile

If you have an implementation of `make` installed, as well as my
other programme [install](https://gitlab.com/lfknudsen/install),
you can execute:

```
make
```

Alternatively:


```
go build -C cmd/clock -o ../../
go build -C cmd/stopwatch -o ../../
go build -C cmd/timer -o ../../
```
