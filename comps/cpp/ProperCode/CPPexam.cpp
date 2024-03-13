using namespace std;
#include <iostream>
#include <string>
#include <format>

double AddTwoNumbers(int a, int b){return a + b; }

template <class T>
bool TypeHandler(const T &TypeChecker);

void ArrayPrinter(int vals[3][2], int size, int** vals2){
    int a = 8;
    int b = 10;


    for(int s = 0; s < 3; s = s + 1){
        cout << vals2[s] << endl;
    
        // if(lastPoint + 4 > (int)vals2[s]){
        // 	for(int sv = 0; sv < sizeof vals[s] / sizeof vals[s][sv]; sv = sv + 1){
        // 				cout << vals2[s][sv] << " ";
        // 	}
        // }
        // cout << "\n";
    }
    // cout << AddTwoNumbers(a=a, b=b);
}

void DDArrayChecker(){
    int vals[3][2] = {{1, 2}, {3, 4}, {5, 6}};
    for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
        if(TypeHandler(vals[v])
            //sizeof(vals[v]) / sizeof(string) > 1  || // && sizeof(vals[v]) % sizeof(string) == 0)||
            //sizeof(vals[v]) / sizeof(char) > 1  || //&& sizeof(vals[v]) % sizeof(char) == 0)||
            //sizeof(vals[v]) / sizeof(int) > 1 //&& sizeof(vals[v]) % sizeof(int) == 0)
            ){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
            // when giving value to it also changes the type of inner array
            
            //next lines below no need of usage
            //*(vals2 + v) = new string[sizeof vals[0] / sizeof vals[0]];
            for(int v2 = 0; v2 < sizeof vals[v] / sizeof *(*(vals + v) + v2); v2++){
                //*(*(vals2 + v) + v2) = vals[v][v2];
                cout << vals[v][v2] << " ";
            }
            cout <<  endl;
        }
    }
}

void DArrayChecker(){
    // TYPE OF ARRAY CHANGES EVERYD TIME WHEN  THE STRING CODE IS RAN
    //int vals[3][2] = {{1, 2}, {3, 4}, {5, 6}};
    int vals[3] = {1, 2, 3};

    //THERE WE GOTTA GIVE VALUE TYPE IN STRING TO FURTHER WRIT IN TOP THAT 

    for(int v = 0; v < sizeof vals / sizeof vals[0]; v++){
        // Only prints types with single value
        if(!TypeHandler(vals[v])
           /* &&  (sizeof(vals[v]) == sizeof(int) ||
            sizeof(vals[v]) == sizeof(string) ||
            sizeof(vals[v]) == sizeof(char))*/ 
            ){
            cout << vals[v] << " ";
        }
    }
}

template <class T> bool TypeHandler(const T &TypeChecker)
{
    cout << is_array<T>::value << endl;
    if(!is_array<T>::value){
        return 0;
    }
    return 1;
}

template <typename T, int W, int H> 
void Printer(T value[W][H]){
    cout << "ss" << endl;
}

int main(){
    string sa[1][1]{{"2"}};
    Printer<string, 1, 1>(sa);
    //cout << typeid(type1).name() << ",, " << format("{}", typeid(type2).name()) == typeid(type2).name() << endl;
    return 0;	

    int dd = 12;
    int *val = &dd;
    *(val + 1) = (int)14;
    cout << *val << endl << val << endl << *(val + 1) << endl << (val + 1) << endl;
        

    // string CheckString = "Apple";
    // string Sizer = "hhhh";
    // char  charSizer = 'h';
    // long sizeChecker = 1212121212;
    // cout << "String size checker: " << sizeof(string) << ", " << sizeof(char) << ", "<< sizeof(long) << CheckString.size() << ", "<< sizeof(Sizer) << "---" << sizeof(charSizer) << endl;

    //cout << sizeof(vals[0]) << ", " << vals2 << ", " << *(vals2[0]) << endl;
    //cout << sizeof(vals[0]) << ", " << vals2 << ", " << *(vals2[0]) << endl;
    //cout << sizeof vals2 / sizeof vals2[0]  << ", " << sizeof(vals2[0]) << ", " <<sizeof(vals2[0][0]) << endl;
    
}



