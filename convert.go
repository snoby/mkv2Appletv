package main

import (
	"errors"
	"fmt"
)

var (
	ffmpegCmd = new(ffmpegOut)
)

type ffmpegOut struct {
	outFile string
	header  string
	Video   string
	Audio0  string
	Audio1  string
}

func (buff *ffmpegOut) genVideoConversion() error {

	switch media.outVideo {
	case "copy":
		buff.Video = fmt.Sprintf("-map %d:0 -c:v copy ", media.masterVideoStream.Index)
	case "convert":
		buff.Video = fmt.Sprintf("-map %d:0 -c:v libx264 -preset slow -crf 20 -profile:v high -level 4.0 ", media.masterVideoStream.Index)
	default:
		err := errors.New("unknown or not set Video settings\n")
		return err
	} // end of switch
	return err
}

func (buff *ffmpegOut) genAudioConversion() error {

	//
	// This is the best case senerio.  We have an aac and an ac3 stream
	//
	if media.outAudio0 == "copy" && media.outAudio1 == "copy" {
		buff.Audio0 = fmt.Sprintf("-map 0:%d -c:a:0 copy ", media.aacAudioStream.Index)
		buff.Audio1 = fmt.Sprintf("-map 0:%d -c:a:1 copy ", media.masterAudioStream.Index)
		return err
	} else if media.outAudio0 == "convert" && media.outAudio1 == "convert" {
		//Handle situation where all codecs need to be generated
		//TODO Handle audio1 situation
		if media.masterAudioStream.Channels > 2 {
			// Output a 2 channel aac stream from the master audio which is a multichannel audio
			// the asplit filter takes the input and splits it into 2 dual streams.  one called 2ch and another called 6ch
			// The 2ch is fed into the pan filter and the output is placed into the "aac" pad
			//
			buff.Audio0 = fmt.Sprintf("-filter_complex \"[0:%d]asplit[2ch][6ch];[2ch]pan=stereo|FL=FC+0.6FL+0.2BL|FR=FC+0.6FR+0.2BR[aac]\" ", media.masterAudioStream.Index)
			// Now append the output pad mappings [aac] and [6ch]
			// if the 6ch is NOT ac3 we will have to transcode it.
			buff.Audio0 = buff.Audio0 + fmt.Sprintf(" -map [aac] -map [6ch] -c:a:0 aac -c:a:1 ac3 ")
		} else {
			// if the master audio is NOT aac and is only 2 channel
			buff.Audio0 = fmt.Sprintf("-map 0:%d -c:a:0 aac -b:a 256k ", media.masterAudioStream.Index)
		}
		return err
	}

	switch media.outAudio0 {
	case "copy":
		// This is a situation where the only audio available is an aac 2 channel
		buff.Audio0 = fmt.Sprintf("-map 0:%d -c:a:0 copy ", media.masterAudioStream.Index)

	case "convert":
		// we need to figure out if this is multichannel
		buff.Audio0 = fmt.Sprintf("-map 0:%d -c:a:0 aac -b:a 256k ", media.masterAudioStream.Index)

	} // end of switch

	switch media.outAudio1 {
	case "copy":
		// This situation is when we have an ac3 codec and we have to generate the aac codec
		buff.Audio1 = fmt.Sprintf("-map 0:%d -c:a:1 copy ", media.masterAudioStream.Index)
	case "convert":
		//this is currently handled in case where numb audio streams > 2
		// we need to figure out if this is multichannel
		//		fmt.Println("Not converting the second audio yet, this is not yet implemented")
		//		buff.Audio1 = fmt.Sprintf("-c:a:1 copy ")

	} // end of switch
	return err
}
func (buff *ffmpegOut) setupHeader() {

	// copy all metadata from the source
	metadata := fmt.Sprintf(" -map_metadata 0:g ")
	metadata = metadata + fmt.Sprintf(" -t 00:00:10 ")
	buff.header = metadata

}

func convertSource(in string) {

	suggestConvSettings(in)
	// Media object is now setup

	ffmpegCmd.outFile = fmt.Sprintf("%s.mp4", in)
	ffmpegCmd.setupHeader()
	ffmpegCmd.genVideoConversion()
	ffmpegCmd.genAudioConversion()
	// Format String to send to ffmpeg
	out := "-i " + in + ffmpegCmd.header + ffmpegCmd.Video + ffmpegCmd.Audio0 + ffmpegCmd.Audio1 + ffmpegCmd.outFile
	fmt.Println("Sending to ffmpeg:")
	fmt.Println(out)
	err := checkFFmpegVersion()
	if err != nil {
		fmt.Println("Not sending commands to ffmpeg because: %s\n", err)
		return
	}

	// Call ffmpeg here

}
