package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("mkv2Appletv", "Convert as efficiently as possible media to AppleTV mp4 format.")
	debug = app.Flag("debug", "Enable debug mode.").Short('d').Bool()
	try   = app.Flag("try", "When set to true only the first 10 seconds of conversion will be done").Short('t').Bool()

	show      = app.Command("show", "Using ffprobe show relavant information about a input file")
	showinput = show.Arg("input", "Location of input File").Required().ExistingFile()

	suggest      = app.Command("suggest", "Show what the suggested output of the transformation would look like.")
	suggestinput = suggest.Arg("input", "Location of input File").Required().ExistingFile()

	convert      = app.Command("convert", "Take input and run ffmpeg to generate an optimal mp4 file")
	convertinput = convert.Arg("input", "Location of input File").Required().ExistingFile()
	outputFile   = convert.Arg("out", "Location of output mp4 File").ExistingDir()

	check = app.Command("check", "log information about ffprobe and ffmpeg")
)

func main() {
	var ()

	kingpin.Version("0.0.3")
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case show.FullCommand():
		showFFprobeInfo((*showinput))

	case suggest.FullCommand():
		suggestConvSettings((*suggestinput))

	case convert.FullCommand():
		convertSource((*convertinput))

	case check.FullCommand():
		checkVersion()
	}
}
