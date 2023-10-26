package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"golang.org/x/exp/slices"
)

func Delayers(givenText string, timed int, checkChan chan string) {
	// dd, _ := time.ParseDuration("2s")
	time.Sleep(time.Duration(timed * int(math.Pow(10, 9))))
	fmt.Println(givenText, "Chan: ", checkChan)
	checkChan <- givenText
}

func RunnerDo() {

	nch := make(chan string)
	nch <- "Start chan"
	go Delayers("Apple", 2, nch)
	go Delayers("Peach", 1, nch)

	fmt.Println("End chan: ", nch)

}

func TestingSlice() {
	var dd = func(v int, v2 int) bool {
		if v > v2 {
			return false
		}
		return true
	}
	// var faanger []string = []string{"a1", "a2", "a3"}
	var faanger []int = []int{1, 2, 3, 9, 10}

	var sl = slices.CompactFunc(faanger, dd)
	fmt.Println(sl)
	for dd, dd2 := range faanger[1:] {
		fmt.Println(dd2, dd)
	}

}

func SaveFileTest() {
	newTab, _ := os.Getwd()
	var newString = newTab + "\\comps\\cpp\\UserCode\\Testing.cpp"

	wr, err := os.Create(newString)
	defer wr.Close()

	if err != nil {
		panic(err)
	}
	wr.WriteString(`int main(){std::cout << "I am god mf";}`)
}
