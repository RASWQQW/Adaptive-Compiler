package inputing

import (
	cppComp "example/inputing/CppPr"
	obj "example/inputing/obj"
	"os"
	"strings"
)

var Globalpath, _ = os.Getwd()
var longstring []string = strings.Split(Globalpath, "\\")
var GlobalPath = strings.Join(longstring[0:len(longstring)-2], "\\") + "\\comps"

// here happends getting a task number and its code and language
type Languages struct {
	java   bool
	cpp    bool
	csharp bool
	python bool
}

type CompilerRunner interface {
	cppCompiler() string
	javaCompiler() string
	pythonCompiler() string
	csharpCompiler() string
}

func cppCompiler(cont obj.Container) string {
	return cppComp.Compiler(cmp)
}

func javaCompiler(cont obj.Container) string {
	return ""
}

func pythonCompiler(cont obj.Container) string {
	return ""
}

func csharpCompiler(cont obj.Container) string {
	return ""
}

// here goes little runner like main Func
func Main(lang string, code string, topic string) string {
	// var Mainer Compiller = Compiller{"Find out that is the most characterized", topic, code}
	var cont = obj.SetVals(code, topic)
	if lang == "cpp" {
		return cont.compiler.cppCompiler()
	} else if lang == "java" {
		return cont.compiler.javaCompiler()
	} else if lang == "csharp" {
		return cont.compiler.csharpCompiler()
	} else if lang == "python" {
		return cont.compiler.pythonCompiler()
	}
	panic("Here is no thing that can be called")
}
