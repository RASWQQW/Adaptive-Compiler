using namespace std;
#include <iostream> 
#include <unistd.h>
#include <thread>
#include<chrono>
#include <vector>
#include <future>
#include <string>
#include <cstdarg>
#include <random>
#include <format>


int** AddTwoNumbers(int a, int b){
	
	//return a + b;
	int **add;
	return add;

}


void Representers(int pp[],  string val){
    cout << "Start: " << val << endl <<"pointer: " << pp;
    std::this_thread::sleep_for(std::chrono::milliseconds(3175));
    (pp)[0] = 1;
    cout << "pp val was given to " << pp << endl << pp[0] << endl;
}


void checkerdd(){        
		std::chrono::steady_clock::time_point begin_t = std::chrono::steady_clock::now();
		int pointer[1] = {0};
        auto runner = [&pointer]() {
        Representers(pointer, "");};
        std::thread ddx(runner);
        //ddx.join();
        // float checker = ((place + rand() % 3)) * ((float) 1 / 21);
        cout << "FUNC EXCEEDED CHECK1 pointer: " << pointer << endl;
        int st = 0;        
        float fin_time_to = 3.16;
        while(st < 3160/15.6436){
            std::this_thread::sleep_for(std::chrono::microseconds(500));
            std::this_thread::sleep_for(std::chrono::microseconds(500));
			
			st = st + 1;
            if(pointer[0] == 1){
                cout << "broke" << endl;
                break;
            }        
        }
		
		std::chrono::time_point end_t = std::chrono::steady_clock::now();
		string diff = to_string((float)(std::chrono::duration_cast<std::chrono::milliseconds>(end_t-begin_t).count()));
		cout <<  "END SEC: " << st << " get pointer: " << pointer[0]  << "Time diff: " << diff<< endl;
}


int main(){
	// int a = 207;
	// int b = 338;
 	// cout << AddTwoNumbers(a=a, b=b);
	// int  dd[1][2] = {{1, 2}};
	// cout << dd << endl << *(*(dd)) << endl << *(*(dd)+1) << endl << dd+1<< endl << dd + 2;
	checkerdd();
	return 0;
}