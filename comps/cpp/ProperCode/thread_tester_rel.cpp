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
#include <sstream>

//https://en.cppreference.com/w/cpp/thread/thread/native_handle
auto ffer = [](int timesd){
    cout << "CheckerCode" << endl;
    //sleep(timesd);
    cout <<  "CheckerCode2 >>" << timesd << endl;
};


// THERE IT HAVE TO BE REWRITTEN TO THE FUNCTIONS
// BETTER CAN TRY MACROS TO MAKE IT MORE FITTING
// ALL OF THEM FUNC_PREPARE VALUES AND STATIC
typedef string to_run_ret;
typedef to_run_ret (*run_code_temp)(string); //  there we need just pass parameters of func when preparing
double TIME_LIMIT = 3.12;   

string Printer(string PrintType){
    return PrintType;
}

void Representers(string getRetVal[], to_run_ret val){
    //cout << "Start: " << val << endl <<"pointer: " << pp;
    std::this_thread::sleep_for(std::chrono::milliseconds(3110 + stoi(val)));
    (getRetVal)[1] = "1";
    //cout << "pp val was given to " << pp << endl << pp[0] << endl;
    getRetVal[0] = val + "\n" + "pointer: " + getRetVal[1];
}


// FUNC PREPARE TEMPLATE
// COMPLETE TRANSMIITER
void transmitFunc(run_code_temp to_run, 
                  string(*Representer)(string[], to_run_ret), 
                  vector<vector<string>> *arr, 
                  string FuncId,
                  int place,
                  string parameter /*THERE FUNC WILL BE PREPARED AS TEMPLATE TO OWN PARAMS*/){

    std::chrono::steady_clock::time_point begin_t = std::chrono::steady_clock::now();
    std::stringstream s;

    try{
        std::cout << "FUNC EXCEEDED CHECKbefore" << endl;
        string pointer[2] = {0};
        auto runner = [Representer, to_run, parameter, &pointer]() {
            Representer(pointer, to_run(parameter));};

        std::thread ddx(runner);
        ddx.detach();
        // ddx.join();
        // float checker = ((place + rand() % 3)) * ((float) 1 / 21);
        //cout << "FUNC EXCEEDED CHECK1 pointer: " << pointer << endl;
        int st = 0;        
        while(st < TIME_LIMIT*1000/15.6436){
            std::this_thread::sleep_for(std::chrono::milliseconds(1));
            st = st + 1;
            if(pointer[1] == "1"){
                //cout << "broke" << endl;
                break;
            }        
        }
        std::chrono::time_point end_t = std::chrono::steady_clock::now();
        string diff = to_string((float)(std::chrono::duration_cast<std::chrono::milliseconds>(end_t-begin_t).count()) / 1000);
        //cout <<  "END SEC: " << st << " get pointer: " << pointer[0]  << "Time diff: " << diff<< endl;

        if(pointer[1] != "1"){
        (*arr).at(place) = {FuncId, diff};/// )};
        } else { 
            (*arr).at(place) = {FuncId, diff, pointer[1]}; /// )};
        }
        //cout << "out_res: " + (*arr).at(place)[0] << " " << (*arr).at(place)[1];

    }catch(exception e){
        // make sense to notify it gone from outer codes not from templates
        s << e.what();
        (*arr).at(place) = {FuncId, "<code_exec_error>" + s.str() + "</code_exec_error>"}; /// )};
    }

}

void CheckerDD(string(*checkerFunc)(string), string text){
    checkerFunc(text);
}

// I HAVE TO CREATE TRANSMIT FUNCTION TO PERFORM A REPRESENT FUNCTION
int main(){
    // THERE GOES CHECKER OF TIME LIITER LOGIC IF THERE COMES AT LEAST UNMATCHED TIME
    // IT WILL JUST STOP AND KILL ANY OTHER PROCESESS
    //std::thread(CheckerDD, Printer, "CheckerTest")

     
    string val[2] = {"1", "2"};
    string FuncIds[2] = {"FuncId1", "FuncId2"},
    vector<string(*)(string )> objectsd = {Printer, Printer};
    vector<vector<string>> catcher = {};
    vector<std::thread> threads;
    

    //int Objects_size = sizeof(objectsd) / sizeof(int(*)(int));

    // promise<vector<string>> promiser;
    // std::future<vector<string>> ff = promiser.get_future();
    for(int v = 1; v < 11; v ++){
        // THERE IT GOES VIA RUNTIME AND 
        // EACH PASSING ARGUMENT WILL HAVE PARAM NAME AS WELL WHEN PASSING
        catcher.push_back({FuncIds[0], "-1", "-1"});
        threads.push_back(std::thread(transmitFunc, Printer, Representers, &catcher, FuncIds, v, val[0] + to_string(v)));//
    }
    
    for(auto&& isd : threads) {
        std::cout << "joining: " << endl;
        isd.join();
    }
    
    std::cout << "FUNC GETTIME" << endl;
    vector<vector<string>> FinalCatcher;
    // fetching the results
    int conc = 0; //to handle finsihed funcs
    int conc2 = conc; // to make cycle on func check proceess

    while (true) {
        if(conc2 >= 10){
            break;
        }
        // vector<string> catcher = ff.get();
        cout << "<rrn::"<< catcher[conc][0] << ">";
        if(sizeof(catcher) / sizeof(vector<string>) > 0 
            && (catcher[conc][0] != "-1" 
            && catcher[conc][1] != "-1" 
            && catcher[conc][1] != "-1")
            ){
        
            cout << "Run time: " << catcher[conc][0] << endl << catcher[conc][1] << sizeof(catcher) / sizeof(vector<string>) << endl;
            if(catcher[conc][1].find("code_exec_error") != std::string::npos){
                cout << catcher[conc][1]; // to print the error
                break; // no more will come any compilation
            }   
            
            if(stod(catcher[conc][1]) > TIME_LIMIT){
                // THERE HAVE TO BE LOGIC OF STOPPING THE PROCESS
                cout << "<time_exceeded_err" << " time::" << catcher[conc][0] << "/>" << endl;
                break; // no more will come any compilation
            }else{
                cout << catcher[conc][1]; // finally printing the fun result
            }
            // THERE IS JUST NO NEED OF PRINTING TWICE ALL WE NEED
            // IS GETTING ITS ERROR OR TIME LIMIT EXEED ERROR
            // cout << "<" + catcher.at(0) + ">" + catcher.at(1) + "<>";
            // FinalCatcher.push_back(catcher);
            conc2 = conc2 + 1;
        }
        conc = conc + 1;
        cout << "</rrn::"<< catcher[conc][0] << ">";
    }

    // std::thread object_t(ffer, 3);
    // std::thread object_t2(ffer, 1);
    // auto  funcs[]  = {{ffer, }, ffer};  
    // object_t.join();
    // object_t2.join();
    // object_t2.detach();

    return 0;
}

// WRITE THE THING MAXIMALLY IN MERE TEMPLATE
// FOR INSTANCE MAKE 
