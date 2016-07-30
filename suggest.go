package main

import (
	"errors"
	"fmt"

	"github.com/snoby/go-ffprobe"
)

var ()

//
// Take a list of audio streams and suggest the best one
// to use as a master to transcode from.
// TODO ADD a config for preferred language
//
// in: a list of streams
// out. an ffproble stream that is the audio master
func masterAudio(fileStreams []*ffprobe.Stream) (streamIndex int, err error) {

	/*
	*
	* This may not be the most efficient way to do it, but these files are small
	*
	 */

	// see if there are any ac3 5.1 surround sound streams we can use.
	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			if stream.Channels > 2 {
				if stream.CodecName == "ac3" {
					streamIndex = stream.Index
					return
					//FOUND IT
				}
			}
		}
	} //end of search for master audio that we can just copy over and not transcode.

	//
	// Search for another mutli channel stream we can use.  It will require us to transcode it to ac3
	// but at least it's multi channel.  Also it doesn't matter if we multiple streams like
	// 1 AAC Multi channel & 1 DTS Mutlichannel, at the end we just need one of these.
	//

	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			if stream.Channels > 2 {
				switch stream.CodecName {
				case "aac":
					fallthrough
				case "dts":
					streamIndex = stream.Index
					return

				} // end of switch
			} // end of multi channel
		}
	}

	//
	//
	// At this point there is only 2 channel audio available, AAC is perferred, but if we have another type we
	// can use it in a pinch
	//
	//
	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			if stream.Channels == 2 {
				switch stream.CodecName {
				case "aac":
					fallthrough
				case "ac3":
					streamIndex = stream.Index
					return

				} // end of switch
			} // end of multi channel
		}
	}

	//
	// If we get to this point we are struggling
	//
	//
	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			if stream.Channels == 2 {
				switch stream.CodecName {
				case "dts":
					streamIndex = stream.Index
					return

				} // end of switch
			} // end of multi channel
		}
	}
	err = errors.New("Could not find an audio stream to use as master")
	return
}

func masterVideo(fileStreams []*ffprobe.Stream) (streamIndex int, err error) {
	for _, stream := range fileStreams {
		if stream.CodecType == "video" {
			switch stream.CodecName {
			case "h264":
				streamIndex = stream.Index
				return
			case "h265":
				err = errors.New("Not currently handling h265 streams yet")
				return

			} // end of switch
		}
	}
	err = errors.New("Could not find a video stream to use as master")
	return
}

func suggestConvSettings(in string) {

	println("Input filename:", in)

	h := ffprobe.File(in)

	fileFormat, err = h.Format()
	if err != nil {
		fmt.Println("This file format doesn't seem to be known, exiting")
		fmt.Println(err)
		return
	}

	fileStreams, err = h.Streams()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(" Information about file: %s \n", fileFormat.Filename)
	fmt.Printf(" Number of Streams: %d \n", fileFormat.NBStreams)
	fmt.Printf(" File has duration: %s (s)\n", fileFormat.Duration)

	masterVideoInx, err := masterVideo(fileStreams)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	Videostream := fileStreams[masterVideoInx]

	masterAudioInx, err := masterAudio(fileStreams)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	Audiostream := fileStreams[masterAudioInx]

	fmt.Printf("Master video codec: %s \n", Videostream.CodecName)
	fmt.Printf("Master Audio codec: %s numChannels:%d\n", Audiostream.CodecName, Audiostream.Channels)

}
