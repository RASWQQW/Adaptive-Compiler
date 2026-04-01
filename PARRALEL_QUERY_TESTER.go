package main

import (
	csd "ep/Execs/BotCompiler"
	"ep/LevelFuncs"
	"fmt"
	"sync"
	"time"
)

func Tester() {
	fmt.Println("Testing is starting")
}

type LockerObj struct {
	ism       sync.Mutex
	isValsAdd []int
}

func (ct *LockerObj) Printer(digit int) {
	ct.ism.Lock()
	time.Sleep(time.Second / 2)
	defer ct.ism.Unlock()
	fmt.Println("SAN: ", digit)
	ct.isValsAdd = append(ct.isValsAdd, digit)
}

func PrinterV2(ct *LockerObj, digit int) {
	ct.ism.Lock()
	time.Sleep(time.Second / 2)
	defer ct.ism.Unlock()
	fmt.Println("SAN: ", digit)
	ct.isValsAdd = append(ct.isValsAdd, digit)
}

func Counter() {
	// var syncer sync.Mutex
	// fmt.Println(len(qq))
	nobject := LockerObj{isValsAdd: []int{}}

	for _, ss := range [6]int{1, 2, 3, 4, 5, 6} {
		//go nobject.Printer(ss);
		go PrinterV2(&nobject, ss)
	}
	for {
		if len(nobject.isValsAdd) >= 6 {
			fmt.Println("It  stopped...")
			break
		}
	}
}

func CheckerFunc() {
	// rtValGetcheck := make(chan []string)
	rtValGetcheck := make(chan string, 3)

	go func() {
		for v := 0; v < cap(rtValGetcheck); v++ {
			select {
			case dd := <-rtValGetcheck:
				fmt.Println(dd)
				// default:
				// 	continue
			}

		}
	}()

	rtValGetcheck <- "12"
	rtValGetcheck <- "13"
	rtValGetcheck <- "14"
	rtValGetcheck <- "15"

}

func SiteTester() {
	//var cd sync.Mutex

	nobject := LevelFuncs.MutexLog{}

	rtValGet1 := make(chan []string, 1)
	rtValGet2 := make(chan []string, 1)
	fmt.Println(cap(rtValGet1), cap(rtValGet2))

	go csd.WebSocketRunner(map[string]any{"code": "", "lang": "cpp"}, rtValGet1, &nobject)
	go csd.WebSocketRunner(map[string]any{"code": "", "lang": "cpp"}, rtValGet2, &nobject)
	fmt.Println("res1: ", <-rtValGet1)
	fmt.Println("res2: ", <-rtValGet2)
}
