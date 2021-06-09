// Code generated by "stringer -type=State"; DO NOT EDIT.

package health

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Stopped-0]
	_ = x[Starting-1]
	_ = x[Healthy-2]
	_ = x[Unhealthy-3]
}

const _State_name = "StoppedStartingHealthyUnhealthy"

var _State_index = [...]uint8{0, 7, 15, 22, 31}

func (i State) String() string {
	if i < 0 || i >= State(len(_State_index)-1) {
		return "State(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}