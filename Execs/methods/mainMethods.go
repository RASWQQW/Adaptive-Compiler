package methods

import (
	"fmt"
	"math/rand"
	"regexp"
	re "regexp"
	"strconv"
	"strings"
)

func Aligner(UserCode string) (string, string) {

	// Work when working only with semicolon based languages
	// The main purpose not just keep proper probels but also indentation and lininges
	command := re.MustCompile("\\s{1,}")

	var newCode string = UserCode
	for strings.Index(UserCode, "  ") != -1 {
		space := strings.Index(UserCode, "  ")
		UserCode = UserCode[:space+1] + UserCode[space+2:]
	}
	UserCode = strings.Trim(UserCode, " ")
	var valuesPre []string = command.FindAllString(newCode, -1)
	var indexing [][]int = command.FindAllStringIndex(newCode, -1)
	var indexes [][]int = [][]int{}
	var values []string = []string{}
	for v := 0; v < len(indexing); v++ {
		// fmt.Println(indexing[v][1] - indexing[v][0])
		if (indexing[v][1] - indexing[v][0]) > 1 {
			indexes = append(indexes, indexing[v])
			values = append(values, valuesPre[v])
		}
	}

	fmt.Println(indexes)
	for ind := range indexes {
		fmt.Println("String: ", len(values[ind]), indexes[ind])
		var addVal int = 1
		var AllMarks string = ")(}{;"
		var Index []int = indexes[ind]
		if Index[0] != 0 && !(Index[1] > len(newCode)) && strings.Contains(AllMarks, string(newCode[Index[0]-1])) ||
			strings.Contains(AllMarks, string(newCode[indexes[ind][1]])) {
			addVal = 0
		}
		newCode = newCode[:indexes[ind][0]+addVal] + newCode[indexes[ind][1]:]
		// fmt.Println(newCode[:indexes[ind][0]+1], "++", newCode[indexes[ind][1]:])
		for vv := ind + 1; vv < len(indexes); vv++ {
			fmt.Println("Before:", indexes[vv])
			for ind3 := range indexes[vv] {
				indexes[vv][ind3] = indexes[vv][ind3] - (len(values[ind]) - addVal)
			}
			fmt.Println("After:", indexes[vv])
		}
	}

	fmt.Println("Aligned code: ", newCode)
	return newCode, UserCode
}

func Replacer(forrep []string, value string) string {
	// forrep+=[""]
	for _, dd := range forrep {
		value = strings.ReplaceAll(value, dd, "")
	}
	return value
}

func TypeGuesser(text string) any {
	var retval interface{}
	if text == "int" {
		retval = rand.Intn(355)
	} else if text == "double" {
		retval = rand.Float64()
	}
	return retval
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
