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

## Examples

```
timer 12m30s
```

## Compile

```
make
```

or


```
go build -C cmd/clock -o ../../
go build -C cmd/stopwatch -o ../../
go build -C cmd/timer -o ../../
```
