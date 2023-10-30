package main

import (
	"bytes"
	"example/base/abst"
	"fmt"
	"os/exec"
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
		time.Sleep(time.Second * 1)
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
	for val := <-dd; val != ""; {
		time.Sleep(time.Second * 2)
		// for v := range dd {
		// 	fmt.Println(v)
		// }

		fmt.Println(val)
		fmt.Println(val)
		fmt.Println(<-dd)
		fmt.Println("End Step")
	}
}

func LoopCheckerPar(str chan string) {
	var gottenText string = <-str
	fmt.Println(gottenText)
	str <- gottenText + "Passed"
}

func Inserter(inserter chan string) {
	inserter <- "Apple"
}

func Runner() {
	maker := make(chan string)
	go Inserter(maker)
	fmt.Println(<-maker)
}

func Runner1() {

	var meths []func(chan string) = []func(chan string){LoopCheckerPar, LoopCheckerPar}
	var results []chan string = []chan string{}
	for st, _ := range meths {
		var newval = make(chan string, 1)
		newval <- string(fmt.Sprintf("%d", newval))
		results = append(results, newval)
		go meths[st](newval)
	}

	time.Sleep(time.Second * 2)
	for val := range results {
		newv := <-results[val]
		fmt.Println("Result: ", newv)

	}

	time.Sleep(time.Second * 5)
	// resnew := make(chan string, 2)
	// fmt.Println(<-resnew)
	// resnew <- "Result is empty"
	// resnew := make(chan string)
	// go Running(resnew)
	// var syncer sync.WaitGroup
	// syncer.Add(1)

	// go Checker1(resnew)
	// go Checker2(resnew)

	// time.Sleep(time.Second * 50)

	// go func() {
	// 	go Checker2(resnew)
	// 	syncer.Done()
	// }()

	// time.Sleep(time.Second * 2)
	// fmt.Println(<-resnew)
	// syncer.Wait()
}

func TesterGo(val *string) {
	fmt.Println(val, *val)
	*val = "Apple"
}

func PointerRunner() {
	var newstr string = "Lord is gone"
	go TesterGo(&newstr)
	time.Sleep(time.Second * 2)
	fmt.Println(newstr, &newstr)
}

func BaseChecker() {
	var base abst.BaseConnection = abst.GetBase()
	var taskType = make(chan []string, 2)
	var id = make(chan []string, 2)
	go base.GetTaskByName(map[string]interface{}{"task_name_id": "AppleAreSweet"},
		map[string]chan []string{"taskType": taskType, "GetTaskId": id})

	go base.GetTaskByName(map[string]interface{}{"task_name_id": "AddTwoNumbers"},
		map[string]chan []string{"taskType": taskType, "GetTaskId": id})
	time.Sleep(time.Second * 5)

	var tt = <-taskType
	fmt.Println("Tasks: ", tt, cap(tt))
	fmt.Println("IDS: ", <-id)
}
