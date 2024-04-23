package main

import (
	"bytes"
	obc "ep/Execs/obj"
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
	time.Sleep(time.Second * 1)
	// var gottenText string = <-str
	// fmt.Println(gottenText)
	str <- "Passed"
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
		var newval = make(chan string)
		// newval <- string(fmt.Sprintf("%d", newval))
		results = append(results, newval)
		go meths[st](newval)
	}

	// time.Sleep(time.Second * 2)
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
	// var base abst.BaseConnection = abst.GetBase()
	var taskType = make(chan []string, 2)
	var id = make(chan []string, 2)
	// go base.GetTaskByName(map[string]interface{}{"task_name_id": "AppleAreSweet"},
	// 	map[string]chan []string{"taskType": taskType, "GetTaskId": id})

	// go base.GetTaskByName(map[string]interface{}{"task_name_id": "AddTwoNumbers"},
	// 	map[string]chan []string{"taskType": taskType, "GetTaskId": id})
	time.Sleep(time.Second * 5)

	var tt = <-taskType
	fmt.Println("Tasks: ", tt, cap(tt))
	fmt.Println("IDS: ", <-id)
}

func ListChecker(wTime time.Duration, val *[]int, val2 []int, val3 map[string]string) {
	time.Sleep(time.Second * wTime)
	val2 = []int{12, 12, 12}
	val3 = map[string]string{"Apple": "sweet"}
	fmt.Println([]interface{}{"In go", val, val2, val3}...)
}

func NewRunner() {
	var val *[]int = &[]int{1, 2, 3}
	var val2 []int = []int{4}
	var val3 map[string]string = map[string]string{"apple": "app"}

	go ListChecker(1, val, val2, val3)
	go ListChecker(2, val, val2, val3)

	for {
		fmt.Println([]interface{}{val, val2, val3}...)
		time.Sleep(time.Millisecond * 500)
	}
}

func Runner2(cc *obc.Career) {
	fmt.Println(cc.OUTS["apple"])
	cc.OUTS["apple"] = "bad"
}

func Runner22(cc *obc.Career) {
	fmt.Println(cc.OUTS["apple"])
	cc.OUTS["apple"] = "damn"
}

func Testings() {
	var nval *obc.Career = &obc.Career{map[string]interface{}{"apple": "good"}, map[string]any{"apple1": "ss"}}
	go Runner2(nval)
	fmt.Println("Out1 ", nval, nval.OUTS)
	time.Sleep(time.Second * 1)
	go Runner22(nval)
	time.Sleep(time.Second * 1)
	fmt.Println("Out2 ", nval, nval.OUTS)
}

func TestingAsChan(val *obc.Career) {
	var iis bool = true
	for {
		for key, _ := range val.OUTS {
			if key == "Apple" {
				fmt.Printf("Before called method retrieved,%s, %s", key, val.OUTS[key])
				iis = false
				break
			}
		}
		if iis == false {
			break
		}
		time.Sleep(time.Millisecond * 500)

	}
}
func GivingVal(val1 *obc.Career) {
	val1.OUTS["Apple"] = "Sweet"
}

func Checking() {
	var nval *obc.Career = &obc.Career{map[string]any{}, map[string]any{}}
	go TestingAsChan(nval)
	time.Sleep(time.Second * 2)
	go GivingVal(nval)
	time.Sleep(time.Second * 10)
}
