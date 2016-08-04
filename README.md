# mkv2Appletv


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
usage: mkv2Appletv [<flags>] <command> [<args> ...]

Convert as efficiently as possible media to AppleTV mp4 format.

Optional flags:
  --help   Show context-sensitive help (also try --help-long and --help-man).
  --debug  Enable debug mode.

Commands:
  help [<command>...]
    Show help.

  show <input>
    Using ffprobe show relavant information about a input file

  suggest <input>
    Show what the suggested output of the transformation would look like.
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

