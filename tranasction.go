package health

type transaction interface{}

type stateChange struct {
	c        Check
	old, new State
}
