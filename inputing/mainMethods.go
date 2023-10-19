package inputing

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func Replacer(forrep []string, value string) string {
	// forrep+=[""]
	for _, dd := range forrep {
		value = strings.ReplaceAll(value, dd, "")
	}
	return value
}

func TypeGuesser(text string) string {
	var retval interface{}
	if text == "int" {
		retval = rand.Intn(355)
	} else if text == "double" {
		retval = rand.Float64()
	}
	return fmt.Sprintf("%s", retval)
}

func GetParamType(comp string) string {
	var GetTypeString string = regexp.MustCompile("[a-z]+").FindString(comp)
	if len(GetTypeString) > 0 {
		return GetTypeString
	}
	return ""
}

func GetLen(comp string) []string {
	var GetListLen = regexp.MustCompile("[\\[\\d*\\]]+").FindAllString(comp, -1)
	if len(GetListLen) > 0 {
		for id := range GetListLen {
			GetListLen[id] = regexp.MustCompile("\\d*").FindString(GetListLen[id])
		}
		return GetListLen
	}
	return []string{}
}

func GetFullType(ValueName string, paramtype string, paramlen []string) string {
	var params string
	for _, val := range paramlen {
		params = params + fmt.Sprintf("[%s]", val)
	}
	return paramtype + fmt.Sprintf(" %s%s", ValueName, params)
}

func getRandom(TypeString string) string {
	return fmt.Sprintf("%s, ", TypeGuesser(TypeString))
}

func RandValSetting(paramtype string, paramlen string, istype string) string {
	var IntLen int = 0
	var RandVal string = "{"
	if len(paramlen) > 0 {
		IntLen, _ = strconv.Atoi(paramlen)
	} else {
		var toint int = 15
		if istype == "matrix" {
			toint = 8
		} else if istype == "list" {
			toint = 15
		}
		IntLen = rand.Intn(toint)
	}
	for i := 0; i < IntLen; i++ {
		RandVal = RandVal + getRandom(paramtype)
	}
	return RandVal + "}"
}

func FindMatrix(slice string) string {
	return string(regexp.MustCompile("([A-Za-z]+[\\[\\]]{4:}").FindString(slice))
}

func FindList(slice string) string {
	return string(regexp.MustCompile("[A-Za-z]+[\\[]{1}[\\]]{1}").FindString(slice))
}
