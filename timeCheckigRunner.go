package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

func Running(res chan string) {

	command := exec.Command(`powershell.exe", "Time-Sleep -s 3; echo "Apple is big"`)
	buf := new(bytes.Buffer)
	d, _ := command.StdoutPipe()

	command.Start()

	buf.ReadFrom(d)
	res <- buf.String()
}

func Checker1(vow chan string) {
	for {
		vow <- "text1"
		vow <- "text2"
		// time.Sleep(time.Second)
		// if dd2 := <-vow; dd2 == "Text" {
		// 	vow <- "Text2"
		// } else {
		// 	vow <- "Text"
		// }
	}
}
func Checker2(dd chan string) {
	for {
		time.Sleep(time.Second * 2)
		for v := range dd {
			fmt.Println(v)
		}
	}
}

func Runner() {
	resnew := make(chan string, 2)
	fmt.Println(<-resnew)
	// resnew <- "Result is empty"
	// resnew := make(chan string)
	// go Running(resnew)
	var syncer sync.WaitGroup
	syncer.Add(1)
	go func() {
		go Checker1(resnew)
		go Checker2(resnew)
		syncer.Done()
	}()

	// time.Sleep(time.Second * 2)
	// fmt.Println(<-resnew)
	syncer.Wait()
}
