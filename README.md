# Intervals between timestamps

A program to calculate the interval in seconds between
text timestamps in a stream of timestamps.

Suppose you have a file (or a source of data) where timestamps
appear in ascending order, like this:

```
2026-01-16T00:51:52Z
2026-01-16T03:05:05Z
2026-01-16T03:35:25Z
2026-01-16T05:13:46Z
...
2026-02-05T16:23:29Z
2026-02-05T19:08:21Z
2026-02-05T19:28:17Z
2026-02-05T21:45:22Z
```
That's RFC3339 format, a.k.a. ISO8601 format.
You should keep dates and timestamps in this format.

### Building

I wrote `dateintervals` on a Linux laptop, but it's written in Go,
which is portable.

```
$ git clone https://github.com/bediger4000/dateintervals.git
$ cd dateinterval
$ go build $PWD
```

### Running

```
  -o string
        floating point output format (default "%.0f")
  -t string
        time.Parse timestamp format (default "2006-01-02T15:04:05Z07:00")
```

`-o` allows you to specify a Go `fmt` package output verb for `float64` type values.
The default is no places after the decimal point,
but you could use `%.3f`, or `%g` or `%3.2f`.
The code puts a newline on the output, you don't need to futz around
with any `\n` in the output verb string.

`-t` lets you choose a timestamp format,
in the Go [time](https://pkg.go.dev/time) package's
[layout format](https://pkg.go.dev/time#Parse).
By default, `dateintervals` uses [RFC 3339]()
format: `2006-01-02T15:04:05Z07:00".
That layout matches input strings like "2026-01-17T15:57:17Z".

You can put one input file on the command line, after any flags.
Otherwise, `dateintervals` reads from stdin.
Output is always on stdout.

#### Examples

```
$ ./dateintervals sorted.timestamps
1820
5901
10295
2734
...
```
Read text timestamps from stdin, specify a timestamp layout:

```
$ cat spork                                     # /home/bediger/src/go/src/combined
2026-01-16
2026-01-17
2026-01-18
2026-01-19
2026-01-20
$  cat spork | ./dateintervals -t '2006-01-02'
86400
86400
86400
86400
$
```
