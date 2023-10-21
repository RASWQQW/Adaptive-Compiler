package py

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	Runner("Apple is so sweat")
}

func Runner(command string) string {
	command = "echo Apple"
	var output, _ = exec.Command("powershell.exe", "-c", "pwd").Output()

	if strings.Contains(strings.ToLower(string(output)), "goporject") {
		var executor = exec.Command("powershell.exe", "-c", "cd comps/py; py -3.11 -m compiller")
		var out, err strings.Builder
		executor.Stderr = &err
		executor.Stdout = &out

		executor.Run()

		fmt.Println("Error:", err.String())
		fmt.Printf("Output: %s", out.String())
	}

	return ""
}

type Writer int

func (*Writer) Convertor() {

}
