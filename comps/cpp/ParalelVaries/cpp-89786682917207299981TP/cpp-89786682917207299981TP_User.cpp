using namespace std;
#include <iostream> 
#include <string> 
#include <array>



double AddTwoNumbers(int a, int b){ return a + b; }


template<typename T>
		void Printer(T PrintType){
			cout << std::to_string(PrintType);
		}

int main(){
	int a = 7;
int b = 2;
 

Printer<double>(AddTwoNumbers(a=a, b=b));


}