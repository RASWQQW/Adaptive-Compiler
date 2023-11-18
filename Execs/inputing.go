package inputing

import (
	cppComp "ep/Execs/CppPr"
	obj "ep/Execs/obj"
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

func cppCompiler(cont *obj.Container) string {
	return cppComp.Compiler(cont)
}

func javaCompiler(cont *obj.Container) string {
	return ""
}

func pythonCompiler(cont *obj.Container) string {
	return ""
}

func csharpCompiler(cont *obj.Container) string {
	return ""
}

// here goes little runner like main Func
func Main(lang string, code string, topic string) string {
	// var Mainer Compiller = Compiller{"Find out that is the most characterized", topic, code}
	var cont = obj.SetValsOfContainer(code, topic)
	if lang == "cpp" {
		return cppCompiler(cont)
	} else if lang == "java" {
		return javaCompiler(cont)
	} else if lang == "csharp" {
		return csharpCompiler(cont)
	} else if lang == "python" {
		return pythonCompiler(cont)
	}
	panic("Here is no thing that can be called")
}
