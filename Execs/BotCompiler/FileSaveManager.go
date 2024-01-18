// THE THING THAT I CAND HANDLE TO MAKE IT COMPILED IN THE INTERNET
// IS JUST USING PYTHON NEW SERVER THAT CONVERTS CODE FROM ANY TO PYTHON
// THEN JUST COMPILE IT THROUGH USING COMPILE()
package BotCompiler

import (
	lv "ep/LevelFuncs"
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/exp/slices"
)

var PATHG, _ = os.Getwd()

type Profile struct {
	Name        string
	UserCName   string
	ProperCName string
	Lang        string
}

func FileCodeGenerator(TaskPeriod int) (NCode string) {
	NCode = lv.ToString(10000+rand.Int63()+rand.Int63()) + lv.ToString(TaskPeriod) + "TP"
	rr, _ := os.ReadDir(PATHG + "\\comps\\cpp\\ParalelVaries")
	if slices.ContainsFunc(rr, func(name os.DirEntry) bool {
		if fmt.Sprintf("%s", name.Name) == NCode {
			return true
		}
		return false
	}) {
		FileCodeGenerator(TaskPeriod)
	}
	fmt.Println("Generated Code " + NCode)
	return
}

// profile contains prop and user code files
// and spec code to interact with
func SaveProfile(UserCode string, ProperCode string, lang string, tp chan int) *Profile {
	var counter int = <-tp
	counter = counter + 1
	// tp increase id to not get same value in same time
	var getName string = lang + FileCodeGenerator(counter)
	os.Mkdir(PATHG+"\\comps\\cpp\\ParalelVaries\\"+getName, os.ModePerm)
	var nP *Profile = &Profile{Name: getName, UserCName: getName + "_User.cpp", ProperCName: getName + "_Proper.cpp", Lang: lang}
	tp <- counter
	return nP
}

// func CreateProfile()
