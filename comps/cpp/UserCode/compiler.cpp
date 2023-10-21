#include <iostream>
#include<list>
#include <cstdio>
using namespace std;



// struct ArResturn{
//     int val1;
//     int val2;
//     int array[val1][val2]{};
// }

int main(){
    capper();
    return 0;
}

static int capper(){

}


list<list<int>> Mainer(int apple[1][1], string apples[], int goofer){
    cout << "Apple is tatsy but is not to much";
    int apple[2][4] = {{1, 2, 3, 4, }, {1, 2, 3, 4, }};
    list<list<int>> lister = list<list<int>>{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}; 

    list<list<int>> newList;
    for(int i = 0; i < sizeof(apple); i++ ){
        list<int> islist;
        for(int j = 0; j < sizeof(apple[i]); j++){
            islist.push_back(i);
        }
        newList.push_back(islist);
    }

    return lister;
}

static void Checking(){
    int** Array = 0;
    Array = new int*[10];
    Array[0] = new int[10];
}



static void Sender(){
    int** val = new int*[10];
    int* apples = new int[10]{1, 2, 3, 4, 5, 6, 7, 8, 9};
    val[0] = apples;

    list<list<int>> val2 = list<list<int>>{{1, 2}, {1, 2}}; 
    Diogonal(val2);
}
static string Diogonal(list<list<int>> matrix){
    // printing
    return "";
}

