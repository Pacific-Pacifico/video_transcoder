package main

import (
	"fmt"
	"strings"
)

func ArgsSplitter(str string) (ArgsStringsList []string) {
	ArgsStringsList = strings.Split(str, " ")
	for i := 0; i < len(ArgsStringsList); i++ {
		// fmt.Println(index, item)
		if ArgsStringsList[i][0] == '"' {
			ArgsStringsList[i] = "\\" + ArgsStringsList[i] + "\\"
		}
		ArgsStringsList[i] = "\"" + ArgsStringsList[i] + "\"" + ","
		fmt.Println(ArgsStringsList[i])
	}
	return ArgsStringsList
}

// func main() {
// 	fmt.Println(ArgsSplitter("ffmpeg -i ./samples/sample.mp4 -vf scale=480:360 ./samples/output2.mp4"))
// 	// fmt.Println(ArgsSplitter("ffprobe -v quiet -print_format json -show_format -show_streams \"lolwut.mp4\" > \"lolwut.mp4.json\""))
// fmt.Println(ArgsSplitter("ffmpeg -i some_fun_video_name.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls ./media/some_fun_video_name/hls/360_out.m3u8"))
// }
