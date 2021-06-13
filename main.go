package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func GetMediaInfo(res string, fileName string) {
	if runtime.GOOS == "windows" {
		fmt.Println("can't run on windows")
	}
	var width, height string
	if res == "240" {
		width = "352"
		height = "240"
	} else if res == "360" {
		width = "480"
		height = "360"
	} else if res == "480" {
		width = "854"
		height = "480"
	} else if res == "720" {
		width = "1280"
		height = "720"
	}
	// out, err := exec.Command("ls", "-al").Output()

	// cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format",
	// 	"-show_streams", "./samples/sample.mp4")

	cmd := exec.Command("ffmpeg", "-i", "./samples/"+fileName, "-vf", "scale="+width+":"+height, "./samples/"+res+"_"+fileName)

	// var stdout bytes.Buffer
	// var stderr bytes.Buffer
	outfile, err := os.Create("./samples/output+" + res + "_" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	errfile, err := os.Create("./samples/err_output" + res + ".txt")
	if err != nil {
		panic(err)
	}
	defer errfile.Close()
	cmd.Stdout = outfile
	cmd.Stderr = errfile
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	err = cmd.Run()
	fmt.Println("Executing...")
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Command executed successfully")
	// outstr := stdout.String()
	// errstr := stderr.String()
	// fmt.Printf("Output : %s", outstr)
	// fmt.Printf("\nError: %s", errstr)
}

func main() {
	// fmt.Println("Hello Wsl")
	//GetMediaInfo("480", "sample.mkv")
	MakeDir("sample")
	// Transcode("sample.mp4", "640", "360")
	// Transcode("sample.mp4", "854", "480")
	// Transcode("sample.mp4", "1280", "720")
	// MakeMasterFile("master_file")
	TranscodeAIO("pass.mkv")

}
