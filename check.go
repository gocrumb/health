package health

type Check interface {
	Notify(ch chan<- error)
	Stop(ch chan<- error)
}
