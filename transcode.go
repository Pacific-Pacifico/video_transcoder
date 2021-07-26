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

func MakeMasterFile(fileName string) {
	str := `#EXTM3U
#EXT-X-VERSION:3
#EXT-X-STREAM-INF:BANDWIDTH=800000,RESOLUTION=640x360
360p.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=1800000,RESOLUTION=960x540
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
	// The larger the GOP size,
	// the more efficient the compression and the less bandwidth you will need.
	cmd := exec.Command("ffmpeg", "-y", "-i", "./temp/"+fileName,
		"-vf", "scale=w=640:h=360:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "23", "-sc_threshold", "0", "-g", "60", "-keyint_min", "60", "-hls_time", "10", "-hls_playlist_type", "vod", "-b:v", "800k", "-maxrate", "856k", "-bufsize", "1200k", "-b:a", "96k", "-hls_segment_filename", "./temp/sample/hls/360p_%03d.ts", "./temp/sample/hls/360p.m3u8",
		// "-vf", "scale=w=842:h=480:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "23", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "10", "-hls_playlist_type", "vod", "-b:v", "1400k", "-maxrate", "1498k", "-bufsize", "2100k", "-b:a", "128k", "-hls_segment_filename", "./temp/sample/hls/480p_%03d.ts", "./temp/sample/hls/480p.m3u8",

		"-vf", "scale=w=960:h=540:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "23", "-sc_threshold", "0", "-g", "60", "-keyint_min", "60", "-hls_time", "10", "-hls_playlist_type", "vod", "-b:v", "1800k", "-maxrate", "2000k", "-bufsize", "3000k", "-b:a", "128k", "-hls_segment_filename", "./temp/sample/hls/540p_%03d.ts", "./temp/sample/hls/540p.m3u8",

		"-vf", "scale=w=1280:h=720:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "23", "-sc_threshold", "0", "-g", "60", "-keyint_min", "60", "-hls_time", "10", "-hls_playlist_type", "vod", "-b:v", "2800k", "-maxrate", "2996k", "-bufsize", "4200k", "-b:a", "128k", "-hls_segment_filename", "./temp/sample/hls/720p_%03d.ts", "./temp/sample/hls/720p.m3u8")
	//"-vf", "scale=w=1920:h=1080:force_original_aspect_ratio=decrease", "-c:a", "aac", "-ar", "48000", "-c:v", "h264", "-profile:v", "main", "-crf", "23", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-hls_time", "10", "-hls_playlist_type", "vod", "-b:v", "5000k", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k", "-hls_segment_filename", "./temp/sample/hls/1080p_%03d.ts", "./temp/sample/hls/1080p.m3u8")

	// outfile, err := os.Create("./temp/sample/output_" + fileName + ".txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer outfile.Close()

	errfile, err := os.Create("./temp/sample/" + fileName + "_err_output.txt")
	if err != nil {
		panic(err)
	}
	defer errfile.Close()
	// cmd.Stdout = outfile
	cmd.Stderr = errfile

	err = cmd.Run()
	if err != nil {
		log.Fatalf("in transcode cmd.Run() failed with %s\n", err)
	}
	MakeMasterFile("master_file")
	fmt.Println("Ran successfully")
}
