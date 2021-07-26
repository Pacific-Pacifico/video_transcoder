package main

import (
	"fmt"
	"time"
)

func consume() {
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	LogIfError(true, err, "Failed to set QoS")
	msgs, err := ch.Consume(
		"TranscodeQueue",
		"",
		false, //make autoack true if task hangs or something wrong happens
		false,
		false,
		false,
		nil,
	)
	LogIfError(true, err, "error consuming task")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved msg:%s\n", d.Body)
			start := time.Now()
			TranscodeAIO(string(d.Body))
			fmt.Println("Processing complete! Took ", time.Since(start))
			d.Ack(false)
			fmt.Println("~~waiting for task~~")
		}
	}()

	fmt.Println("~~waiting for task~~")
	<-forever
}
