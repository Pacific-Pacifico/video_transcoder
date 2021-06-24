package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func MakeDir(uuid string) {
	// out, err := exec.Command("mkdir", "-p", "./temp/"+uuid+"/hls").Output()
	os.MkdirAll(filepath.Join("temp", uuid, "hls"), 0700)
}

func Transcode(uuid string, w string, h string) {
	// 640x360 , 854x480 , 1280x720
	cmd := exec.Command("ffmpeg", "-i", "./temp/"+uuid,
		"-profile:v", "baseline",
		"-level", "3.0", "-s", w+"x"+h,
		"-start_number", "0",
		"-hls_time", "10",
		"-hls_list_size", "0", "-f", "hls",
		"./temp/sample/hls/"+h+"_out.m3u8",
	)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("in transcode cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Ran successfully")
}

func MakeMasterFile(fileName string) {
	// 	str := `#EXTM3U
	// #EXT-X-STREAM-INF:BANDWIDTH=375000,RESOLUTION=640x360
	// 360_out.m3u8
	// #EXT-X-STREAM-INF:BANDWIDTH=750000,RESOLUTION=854x480
	// 480_out.m3u8
	// #EXT-X-STREAM-INF:BANDWIDTH=2000000,RESOLUTION=1280x720
	// 480_out.m3u8
	// #EXT-X-STREAM-INF:BANDWIDTH=3500000,RESOLUTION=1920x1080
	// 480_out.m3u8`

	str := `#EXTM3U
#EXT-X-VERSION:3
#EXT-X-STREAM-INF:BANDWIDTH=800000,RESOLUTION=640x360
360p.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=1400000,RESOLUTION=842x480
480p.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=2800000,RESOLUTION=1280x720
720p.m3u8`
	outfile, err := os.Create("./temp/sample/hls/" + fileName + ".m3u8")
	if err != nil {
		panic(err)
	}
	_, err = outfile.Write([]byte(str))
	if err != nil {
		log.Fatal("write error!")
	}
}

func TranscodeAIO(fileName string) {
	// 640x360 , 854x480 , 1280x720

	cmd := exec.Command("ffmpeg", "-y", "-i", "./temp/"+fileName,
		"-vf", "scale=w=640:h=360:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "4", "-hls_playlist_type", "vod", "-b:v", "800k", "-maxrate", "856k", "-bufsize", "1200k", "-b:a", "96k", "-hls_segment_filename", "./temp/sample/hls/360p_%03d.ts", "./temp/sample/hls/360p.m3u8",
		"-vf", "scale=w=842:h=480:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "4", "-hls_playlist_type", "vod", "-b:v", "1400k", "-maxrate", "1498k", "-bufsize", "2100k", "-b:a", "128k", "-hls_segment_filename", "./temp/sample/hls/480p_%03d.ts", "./temp/sample/hls/480p.m3u8",
		"-vf", "scale=w=1280:h=720:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "4", "-hls_playlist_type", "vod", "-b:v", "2800k", "-maxrate", "2996k", "-bufsize", "4200k", "-b:a", "128k", "-hls_segment_filename", "./temp/sample/hls/720p_%03d.ts", "./temp/sample/hls/720p.m3u8")
	//"-vf", "scale=w=1920:h=1080:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "4", "-hls_playlist_type", "vod", "-b:v", "5000k", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k", "-hls_segment_filename", "./temp/sample/hls/1080p_%03d.ts", "./temp/sample/hls/1080p.m3u8")

	// cmd := exec.Command("ffmpeg", "-y", "-i", "./temp/"+fileName,
	// 	"-preset", "medium", "-g", "48", "-sc_threshold", "0",
	// 	"-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1", "-map", "0:0", "-map", "0:1",
	// 	"-s:v:0", "640x360", "-c:v:0", "libx264", "-b:v:0", "365k",
	// 	"-s:v:1", "960x540", "-c:v:1", "libx264", "-b:v:1", "2000k",
	// 	"-s:v:2", "1280x720", "-c:v:2", "libx264", "-b:v:2", "3000k",

	// 	"-c:a", "copy",
	// 	"-var_stream_map", "v:0,a:0 v:1,a:1,v:2,a:2",
	// 	"-master_pl_name", "master.m3u8",
	// 	"-f", "hls",
	// 	"-hls_time", "6",
	// 	"-hls_playlist_type", "vod",
	// 	"-hls_list_size", "0",
	// 	"-hls_segment_filename", "./temp/sample/hls/v%v/fileSequence%d.ts",
	// 	"./temp/sample/hls/v%v/prog_index.m3u8",
	// )

	// cmd := exec.Command("ffmpeg", "-i", "./temp/"+fileName,
	// 	"-filter_complex",
	// 	"[0:v]split=3[v1][v2][v3]; [v1]copy[v1out]; [v2]scale=w=1280:h=720[v2out]; [v3]scale=w=640:h=360[v3out]",
	// 	"-map", "[v1out]", "-c:v:0", "libx264", "-x264-params", "\"nal-hrd=cbr:force-cfr=1\"", "-b:v:0", "5M", "-maxrate:v:0", "5M", "-minrate:v:0", "5M", "-bufsize:v:0", "10M", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-keyint_min", "48",
	// 	"-map", "[v2out]", "-c:v:1", "libx264", "-x264-params", "\"nal-hrd=cbr:force-cfr=1\"", "-b:v:0", "3M", "-maxrate:v:0", "3M", "-minrate:v:0", "3M", "-bufsize:v:0", "3M", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-keyint_min", "48",
	// 	"-map", "[v3out]", "-c:v:2", "libx264", "-x264-params", "\"nal-hrd=cbr:force-cfr=1\"", "-b:v:0", "1M", "-maxrate:v:0", "1M", "-minrate:v:0", "1M", "-bufsize:v:0", "1M", "-preset", "slow", "-g", "48", "-sc_threshold", "0", "-keyint_min", "48",
	// 	"-map", "a:0", "-c:a:0", "aac", "-b:a:0", "96k", "-ac", "2",
	// 	"-map", "a:0", "-c:a:1", "aac", "-b:a:1", "96k", "-ac", "2",
	// 	"-map", "a:0", "-c:a:2", "aac", "-b:a:2", "48k", "-ac", "2",
	// 	"-f", "hls",
	// 	"-hls_time", "2",
	// 	"-hls_playlist_type", "vod",
	// 	"-hls_flags", "independent_segments",
	// 	"-hls_segment_type", "mpegts",
	// 	"-hls_segment_filename", "./temp/sample/hls/stream_%v/data%02d.ts",
	// 	"-master_pl_name", "master.m3u8",
	// 	"-var_stream_map", "v:0,a:0 v:1,a:1 v:2,a:2", "./temp/sample/hls/stream_%v.m3u8",
	// )

	outfile, err := os.Create("./temp/sample/output_" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	errfile, err := os.Create("./temp/sample/err_output" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer errfile.Close()
	cmd.Stdout = outfile
	cmd.Stderr = errfile

	err = cmd.Run()
	if err != nil {
		log.Fatalf("in transcode cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Ran successfully")
}
