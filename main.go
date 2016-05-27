package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	app     = kingpin.New("mkv2Appletv", "Convert as efficiently as possible media to AppleTV mp4 format.")
	debug   = app.Flag("debug", "Enable debug mode.").Bool()
	input   = app.Flag("input", "Location of input File").Short('i').Required().ExistingFile()
	outFile = app.Flag("output", "Name of the output mp4 file, this is optional").Short('o').String()

	show = app.Command("show", "Using ffprobe show relavant information about a input file").Default()
	//inputFile = show.Arg("input", "Name of Input File passed to ffprobe ").Required().ExistingFile()

	convert = app.Command("convert", "convert the input file to the best possible output file")
	//inFile  = convert.Arg("input", "Name of Input File used as the source for conversion to an mp4 ").Required().ExistingFile()

)

func main() {
	var ()

	kingpin.Version("0.0.1")
	app.UsageTemplate(kingpin.SeparateOptionalFlagsUsageTemplate)

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// Register user
	case show.FullCommand():
		println("Input file ", (*input))
		show_ffprobe_info((*input))

	case convert.FullCommand():
		//println("Input file", input.Name())

	}
}
