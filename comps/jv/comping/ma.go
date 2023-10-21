package comping

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Java compiler entrance")
}

func PathCheck() {
	val, err1 := os.Getwd()
	if err1 == nil {
		fmt.Print("Current run path: ", val)
	}

	var path string = "comps/jv"
	res, pathtesting := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; ls; pwd", path)).Output()
	fmt.Println(string(res), pathtesting)

}

func Runner(st string) []string {

	var path string = "comps"
	var shell string = "powershell.exe"
	var precompiler strings.Builder
	var wholebyter []byte

	precheck := exec.Command(shell, "-c", fmt.Sprintf("cd %s; javac jv/comp.java", path))
	precheck.Stderr = &precompiler
	precheck.Stderr.Write(wholebyter)
	precheck.Run()

	var prepstring string = precompiler.String()

	fmt.Println("Pre check error: ", prepstring, len(prepstring))
	fmt.Println("Pre check error2: ", string(wholebyter), len(wholebyter))

	if !strings.Contains(precompiler.String(), "error") || len(prepstring) == 0 {
		// check
		// here goes if here is not build or syntax error but runtime may be cause at least
		var resulting = exec.Command(shell, "-c", fmt.Sprintf("cd %s; java jv/comp", path))

		var out, runtimerror strings.Builder
		resulting.Stdout = &out
		resulting.Stderr = &runtimerror

		resulting.Start()
		return []string{out.String(), runtimerror.String()}
	}
	return []string{"", prepstring}
}
