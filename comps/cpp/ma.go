package cpp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// here  occurs only commands related to cpp and its compilation
	fmt.Println("Compliation start")
}

func Runner(currentGivenPath string, returns chan []string) []string {
	var currentPath string = ""
	if len(currentGivenPath) >= 1 {
		currentPath = currentGivenPath
	} else {
		path, _ := os.Getwd()
		currentPath = path + "\\comps\\cpp\\ProperCode"

	}
	exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; g++ compiler.cpp", currentPath)).Run()
	var result2 = exec.Command("powershell.exe", fmt.Sprintf("cd %s; cmd /C a.exe ", currentPath))

	var errs strings.Builder
	var stdread strings.Builder

	result2.Stderr = &errs
	result2.Stdout = &stdread
	result2.Run()

	// fmt.Printf("Output of: %s", strings.Trim(string(outpp[:]), " "))
	// fmt.Println("Result at end:", stdread.String(), errs.String())
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
