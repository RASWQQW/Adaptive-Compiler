package cpp

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	// here  occurs only commands related to cpp and its compilation
	fmt.Println("Compliation start")
}

func Runner(currentGivenPath string) []string {
	var currentPath string = ""
	if len(currentGivenPath) < 1 {
		currentPath = currentGivenPath
	} else {
		currentPath = "comps/cpp/ProperCode"
	}

	exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; g++ compiler.cpp", currentPath))
	var result = exec.Command("powershell.exe", "c-", fmt.Sprintf("cd %s; a.exe", currentPath))
	var errs strings.Builder

	result.Stderr = &errs
	out, _ := result.Output()

	return []string{string(out), errs.String()}
}
