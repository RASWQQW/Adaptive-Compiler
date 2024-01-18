package cpp

import (
	execs "ep/Execs/BotCompiler"
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"GOPROJECT/Execs/CppPr/FileSaveManager"
)

type apple = interface {
	~[]string
}

func main() {
	// here  occurs only commands related to cpp and its compilation
	fmt.Println("Compliation start")
}

func Runner(currentGivenPath string, returns chan []string, profEl *execs.Profile, filename string) []string {
	var currentPath string = ""
	if len(currentGivenPath) >= 1 {
		currentPath = currentGivenPath
	} else {
		path, _ := os.Getwd()
		currentPath = path + "\\comps\\cpp\\ParalelVaries"

	}
	var profPath string = currentPath + "\\" + profEl.Name
	//var filePath string = currentPath + "\\" + profEl.Name + "\\" + filename

	fmt.Println("PATH: " + profPath)
	err := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; g++ "+filename, profPath)).Run()
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