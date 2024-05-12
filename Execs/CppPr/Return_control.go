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
		if(!std::is_array<T>::value){
			return 0;
		}
		return 1;
	}`

	// <|>int vals[3] = {1, 2, 3};<|>
	var dArrayCode [3]string = [3]string{
		//body
		`template <typename T, int W>
		void DArrayChecker(T vals[W], string getRetVal[]){
		string Printer = "";
		for(int v = 0; v < W; v++){
			// Only prints types with single value
			if(!TypeHandler(vals[v])){
				cout << vals[v] + " ";
			}
		}
	}`,
		`template <typename T, int W>
		string DArrayChecker(T vals[W]){
		string Printer = "";
		for(int v = 0; v < W; v++){
			// Only prints types with single value
			if(!TypeHandler(vals[v])){
				Printer = Printer +  vals[v] + " ";
			}
		}
		(getRetVal)[1] = "1";
		getRetVal[0] = Printer;
	}`,
		"DArrayChecker"}

	//<type1|>int vals[3][2] = {{1, 2}, {3, 4}, {5, 6}};<type1|>
	var ddArrayCode [3]string = [3]string{
		// body
		`template <typename T, int W, int H>
		void DDArrayChecker(T vals[W][H]){
		string Printer = "";
		for(int v = 0; v < H; v++){
			if(TypeHandler(vals[v])
				){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
					for(int v2 = 0; v2 < H; v2++){
						//*(*(vals2 + v) + v2) = vals[v][v2];
						cout << vals[v][v2] + " ";
					}
					cout << "\n";
				}
			}
		}`,
		`template <typename T, int W, int H>
		string DDArrayChecker(T vals[W][H], string getRetVal[]){
		string Printer = "";
		for(int v = 0; v < W; v++){
			if(TypeHandler(vals[v])
				){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
					for(int v2 = 0; v2 < H; v2++){
						//*(*(vals2 + v) + v2) = vals[v][v2];
						Printer = Printer + vals[v][v2] + " ";
					}
					Printer = Printer + "\n";
				}
			}
			(getRetVal)[1] = "1";
			getRetVal[0] = Printer;
			}`, "DDArrayChecker"}

	var SimpleValue [3]string = [3]string{
		`template<typename T>
		void Printer(T PrintType){
			cout << std::to_string(PrintType);
		}`,
		`template<typename T>
		string Printer(T PrintType, string getRetVal[]){
			(getRetVal)[1] = "1";
			getRetVal[0] = std::to_string(PrintType);
		}`,
		"Printer"}

	//var ChangeArgs map[string]string = map[string]string{"type1": ReturnDType}
	//var CostumCode string = ""
	//var CodeToRun string = MainCode + "\n" + Functions
	FuncRetType = strings.ReplaceAll(FuncRetType, "[]", "[sizeof()]") // if there are empty valuus just fill
	var forDeal string = strings.ReplaceAll(strings.ReplaceAll(strings.Join(strings.Split(FuncRetType, "]["), ", "), "[", ""), "]", "")

	if len(mt.FindMatrix(FuncRetType)) > 0 {
		return []string{fmt.Sprintf("%s<%s>", ddArrayCode[2], forDeal), Functions + "\n\n" + ddArrayCode[compilerType]}
	} else if len(mt.FindList(FuncRetType)) > 0 {
		return []string{fmt.Sprintf("%s<%s>", dArrayCode[2], forDeal), Functions + "\n\n" + dArrayCode[compilerType]}
	} else {
		return []string{fmt.Sprintf("%s<%s>", SimpleValue[2], forDeal), SimpleValue[compilerType]} //Functions + "\n\n" +

	}
	return []string{"-1"}
}
