package inputing

// how to clean stack after runtime of file ends in golang

import (
	mt "ep/Execs/methods"
	vals "ep/LevelFuncs"
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
func cpp_control(ReturnDType string) []string {

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
	var dArrayCode []string = []string{
		"DArrayChecker", // header
		//body
		`template <typename T, int W>
	void DArrayChecker(T vals[W]){
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			// Only prints types with single value
			if(!TypeHandler(vals[v])){
				cout << vals[v] << " ";
			}
		}
	}
	`}

	//<type1|>int vals[3][2] = {{1, 2}, {3, 4}, {5, 6}};<type1|>
	var ddArrayCode []string = []string{
		"DDArrayChecker", // header
		// body
		`template <typename T, int W, int H>
	void DDArrayChecker(T vals[W][H]){
		for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
			if(TypeHandler(vals[v])
				){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
				for(int v2 = 0; v2 < sizeof vals[v] / sizeof *(*(vals + v) + v2); v2++){
					//*(*(vals2 + v) + v2) = vals[v][v2];
					cout << vals[v][v2] << " ";
				}
				cout <<  endl;
			}
		}
	}
	`}

	//var ChangeArgs map[string]string = map[string]string{"type1": ReturnDType}
	//var CostumCode string = ""
	//var CodeToRun string = MainCode + "\n" + Functions
	if len(mt.FindMatrix(ReturnDType)) > 0 {
		return []string{ddArrayCode[0], Functions + ddArrayCode[1]}
	} else if len(mt.FindList(ReturnDType)) > 0 {
		return []string{dArrayCode[0], Functions + dArrayCode[1]}
	}
	return []string{"-1"}
}
