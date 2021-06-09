package health

//go:generate stringer -type=Status
type Status int

const (
	Stopped Status = iota
	Starting
	Healthy
	Unhealthy
)
