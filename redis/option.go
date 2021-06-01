package redis

type Option interface {
	apply(*check)
}

type OptionFunc func(*check)

func (o OptionFunc) apply(c *check) {
	o(c)
}
