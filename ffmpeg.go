package main

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func checkVersion() {
	out, err := exec.Command("ffmpeg", "-version").CombinedOutput()
	if err != nil {
		// Not found
		fmt.Println(err)
		fmt.Println("Unable to find ffmpeg in your path\n")
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
		} else {
			fmt.Println("Found the correct version of ffmpeg\n")
		}
		// found and right version
		fmt.Printf("%s\n", out)
	}

	// Now check the ffprobe version
	probeout, err := exec.Command("ffprobe", "-version").CombinedOutput()
	if err != nil {
		// Not found
		fmt.Println(err)
		fmt.Println("Unable to find ffmpeg in your path\n")
	} else {
		// Find out what version of ffmpeg that is installed
		version := regexp.MustCompile(`ffprobe.version.\d\S*`)
		re := regexp.MustCompile("ffprobe.version 3")

		bFound := re.MatchString(string(probeout[:30]))
		if bFound == false {
			//found but wrong version
			fmt.Println("Requires ffprobe >= 3.x.x I found version %s", version.MatchString(string(probeout[:30])))
			fmt.Println("Go here to download binaries for your machine: https://ffmpeg.org/download.html")
			fmt.Println("I would recommend the static compiled version")
			err = errors.New("Wrong version of ffmpeg found ")
		} else {
			fmt.Println("Found the correct version of ffprobe\n")
		}
		fmt.Printf("%s", probeout)
	}
}
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

//
// // add stats: -stats
// func callFFmpeg(cmd Cmd) (string, error) {
//
// 	//fmt.Printf("%s", args)
// 	// var err error
// 	//cmd := exec.Command("ffmpeg", args)
// 	fmt.Printf("%v", cmd)
//
// 	fmt.Println("About to start\n")
//
// 	err = cmd.Start()
// 	if err != nil {
// 		fmt.Printf("Error starting program: %s\n", err)
// 	}
// 	fmt.Println("About to wait\n")
// 	err = cmd.Wait()
// 	if err != nil {
// 		fmt.Printf("Error waiting for  program: %s\n", err)
// 	}
//
// 	return "Success", err
// }
