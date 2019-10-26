package spin

import (
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SpinType string

const Dots = SpinType("dots")
const Lines = SpinType("lines")
const Dots2 = SpinType("dots2")
const csi = "\033[?"

var disableCursor = []byte(csi + "25l")
var enableCursor = []byte(csi + "25h")
var clearAllLine = []byte(csi + "2K")

type Spinner interface {
	Start()
	Stop()
	Close()
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
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for {
		if index >= len(s.frames) {
			index = 0
		}

		frame := "\r" + s.frames[index]

		select {
		case <-time.After(time.Millisecond * 100):
			s.writer.Write([]byte(frame))
			index++
		case <-sig:
			s.writer.Write(enableCursor)
			s.writer.Write([]byte("\n"))
		case <-s.stop:
			return
		}
	}
}

func (s *spinner) Stop() {
	s.writer.Write(clearAllLine)
	s.writer.Write(enableCursor)
	s.stop <- true
}

func (s *spinner) Close() {
	s.writer.Write(enableCursor)
}
