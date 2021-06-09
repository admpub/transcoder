package ffmpeg

import "github.com/admpub/transcoder"

// Config ...
type Config struct {
	FfmpegBinPath   string
	FfprobeBinPath  string
	ProgressEnabled bool
	Verbose         bool
	Env             []string
	Dir             string
	OnMetadata      func(transcoder.Metadata) error
}
