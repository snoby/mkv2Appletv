package main

import (
	"fmt"
	"github.com/snoby/go-ffprobe"
)

var (
	err error
)

//
// Take a list of audio streams and suggest the best one
// to use as a master to transcode from.
// TODO ADD a config for preferred language
//
// in: a list of streams
// out. an ffproble stream that is the audio master
func masterAudio(fileStreams []*ffprobe.Stream) ffprobe.Stream {

}
