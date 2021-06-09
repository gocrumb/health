package health

//go:generate stringer -type=State
type State int

const (
	Stopped State = iota
	Starting
	Healthy
	Unhealthy
)
