package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func checkFFmpegVersion() error {

	out, err := exec.Command("ffmpeg", "-version").CombinedOutput()
	if err != nil {
		// Not found
		fmt.Println(err)
	} else {
		// Find out what version of ffmpeg that is installed
		version := regexp.MustCompile(`ffmpeg.version.\d\S*`)
		re := regexp.MustCompile("ffmpeg.version 3")

		bFound := re.MatchString(string(out[:30]))
		if bFound == false {
			//found but wrong version
			fmt.Println("Requires ffmpeg >= 3.x.x I found version %s", version.MatchString(string(out[:30])))
			fmt.Println("Go here to download binaries for your machine: https://ffmpeg.org/download.html")
			fmt.Println("I would recommend the static compiled version")
			err = errors.New("Wrong version of ffmpeg found ")
		}
		// found and right version
	}
	return err
}

// add stats: -stats
func callFFmpeg(args string) (string, error) {
	cmd := "ffmpeg"
	//args := []string{"-resize", "50%", "foo.jpg", "foo.half.jpg"}
	if err := exec.Command(cmd, args).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return "Success", err
}
