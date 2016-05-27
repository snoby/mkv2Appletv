package main

import (
	"fmt"
	"github.com/snoby/go-ffprobe"
)

var (
	err         error
	fileFormat  *ffprobe.Format
	fileStreams []*ffprobe.Stream
)

func show_ffprobe_info(in string) {
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
		fmt.Println("This stream format doesn't seem to be known, exiting")
		fmt.Println(err)
		return
	}

	fmt.Printf(" Information about file: %s \n", fileFormat.Filename)
	fmt.Printf(" Number of Streams: %d \n", fileFormat.NBStreams)
	fmt.Printf(" File has duration: %s (s)\n", fileFormat.Duration)

	for _, stream := range fileStreams {
		if stream.CodecType == "audio" {
			fmt.Printf("AudioStream: codec: %s channels: %+v bitrate: %s\n", stream.CodecName, stream.Channels, stream.BitRate)
		}
		if stream.CodecType == "video" {
			fmt.Printf("VideoStream: codec: %s Profile: %s   bitrate: %s \n", stream.CodecName, stream.Profile, stream.BitRate)
		}
	}
}
