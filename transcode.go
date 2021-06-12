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
	str := `#EXTM3U
#EXT-X-STREAM-INF:BANDWIDTH=375000,RESOLUTION=640x360
360_out.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=750000,RESOLUTION=854x480
480_out.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=2000000,RESOLUTION=1280x720
480_out.m3u8
#EXT-X-STREAM-INF:BANDWIDTH=3500000,RESOLUTION=1920x1080
480_out.m3u8`

	outfile, err := os.Create("./temp/sample/hls/" + fileName + ".m3u8")
	if err != nil {
		panic(err)
	}
	_, err = outfile.Write([]byte(str))
	if err != nil {
		log.Fatal("write error!")
	}
}
