package main

import (
	"errors"
	"fmt"

	"github.com/snoby/go-ffprobe"
)

var (
	media = new(Convert)
)

// Convert Contains the control block of what we will do with the
// conversion process
type Convert struct {
	inFile            string
	outFile           string
	masterVideoStream *ffprobe.Stream
	masterAudioStream *ffprobe.Stream
	aacAudioStream    *ffprobe.Stream
	outVideo          string
	outAudio0         string
	outAudio1         string
}

func checkforAACsecondaryAudio(fileStreams []*ffprobe.Stream) (streamIndex int, err error) {
	// see if there are any ac3 5.1 surround sound streams we can use.
	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			if stream.Channels == 2 {
				if stream.CodecName == "aac" {
					streamIndex = stream.Index
					fmt.Printf("Found a 2 channel aac stream")
					return
					//FOUND IT
				}
			}
		}
	} //end of search for master audio that we can just copy over and not transcode.
	err = errors.New("Could not find secondary aac audio")
	return -1, err
}

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
				case "truehd":
					fallthrough
				case "aac":
					fallthrough
				case "dca":
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
				case "truehd":
					fallthrough
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

//
//
// Find the masterVideo stream.  At this point it's usually just
// the only video stream, but may need to add code here for the
// situation where we have more than one video
//
//
func masterVideo(fileStreams []*ffprobe.Stream) (streamIndex int, err error) {
	for _, stream := range fileStreams {
		if stream.CodecType == "video" {
			streamIndex := stream.Index
			return streamIndex, nil
		}
	}
	err = errors.New("Could not find a video stream to use as master")
	return
}

func (media *Convert) print() {
	fmt.Println("***")
	fmt.Printf("Primary Video Stream (%s)\n", media.masterVideoStream.CodecName)
	fmt.Printf("Primary Audio Stream (%s)\n", media.masterAudioStream.CodecName)
	fmt.Printf("----------Planned Output-----------------\n")
	fmt.Printf("Video Stream (%s)  operation [%s]\n", "h264", media.outVideo)
	fmt.Printf("Primary  Audio Stream (%s)  operation [%s]\n", "aac", media.outAudio0)
	fmt.Printf("Second  Audio Stream  (%s)  operation [%s]\n", "ac3", media.outAudio1)

}

//
// Adding a method
//
func (media *Convert) setupAudioConversion(fileStreams []*ffprobe.Stream) {

	if media.masterAudioStream.Channels == 2 {
		//
		// Can't surround sound with this.
		//
		switch media.masterAudioStream.CodecName {
		case "aac":
			media.outAudio0 = "copy"
			media.outAudio1 = "none"
		case "ac3":
			fmt.Println("Found ac3 Master audio...")
			stream, err := checkforAACsecondaryAudio(fileStreams)
			if err != nil {
				media.outAudio0 = "convert"
				media.outAudio1 = "none"
			} else {
				//Not really sure how this can happen
				media.aacAudioStream = fileStreams[stream]
				media.outAudio0 = "copy"
			}
			media.outAudio1 = "none"
		case "dts":
			media.outAudio0 = "convert"
			media.outAudio1 = "none"
		default:
			fmt.Printf("Not sure what to do with this codec: %s", media.masterAudioStream.CodecName)

		} // end of switch
	} else {
		// The Master Audio has surround sound
		switch media.masterAudioStream.CodecName {
		case "ac3":
			stream, err := checkforAACsecondaryAudio(fileStreams)
			if err != nil {
				// This means we didn't find an aac alternate
				media.outAudio0 = "convert"
				media.outAudio1 = "copy"
			} else {
				// we found the aac 2 channel stream
				media.aacAudioStream = fileStreams[stream]
				media.outAudio0 = "copy"
				media.outAudio1 = "copy"
			}

		case "aac":
			fallthrough
		case "truehd":
			fallthrough
		case "dca":
			fallthrough
		case "dts":
			media.outAudio0 = "convert"
			media.outAudio1 = "convert"
		default:
			fmt.Printf("Not sure what to do with this codec: %s", media.masterAudioStream.CodecName)

		} // end of switch

	} //end of if channels > 2
}

//
// Adding a method
//
func (media *Convert) setupVideoConversion() {

	//
	// Can't surround sound with this.
	//
	switch media.masterVideoStream.CodecName {
	case "h264":
		media.outVideo = "copy"
	default:
		media.outVideo = "convert"

	} // end of switch

}

//
//
// Find the best audio and video streams in the container
// Then make a suggestion on what to do with the streams
// to make a compatible mp4 file for the appleTV.
//
//
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

	media.inFile = in

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

	media.masterVideoStream = Videostream
	media.masterAudioStream = Audiostream

	fmt.Printf("Master Video codec: %s \n", Videostream.CodecName)
	fmt.Printf("Master Audio codec: %s numChannels:%d\n", Audiostream.CodecName, Audiostream.Channels)

	media.setupAudioConversion(fileStreams)
	media.setupVideoConversion()
	media.print()

}
