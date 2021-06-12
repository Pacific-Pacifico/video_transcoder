package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func GetMediaInfo() {
	if runtime.GOOS == "windows" {
		fmt.Println("can't run on windows")
	}
	// out, err := exec.Command("ls", "-al").Output()

	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format",
		"-show_streams", "./samples/sample.mp4")

	var stdout bytes.Buffer
	// var stderr bytes.Buffer
	cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println("Command executed successfully")
	outstr := stdout.String()
	// errstr := stderr.String()
	fmt.Printf("Output : %s", outstr)
	// fmt.Printf("\nError: %s", errstr)
}

func main() {
	fmt.Println("Hello Wsl")
	GetMediaInfo()
}
