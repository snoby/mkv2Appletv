package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("mkv2Appletv", "Convert as efficiently as possible media to AppleTV mp4 format.")
	debug = app.Flag("debug", "Enable debug mode.").Bool()

	show      = app.Command("show", "Using ffprobe show relavant information about a input file")
	showinput = show.Arg("input", "Location of input File").Required().ExistingFile()

	suggest      = app.Command("suggest", "Show what the suggested output of the transformation would look like.")
	suggestinput = suggest.Arg("input", "Location of input File").Required().ExistingFile()
)

func main() {
	var ()

	kingpin.Version("0.0.2")
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)
	//app.UsageTemplate(kingpin.DefaultUsageTemplate)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// Register user
	case show.FullCommand():
		println("Input file ", (*showinput))
		showFFprobeInfo((*showinput))

	case suggest.FullCommand():
		println("Input file ", (*suggestinput))
		suggestConvSettings((*suggestinput))

	}
}
