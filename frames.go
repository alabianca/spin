package spin

var dotFrames = []string{
	"⠋",
	"⠙",
	"⠸",
	"⠼",
	"⠦",
	"⠇",
}

var dotFrames2 = []string{
	"⣷",
	"⣽",
	"⣻",
	"⢿",
	"⡿",
	"⣟",
	"⣯",
	"⣷",
}

var lineFrames = []string{
	"-",
	"\\",
	"|",
}

func getFrames(st SpinType) []string {
	var frames []string

	switch st {
	case Dots:
		frames = dotFrames
	case Lines:
		frames = lineFrames
	case Dots2:
		frames = dotFrames2
	}

	return frames
}
