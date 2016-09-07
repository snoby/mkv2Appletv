# mkv2Appletv
[![CircleCI](https://circleci.com/gh/snoby/mkv2Appletv/tree/master.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/snoby/mkv2Appletv/tree/master)

Drone CI

[![Build Status](http://drone.mattsnoby.com/api/badges/snoby/mkv2Appletv/status.svg )] (http://drone.mattsnoby.com/snoby/mkv2Appletv)


## Why
I've always wanted to find a program like ffmepgtools that would automatically convert any video files I downloaded to the format / codecs that would play on the appleTV.  Specifically look at the input streams available and output the best quality mp4 videos as fast as possible.

I also wanted to learn the golang programming language.  So in my freetime I wrote this program.  I'm sure it has bugs.

### Requirements
* at minimum output an mp4 that has h264 video and aac audio.
* If multichannel audio is available, output h264 video, aac audio 2 channel, ac3 5.1 channel audio.
* Do the conversion as fast as possible.
* Don't change video if possible just copy
* Don't change audio if possible just copy
* Have a good help menu
* Show information about video file
* package with ffmpeg ( in a docker container )

### Usage
```
usage: mkv2Appletv --input=INPUT [<flags>] <command> [<args> ...]

Convert as efficiently as possible media to AppleTV mp4 format.

Flags:
      --help           Show context-sensitive help (also try --help-long and --help-man).
  -d, --debug          Enable debug mode.
  -t, --try            When set to true only the first 10 seconds of conversion will be done
  -i, --input=INPUT    Location of input File
  -o, --output=OUTPUT  Location of output mp4 File (not required, if not set output file name will be input filename +.mp4)

Commands:
  help [<command>...]
    Show help.


  show
    Using ffprobe show relavant information about a input file


  suggest
    Show what the suggested output of the transformation would look like.


  convert
    Take input and run ffmpeg to generate an optimal mp4 file


  check
    log information about ffprobe and ffmpeg
```
#### show command
The dump below shows that we have a video file that has an h264 video stream with 2 audio's an aac 2 channel and an ac3 2 channel, it has subtitles and stream metadata along with artwork.

```
./mkv2Appletv show ~/Public/public/civil_war.mp4
Input file  /Users/snoby/Public/public/civil_war.mp4
 Information about file: /Users/snoby/Public/public/civil_war.mp4
 Number of Streams: 5
 File has duration: 8836.931000 (s)
VideoStream: codec: h264 Profile: High   bitrate: 826355
AudioStream: codec: aac channels: 2 bitrate: 128253
AudioStream: codec: ac3 channels: 2 bitrate: 128000
Data Stream: mov_text type: subtitle
VideoStream: codec: mjpeg Profile:    bitrate:
```
#### suggest command
```
./mkv2Appletv suggest ~/Public/public/civil_war.mp4
Input file  /Users/snoby/Public/public/civil_war.mp4
Input filename: /Users/snoby/Public/public/civil_war.mp4
 Information about file: /Users/snoby/Public/public/civil_war.mp4
 Number of Streams: 5
 File has duration: 8836.931000 (s)
Master Video codec: h264
Master Audio codec: aac numChannels:2
***
Primary Video Stream (h264)
Primary Audio Stream (aac)
----------Planned Output-----------------
Video Stream (h264)  operation [copy]
Primary  Audio Stream (aac)  operation [copy]
Second  Audio Stream  (ac3)  operation [none]
```
#### Convert Command
```
~/work/myprojects/godevelopment/src/github.com/snoby/mkv2Appletv$ ./mkv2Appletv -t -i ~/Public/public/JL.mkv -o ~/Public/public/test.mp4 convert
Input filename: /Users/snoby/Public/public/JL.mkv
 Information about file: /Users/snoby/Public/public/JL.mkv
 Number of Streams: 2
 File has duration: 4494.688000 (s)
Master Video codec: h264
Master Audio codec: ac3 numChannels:6
***
Primary Video Stream (h264)
Primary Audio Stream (ac3)
----------Planned Output-----------------
Video Stream (h264)  operation [copy]
Primary  Audio Stream (aac)  operation [convert]
Second  Audio Stream  (ac3)  operation [copy]
[-hide_banner -y -i /Users/snoby/Public/public/JL.mkv -map_metadata 0:g -t 00:00:10 -map 0:0 -c:v copy -map 0:1 -c:a:0 aac -b:a:0 256k -map 0:1 -c:a:1 copy /Users/snoby/Public/public/test.mp4]About to start


&{/opt/local/bin/ffmpeg [ffmpeg -hide_banner -y -i /Users/snoby/Public/public/JL.mkv -map_metadata 0:g -t 00:00:10 -map 0:0 -c:v copy -map 0:1 -c:a:0 aac -b:a:0 256k -map 0:1 -c:a:1 copy /Users/snoby/Public/public/test.mp4] []  <nil> <nil> <nil> [] <nil> <nil> <nil> <nil> false [] [] [] [] <nil>}
Input #0, matroska,webm, from '/Users/snoby/Public/public/JL.mkv':
  Metadata:
    encoder         : libebml v1.3.0 + libmatroska v1.4.0
    creation_time   : 2014-05-29 05:41:22
  Duration: 01:14:54.69, start: 0.000000, bitrate: 3914 kb/s
    Chapter #0:0: start 0.000000, end 233.233000
    Metadata:
      title           : Prologue and Credits
    Chapter #0:1: start 233.233000, end 406.406000
    Metadata:
      title           : War Over
    Chapter #0:2: start 406.406000, end 620.620000
    Metadata:
      title           : Like Vigilantes
    Chapter #0:3: start 620.620000, end 818.150000
    Metadata:
      title           : Live for the Moment
    Chapter #0:4: start 818.150000, end 1066.398000
    Metadata:
      title           : Las Vegas Rescue
    Chapter #0:5: start 1066.398000, end 1260.259000
    Metadata:
      title           : Damaged Goods
    Chapter #0:6: start 1260.259000, end 1423.922000
    Metadata:
      title           : You Will All Be Judged
    Chapter #0:7: start 1423.922000, end 1724.222000
    Metadata:
      title           : The Real Ferris Aircraft
    Chapter #0:8: start 1724.222000, end 2005.503000
    Metadata:
      title           : It's All True
    Chapter #0:9: start 2005.503000, end 2248.746000
    Metadata:
      title           : Mission to Mars
    Chapter #0:10: start 2248.746000, end 2538.369000
    Metadata:
      title           : Space Rescue
    Chapter #0:11: start 2538.369000, end 2808.305000
    Metadata:
      title           : Lantern Legacy
    Chapter #0:12: start 2808.305000, end 2949.613000
    Metadata:
      title           : It's Coming
    Chapter #0:13: start 2949.613000, end 3123.453000
    Metadata:
      title           : Biblical Scale
    Chapter #0:14: start 3123.453000, end 3395.225000
    Metadata:
      title           : We Work Together
    Chapter #0:15: start 3395.225000, end 3680.009000
    Metadata:
      title           : Dinosaurs Unleashed
    Chapter #0:16: start 3680.009000, end 3864.861000
    Metadata:
      title           : Fight On
    Chapter #0:17: start 3864.861000, end 4025.521000
    Metadata:
      title           : Godspeed
    Chapter #0:18: start 4025.521000, end 4195.858000
    Metadata:
      title           : The Ring and the Will
    Chapter #0:19: start 4195.858000, end 4310.806000
    Metadata:
      title           : New Frontier
    Chapter #0:20: start 4310.806000, end 4494.688000
    Metadata:
      title           : End Credits
    Stream #0:0: Video: h264 (High), yuv420p, 1920x1080 [SAR 1:1 DAR 16:9], 23.98 fps, 23.98 tbr, 1k tbn, 47.95 tbc (default)
    Metadata:
      title           : ETRG
    Stream #0:1(eng): Audio: ac3, 48000 Hz, 5.1(side), fltp, 640 kb/s (default)
    Metadata:
      title           : ETRG
[mp4 @ 0x7f8cb401e400] track 2: codec frame size is not set
Output #0, mp4, to '/Users/snoby/Public/public/test.mp4':
  Metadata:
    creation_time   : 2014-05-29 05:41:22
    encoder         : Lavf57.25.100
    Chapter #0:0: start 0.000000, end 10.000000
    Metadata:
      title           : Prologue and Credits
    Stream #0:0: Video: h264 ([33][0][0][0] / 0x0021), yuv420p, 1920x1080 [SAR 1:1 DAR 16:9], q=2-31, 23.98 fps, 23.98 tbr, 16k tbn, 1k tbc (default)
    Metadata:
      title           : ETRG
    Stream #0:1(eng): Audio: aac (LC) ([64][0][0][0] / 0x0040), 48000 Hz, 5.1(side), fltp, 256 kb/s (default)
    Metadata:
      title           : ETRG
      encoder         : Lavc57.24.102 aac
    Stream #0:2(eng): Audio: ac3 ([165][0][0][0] / 0x00A5), 48000 Hz, 5.1(side), 640 kb/s (default)
    Metadata:
      title           : ETRG
Stream mapping:
  Stream #0:0 -> #0:0 (copy)
  Stream #0:1 -> #0:1 (ac3 (native) -> aac (native))
  Stream #0:1 -> #0:2 (copy)
Press [q] to stop, [?] for help
frame=  242 fps=159 q=-1.0 Lsize=    5394kB time=00:00:10.01 bitrate=4411.7kbits/s speed=6.58x
video:4285kB audio:1096kB subtitle:0kB other streams:0kB global headers:0kB muxing overhead: 0.243517%
[aac @ 0x7f8cb403d000] Qavg: 153.092
```

#### Demo of the app
Here is an example of the app running with the -t option which only encodes the first 10 seconds of the video


![](https://cloud.githubusercontent.com/assets/724760/17499799/d603494a-5d9d-11e6-8fac-1be62bd62d9f.gif)

