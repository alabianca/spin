package spin

import (
	"io"
	"time"
)

type SpinType string

const Dots = SpinType("dots")
const Lines = SpinType("lines")
const Dots2 = SpinType("dots2")

var disableCursor = []byte("\033[?25l")
var enableCursor = []byte("\033[?25h")

type Spinner interface {
	Start()
	Stop()
}

type spinner struct {
	spinType SpinType
	frames   []string
	stop     chan bool
	writer   io.Writer
}

func NewSpinner(st SpinType, writer io.Writer) Spinner {
	frames := getFrames(st)

	return &spinner{
		spinType: st,
		frames:   frames,
		stop:     make(chan bool, 1),
		writer:   writer,
	}
}

func (s *spinner) Start() {
	index := 0
	// disable cursor while spinning
	s.writer.Write(disableCursor)

	for {
		if index >= len(s.frames) {
			index = 0
		}

		frame := "\r" + s.frames[index]

		select {
		case <-time.After(time.Millisecond * 100):
			s.writer.Write([]byte(frame))
			index++
		case <-s.stop:
			s.writer.Write(enableCursor)
			return
		}
	}
}

func (s *spinner) Stop() {
	s.writer.Write([]byte("\n"))
	s.stop <- true
}
