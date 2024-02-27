package cpp

import (
	"ep/LevelFuncs"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	// "ep/Execs/execs"
	//"GOPROJECT/Execs/CppPr/FileSaveManager"
)

type apple = interface {
	~[]string
}

func main() {
	// here  occurs only commands related to cpp and its compilation
	fmt.Println("Compliation start")
}

// func Runner(CommonPath string, returns chan []string, ProfileObj *execs.Profile, filename string) []string {
func Runner(InnerVals map[string]any, returns chan []string) []string {
	var currentPath string = ""
	var profPath string = ""
	if len(LevelFuncs.ToString(InnerVals["CommonPath"])) >= 1 {
		currentPath = LevelFuncs.ToString(InnerVals["CommonPath"])
	} else {
		path, _ := os.Getwd()
		currentPath = path + "\\comps\\cpp\\ParalelVaries"

	}
	var ProfVal = reflect.ValueOf(InnerVals["ProfileObj"])
	if ProfVal.Kind().String() == "Profile" {
		profPath = currentPath + "\\" + ProfVal.FieldByName("Name").String() //InnerVals["profEl"].Name
		//var filePath string = currentPath + "\\" + profEl.Name + "\\" + filename
	}

	fmt.Println("PATH: " + profPath)
	err := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; g++ "+LevelFuncs.ToString(InnerVals["filename"]), profPath)).Run()
	fmt.Println("OUT ERR: ", err)
	var result2 = exec.Command("powershell.exe", fmt.Sprintf("cd %s; cmd /C a.exe ", profPath))

	var errs strings.Builder
	var stdread strings.Builder

	result2.Stderr = &errs
	result2.Stdout = &stdread
	result2.Run()

	// fmt.Printf("Output of: %s", strings.Trim(string(outpp[:]), " "))
	// fmt.Println("Result at end:", stdread.String(), errs.String())
	fmt.Println("Result: "+stdread.String(), errs.String())
	returns <- []string{stdread.String(), errs.String()}
	return []string{stdread.String(), errs.String()}
}

func GetAllInnerValues(checkPath string) string {
	var result = exec.Command("powershell.exe", fmt.Sprintf("cd %s; pwd; dir", checkPath))
	outcom, _ := result.CombinedOutput()
	result.Run()

	// var readerByter []byte
	// var output, rr = outpp.Read(readerByter)
	// if rr != nil {
	// 	panic(rr)
	// }
	// fmt.Printf("Output of2: %s", output, string(readerByter[:]))

	return string(outcom[:])
}
