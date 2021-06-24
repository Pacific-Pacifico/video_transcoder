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
// 	// 	// fmt.Println(ArgsSplitter("ffmpeg -i ./samples/sample.mp4 -vf scale=480:360 ./samples/output2.mp4"))
// 	// 	// fmt.Println(ArgsSplitter("ffprobe -v quiet -print_format json -show_format -show_streams \"lolwut.mp4\" > \"lolwut.mp4.json\""))
// 	// 	// fmt.Println(ArgsSplitter("ffmpeg -i some_fun_video_name.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls ./media/some_fun_video_name/hls/360_out.m3u8"))
// 	// 	// fmt.Println(ArgsSplitter("ffmpeg -y -i sintel_trailer-1080p.mp4 -preset slow -g 48 -sc_threshold 0 -map 0:0 -map 0:1 -map 0:0 -map 0:1 -s:v:0 640x360 -c:v:0 libx264 -b:v:0 365k -s:v:1 960x540 -c:v:1 libx264 -b:v:1 2000k -c:a copy -var_stream_map \"v:0,a:0 v:1,a:1\" -master_pl_name master.m3u8 -f hls -hls_time 6 -hls_list_size 0 -hls_segment_filename \"v%v/fileSequence%d.ts\" v%v/prog_index.m3u8"))
// 	// 	fmt.Println(ArgsSplitter("ffmpeg -i brooklynsfinest_clip_1080p.mp4 -filter_complex \"[0:v]split=3[v1][v2][v3]; [v1]copy[v1out]; [v2]scale=w=1280:h=720[v2out]; [v3]scale=w=640:h=360[v3out]\" -map [v1out] -c:v:0 libx264 -x264-params \"nal-hrd=cbr:force-cfr=1\" -b:v:0 5M -maxrate:v:0 5M -minrate:v:0 5M -bufsize:v:0 10M -preset slow -g 48 -sc_threshold 0 -keyint_min 48 -map [v2out] -c:v:1 libx264 -x264-params \"nal-hrd=cbr:force-cfr=1\" -b:v:0 3M -maxrate:v:0 3M -minrate:v:0 3M -bufsize:v:0 3M -preset slow -g 48 -sc_threshold 0 -keyint_min 48 -map [v3out] -c:v:2 libx264 -x264-params \"nal-hrd=cbr:force-cfr=1\" -b:v:0 1M -maxrate:v:0 1M -minrate:v:0 1M -bufsize:v:0 1M -preset slow -g 48 -sc_threshold 0 -keyint_min 48 -map a:0 -c:a:0 aac -b:a:0 96k -ac 2 -map a:0 -c:a:1 aac -b:a:1 96k -ac 2 -map a:0 -c:a:2 aac -b:a:2 48k -ac 2 -f hls -hls_time 2 -hls_playlist_type vod -hls_flags independent_segments -hls_segment_type mpegts -hls_segment_filename stream_%v/data%02d.ts -master_pl_name master.m3u8 -var_stream_map \"v:0,a:0 v:1,a:1 v:2,a:2\" stream_%v.m3u8"))
// 	// fmt.Println(ArgsSplitter(`ffmpeg -hide_banner -y -i beach.mkv -vf scale=w=640:h=360:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 800k -maxrate 856k -bufsize 1200k -b:a 96k -hls_segment_filename beach/360p_%03d.ts beach/360p.m3u8 -vf scale=w=842:h=480:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 1400k -maxrate 1498k -bufsize 2100k -b:a 128k -hls_segment_filename beach/480p_%03d.ts beach/480p.m3u8 -vf scale=w=1280:h=720:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 2800k -maxrate 2996k -bufsize 4200k -b:a 128k -hls_segment_filename beach/720p_%03d.ts beach/720p.m3u8 -vf scale=w=1920:h=1080:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename beach/1080p_%03d.ts beach/1080p.m3u8`))

// 	fmt.Println(ArgsSplitter(`ffmpeg -y -i ./hls.mkv \
// 	-vf scale=w=640:h=360:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 800k -maxrate 856k -bufsize 1200k -b:a 96k -hls_segment_filename ./hls/360p_%03d.ts ./hls/360p.m3u8 \
// 	-vf scale=w=842:h=480:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 1400k -maxrate 1498k -bufsize 2100k -b:a 128k -hls_segment_filename ./hls/480p_%03d.ts ./hls/480p.m3u8 \
// 	-vf scale=w=1280:h=720:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 2800k -maxrate 2996k -bufsize 4200k -b:a 128k -hls_segment_filename ./hls/720p_%03d.ts ./hls/720p.m3u8 \
// 	-vf scale=w=1920:h=1080:force_original_aspect_ratio=decrease -c:a aac -ar 48000 -c:v h264 -profile:v main -crf 20 -sc_threshold 0 -g 48 -keyint_min 48 -hls_time 4 -hls_playlist_type vod -b:v 5000k -maxrate 5350k -bufsize 7500k -b:a 192k -hls_segment_filename ./hls/1080p_%03d.ts ./hls/1080p.m3u8`))
// }
