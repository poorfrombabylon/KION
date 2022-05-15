package domain

type Event string

func (e Event) String() string {
	return e.String()
}

const (
	StopVideoEvent    Event = "STOP_VIDEO_EVENT"
	ForwardVideoEvent Event = "FORWARD_VIDEO_EVENT"
	RevertVideoEvent  Event = "REVERT_VIDEO_EVENT"
	ExitVideoEvent    Event = "EXIT_VIDEO_EVENT"
	Nothing           Event = "kek"
)
