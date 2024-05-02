using namespace std;
#include <iostream>



double AddTwoNumbers(int a, int b){
	return a + b;
	}



	template <class T> bool TypeHandler(const T &TypeChecker)
	{
		if(!is_array<T>::value){
			return 0;
		}
		return 1;
	}

template<typename T>
		void Printer(T PrintType){
			return to_string(PrintType);
		}

int main(){
	int a = 4;
int b = 4;
 

Printer(AddTwoNumbers(a=a, b=b));


}