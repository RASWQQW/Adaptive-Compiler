#include <stdio.h>
#include <windows.h> // for EXCEPTION_ACCESS_VIOLATION
#include <iostream>
#include <string>

template <class T> bool TypeHandler(const T &TypeChecker)
	{
		if(!std::is_array<T>::value){
			return 0;
		}
		return 1;
	}

class A{};

int main()
{
    std::cout << std::is_array<A>::value << std::endl;
    std::cout << std::to_string(12) << std::endl;
}