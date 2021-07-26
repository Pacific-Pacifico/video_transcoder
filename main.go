package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	MakeDir("sample")
	TranscodeAIO("sample.mp4")
	fmt.Println(time.Since(start))
}
