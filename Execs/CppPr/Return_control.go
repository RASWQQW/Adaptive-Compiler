package inputing

// how to clean stack after runtime of file ends in golang

import (
	mt "ep/Execs/methods"
	vals "ep/LevelFuncs"
	"fmt"
	"strings"
)

// HAVE TO FULLY CHECK THAT RETURN VALUE CHECK TO WORK FOR 100% PERCENT

// AND CREATE PRINTER OF RESULT AFTER GETTING OF FUNC  RESULT IF RET VALUE ARRAY OR MATRIX

func main(lang string, valueType *vals.ValType) {
	// var cycleCode map[string]string = map[string]string{
	// 	"cpp":    "",
	// 	"java":   "",
	// 	"csharp": "",
	// }
}

// THE CODE TO DYNAMICALLY  DETERMAINE AND REWRITE IF FUNC CALL NEED TO PRINTED
// OR JUST CALLED TO GET RESULT OR PRINTED TO GET SINGLE RETURN VALUE
func Cpp_control(FuncRetType string, compilerType int) []string { //FuncRetType may differ from  {int } to {flaot, 12, 12 }

	/*var MainCode string = `
		using namespace std;
		#include<iostream>
	`*/

	var Functions string = `
	template <class T> bool TypeHandler(const T &TypeChecker)
	{
		if(!is_array<T>::value){
			return 0;
		}
		return 1;
	}`

	// <|>int vals[3] = {1, 2, 3};<|>
	var dArrayCode [3]string = [3]string{
		"DArrayChecker", // header
		//body
		`template <typename T, int W>
		void DArrayChecker(T vals[W]){
		string Printer = "";
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			// Only prints types with single value
			if(!TypeHandler(vals[v])){
				cout << vals[v] + " ";
			}
		}
	}`,
		`template <typename T, int W>
		string DArrayChecker(T vals[W]){
		string Printer = "";
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			// Only prints types with single value
			if(!TypeHandler(vals[v])){
				Printer = Printer +  vals[v] + " ";
			}
		}
		return Printer;
	}`}

	//<type1|>int vals[3][2] = {{1, 2}, {3, 4}, {5, 6}};<type1|>
	var ddArrayCode [3]string = [3]string{
		"DDArrayChecker", // header
		// body
		`template <typename T, int W, int H>
		void DDArrayChecker(T vals[W][H]){
		string Printer = "";
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			if(TypeHandler(vals[v])
				){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
					for(int v2 = 0; v2 < sizeof vals[v] / sizeof *(*(vals + v) + v2); v2++){
						//*(*(vals2 + v) + v2) = vals[v][v2];
						cout << vals[v][v2] + " ";
					}
					cout << "\n";
				}
			}
		}`,
		`template <typename T, int W, int H>
		string DDArrayChecker(T vals[W][H]){
		string Printer = "";
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			if(TypeHandler(vals[v])
				){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
					for(int v2 = 0; v2 < sizeof vals[v] / sizeof *(*(vals + v) + v2); v2++){
						//*(*(vals2 + v) + v2) = vals[v][v2];
						Printer = Printer + vals[v][v2] + " ";
					}
					Printer = Printer + "\n";
				}
			}
			return Printer;
			}`}

	var SimpleValue [3]string = [3]string{
		"Printer",
		`template<typename T>
		void Printer(T PrintType){
			cout << to_string(PrintType);
		}`,
		`template<typename T>
		string Printer(T PrintType){
			return to_string(PrintType);
		}`}

	//var ChangeArgs map[string]string = map[string]string{"type1": ReturnDType}
	//var CostumCode string = ""
	//var CodeToRun string = MainCode + "\n" + Functions
	FuncRetType = strings.ReplaceAll(FuncRetType, "[]", "[50]") // if there are empty valuus just fill
	var forDeal string = strings.ReplaceAll(strings.ReplaceAll(strings.Join(strings.Split(FuncRetType, "]["), ", "), "[", ""), "]", "")

	if len(mt.FindMatrix(FuncRetType)) > 0 {
		return []string{fmt.Sprintf("%s<%s>", ddArrayCode[compilerType], forDeal), Functions + "\n\n" + ddArrayCode[1]}
	} else if len(mt.FindList(FuncRetType)) > 0 {
		return []string{fmt.Sprintf("%s<%s>", dArrayCode[compilerType], forDeal), Functions + "\n\n" + dArrayCode[1]}
	} else {
		return []string{fmt.Sprintf("%s<%s>", SimpleValue[compilerType], forDeal), Functions + "\n\n" + SimpleValue[1]}

	}
	return []string{"-1"}
}
