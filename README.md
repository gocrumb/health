# Health

[![Go Reference](https://pkg.go.dev/badge/github.com/gocrumb/health.svg)](https://pkg.go.dev/github.com/gocrumb/health)

Monitor and track health of services, send heartbeats (to-do), and expose status through HTTP endpoint.

## Installation

Install health using the go get command:

```
$ go get github.com/gocrumb/health
```

The package requires no additional dependencies other than Go itself.

## Usage

``` golang
hm := health.New(
	health.Checks(
		http.Head("http://example.com/", http.Period(3*time.Minute)),
		mongo.Ping(mongoClient, mongo.Period(1*time.Minute), mongo.ReadPref(rp)),
		redis.Ping(redisClient), // Falls back to default period (5 minutes)
	),
	health.Logger(log.New(os.Stderr, "[Health] ", log.LstdFlags|log.Lmsgprefix)),
	health.OnFailure(func() {
		log.Fatal("Health check failed; exiting")
	}),
)
hm.Start()
```

## Documentation

- [Reference](https://godoc.org/github.com/gocrumb/health)

## Contributing

Contributions are welcome.

## License

This package is available under the [BSD (3-Clause) License](https://opensource.org/licenses/BSD-3-Clause).
