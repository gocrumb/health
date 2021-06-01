package health

//go:generate gostringer -type=Status
type Status int

const (
	Stopped Status = iota
	Starting
	Healthy
	Unhealthy
)
