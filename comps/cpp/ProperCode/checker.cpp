using namespace std;
#include <iostream> 
#include <string>
#include <cstring>
#include <sstream>

// template <int W>
void PassFunc(int* vals, int sizef){
    cout << "Size is:" << sizeof(*vals) << "sec: "  << sizeof(vals[0]);
    cout  << "val 0: " << vals[0];

    for(int v = 0; v < sizef; v++){
                // Only prints types with single value
            cout  << "val : " << vals[v] << endl;
    }
}

// template <int W, int H>
void PassFunc2D(int **(vals), int s1, int s2){
    cout << "is" << sizeof(*vals) / sizeof(int) << " Size is:" << sizeof(*vals) << "sec: "  << sizeof(*(*(vals + 1)+1));
    cout  << "val 0: " << vals[0]  << endl;
    for(int v = 0; v < s1; v++){
        for(int j = 0; j < s2; j++){
                // Only prints types with single value
                cout  << "val : " << vals[v][j] << endl;
        }
    }
}

int* checkerfunc(){
    return new int[1]{1};    
}

typedef int *getter;

typedef int (*arr)[1][1]; 

void by_pointer(int (*arr)[1][1])
{
    cout << arr << *(*arr) << endl;
    cout << sizeof(arr) << endl << "val: " << *(*(arr)[0]) << endl;
}


int main(){
    // for(int d = 0; d < 10; d++){
    //     float dd = (rand() % 5)* ((float)1 / 21);
    //     cout << dd << endl;
    // }
    // cout << rand() << endl;
    // cout << rand() << endl;s
    // cout << rand() << endl;
    // int* checkerVal;
    // int checker2[1][1] = {{1}}; 
    // int (*getterval)[1][1] = &checker2;
    // int (&gettervald)[1][1] = checker2;

    // cout << (&gettervald) << "  " << (*gettervald) << endl; 
    // cout <<  "checker2 vals: " << checker2 << endl << *checker2 << endl << &checker2 << endl << "END:" << endl;
    // by_pointer(&checker2);
    // checkerVal = checkerfunc();
    // cout << checkerVal << "  " << *checkerVal << endl;

    // cout << typeid(int*).name();
    

    // getter val = new int [10]{11, 12};
    // cout << (val + 1) << "  " << *(val + 1) << endl;



    // typedef getter changing_getter;
    // std::remove_reference<float*>::type originTypeds;
    // cout << "first: " << typeid(originTypeds).name() << endl;

    // getter vald;
    // while(true){
    //     if(std::is_pointer<changing_getter>::value && strcmp(typeid(changing_getter).name(), "Pi") == 0){
    //         typedef std::remove_reference<changing_getter>::type originType;
    //         typedef originType changing_getter;
    //         cout << "1 " << typeid(originType).name() << endl;
    //         cout << "2 " << typeid(changing_getter).name() << endl;
    //     }else{
    //         break;
    //     }
    // }
    // cout << "last: " << typeid(vald).name() << endl;

    // changing_getter val1 = new int [1]{156};
    // cout << "Its type: " << typeid(changing_getter).name() << val1[0] << endl;

    // int passer[] = {1, 2, 3};//{1, 2, 3, 4, 5};
    // int passer2[][5] = {{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}};


    // int **arrs;
    // arrs = new int *[10];
    // arrs[0] = passer2[0];


    // // int **passer3 = passer2;
    // PassFunc
    // // <sizeof(passer) / sizeof(int)>
    // (passer, sizeof(passer) / sizeof(int));    
    
    // cout << "Size bef passing: " << sizeof(passer2) / sizeof(passer2[0]) << " " 
    // << sizeof(passer2[0]) / sizeof(int) << endl;
    
    // int s1 = sizeof(passer2) / sizeof(passer2[0]);
    // int s2 = sizeof(passer2[0]) / sizeof(int);
    // <sizeof(passer2) / sizeof(passer2[0]), sizeof(passer2[0]) / sizeof(int)>
    //(passer2, s1, s2);

    // int **checkerList = new int *[5];
    // checkerList[0] = new int[2]{1, 2};
    // checkerList[1] = new int[2]{1, 2};
    // checkerList[2] = new int[2]{1, 2};
    // checkerList[3] = new int[2]{1, 2};
    // checkerList[4] = new int[2]{1, 2};

        
    // cout << checkerList[0] << "   "<<checkerList[1] << endl;
    // cout << typeid(checkerList[4]).name() << "   "<< checkerList[4][0] << "   "<< checkerList[4][0] <<  "   "<<checkerList[4][0] << endl;
    

    // int **getter = new int *[2];
    // getter[0] = new int[2]{25, 2};
    // getter[1] = new int[2]{27, 2};
    // cout << typeid(getter[0]).name() << "   "<< getter[0][0] << "   "<< getter[1][0]<< endl;

    // int **getterd2 = new int *[4];
    // getterd2[0] = new int[2]{654, 4322};
    // getterd2[1] = new int[2]{};

    // cout << typeid(getterd2[0]).name() << "   " <<  getterd2[1][1] << " " << to_string(getterd2[1][0]).size() << "   "<< getterd2[1][0]<< endl;

    // int val = -1427761392;

    std::stringstream trs;
    trs << "apple";
    trs.str("");
    trs << "apple2";
    cout << trs.str();
    return 0;
}   