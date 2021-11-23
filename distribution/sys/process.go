package mysys

type process struct {
	pid       int
	size      int
	startAddr int
}

func NewProcess(pid int, size int) *process {
	return &process{
		pid:       pid,
		size:      size,
		startAddr: 0,
	}
}
