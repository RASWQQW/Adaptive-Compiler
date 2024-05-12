using namespace std;
#include <iostream> 
#include <thread>
#include<chrono>
#include <vector>
#include <string>
#include <sstream>
#include <bits/stdc++.h>
#include <stdio.h>
#include <windows.h> // for EXCEPTION_ACCESS_VIOLATION
#include <excpt.h>

//https://en.cppreference.com/w/cpp/thread/thread/native_handle
auto ffer = [](int timesd){
    cout << "CheckerCode" << endl;
    //sleep(timesd);
    cout <<  "CheckerCode2 >>" << timesd << endl;
};

// int filterException(int code, PEXCEPTION_POINTERS ex) {
//     std::cout << "Filtering " << std::hex << code << std::endl;
//     return EXCEPTION_EXECUTE_HANDLER;
// }
// THERE IT HAVE TO BE REWRITTEN TO THE FUNCTIONS
// BETTER CAN TRY MACROS TO MAKE IT MORE FITTING
// ALL OF THEM FUNC_PREPARE VALUES AND STATIC
// typedef string ret_originType;

int filter(unsigned int code, struct _EXCEPTION_POINTERS *ep)
{
    cout << "in filter.";
    if (code == EXCEPTION_ACCESS_VIOLATION)
    {
        cout << "caught AV as expected.";
        return EXCEPTION_EXECUTE_HANDLER;
    }
    else
    {
        cout << "didn't catch AV, unexpected.";
        return EXCEPTION_CONTINUE_SEARCH;
    };
}

typedef int** to_run_ret;
typedef to_run_ret (*run_code_temp)(string, string[]); //  there we need just pass parameters of func when preparing
double TIME_LIMIT = 3.13;   

template <class T> bool TypeHandler(const T &TypeChecker)
	{
		if(!is_array<T>::value){
			return 0;
		}
		return 1;
	}
    

template<typename T>
void ValPrinter(T PrintType, string getRetVal[]){
    (getRetVal)[1] = "1";
    getRetVal[0] = to_string(PrintType);
}
     
template <typename T, int W, int H>
void DDArrayChecker(T **vals, string getRetVal[]){
    string Printer = "";
    // cout << "Inside " << H <<" " <<  W << " " << vals[0][0];
    for(int v = 0; v < 20; v++){
        cout << "Inside2" << vals[v][0] << endl;
        if(!TypeHandler(vals[v])
            ){ // END FROM HERE TO MAKE SOME SENSE TO CREATE IDEAL SORTING ON TYPE RECOGNITION
                for(int v2 = 0; v2 < H; v2++){
                    //*(*(vals2 + v) + v2) = vals[v][v2];
                    cout << "to print " << vals[v][v2] << " ";
                    Printer = Printer + to_string(vals[v][v2]) + " ";
                }
                Printer = Printer + "\n";
            }
        }
        cout << "AFTER LOOP: " << endl;
        (getRetVal)[1] = "1";
        getRetVal[0] = Printer;
        cout << "Final " << H << endl;
}

template <typename T, int W>
void DArrayChecker(T *vals,  string getRetVal[]){
    string Printer = "";
    for(int v = 0; v < W; v++){
        // Only prints types with single value
        if(!TypeHandler(vals[v])){
            Printer = Printer + to_string(vals[v]) + " ";
        }
    }
    (getRetVal)[1] = "1";
    getRetVal[0] = Printer;
}



string Printerd(string PrintType, string PrintType2[]){
    return PrintType;
}

int check(){
    return 12;
}

int** CodePrinter2d(string val1, string val2[]){
    int val[2][2] = {{1, 2}, {1, 2}};
    int (*vv)[2][2] = &val;
    int **checker;
    checker  = new int *[2];
    (*checker) = new int[2]{1, 2};
    *(checker+1) = new int[2]{3, 4};
    return checker;
}

int* CodePrinter(string val1, string val2[]){
    return new int[2]{12, 45};
}


// void Representers(string getRetVal[], to_run_ret val){
//     //cout << "Start: " << val << endl <<"pointer: " << pp;
//     std::this_thread::sleep_for(std::chrono::milliseconds(3050 + stoi(val[0][0])));
//     (getRetVal)[1] = "1";
//     getRetVal[0] = "pointer: " + val[0][0];
//     //cout << "pp val was given to " << getRetVal[1] << endl << getRetVal[0] << endl;
// }



// FUNC PREPARE TEMPLATE
// COMPLETE TRANSMIITER
void transmitFunc(run_code_temp to_run, 
                //   void(*Representer)(string[], to_run_ret), 
                  string rep_type,
                  vector<vector<string>> *arr, 
                  string FuncId,
                  int place,
                  float time_limit_val,
                  string parameter1,
                  string parameter2[1] /*THERE FUNC WILL BE PREPARED AS TEMPLATE TO OWN PARAMS*/){

    std::chrono::steady_clock::time_point begin_t = std::chrono::steady_clock::now();
    std::stringstream stream_gather;

    try{
        //std::cout << "FUNC EXCEEDED CHECKbefore" << endl;
        string pointer[2] = {"0", "0"};
        ///throw runtime_error("CHECK_TEXT!");
        auto runner = [rep_type, to_run, parameter1, parameter2, &pointer]() {
            
            to_run_ret valer = to_run(parameter1, parameter2);
            // typedef std::remove_reference<decltype( **valer )>::type s1_remove_originType;

            if(1 != 1){

            }
            // else if (rep_type == "Printer"){
            //     ValPrinter<to_run_ret>(valer, pointer);
            // }    
            else if (rep_type == "DDArrayChecker" ) {//&& std::is_pointer<s1_remove_originType>::value()
                to_run_ret pointer_pass_val;
                //https://en.cppreference.com/w/cpp/types/is_pointer#:~:text=std%3A%3Ais_pointer%20is%20a,value%20is%20equal%20to%20false.
                
                int setter_sizer = sizeof(valer) / sizeof(valer[0]);

                //BEFORE EXPLICIT CONVERT                
                // pointer_pass_val = new ret_originType *[setter_sizer];
                // for(int s = 0; s < setter_sizer; s++){
                //     pointer_pass_val[s] = valer[s];
                // }

                // typedef std::remove_reference<decltype( *valer )>::type deffer;
                // if (typeid(deffer).name() != typeid(ret_originType).name()){

                // to_run_ret (&conv_type)[][(sizeof(valer[0]) / sizeof(ret_originType))] = valer; 
                // HERE I CAN WRITE VIA STRING BUT PROBLEM IS ON ONE LEVEL POINTER
                // NOT IN THE SECOND, SO I HAVE TO CHECK IN FIRST AND DETERMINE IT HAS VALUE OR NOT
                std::stringstream string_tug;
                typedef std::remove_reference<decltype( **valer )>::type ret_originType;
                int enrichStat = 50;
                int CountGetter = 0;
                for(int dds = 0; dds < enrichStat; dds++){
                    __try {
                        try{
                            if(to_string(valer[dds][0]).size() >= 6){  // have to set some  limit how generated values are will be passed below 1kk
                                break;
                            }
                            CountGetter = CountGetter + sizeof(valer[0]);
                        }
                        catch (const exception &exp){
                            string_tug << exp.what();
                            if(string_tug.str().find("segmentation") != std::string::npos){
                                if(string_tug.str().find("fault") != std::string::npos){
                                    break;
                                }
                            }
                        }
                    } __except(filter(GetExceptionCode(), GetExceptionInformation())){
                        std::cout << "caught" << std::endl;
                    }   
                    
                }
                // int (*val)[10][10] = &valer;
                // std::array<valer>();

                std::cout <<  "ss: " << CountGetter << " before run size: " << sizeof(*valer) << "  " << sizeof(valer[0]) << "||" <<  sizeof(valer[0]) / sizeof(ret_originType) << endl;
                DDArrayChecker<ret_originType, sizeof(CountGetter) / sizeof(valer[0]), sizeof(valer[0]) / sizeof(ret_originType)>(valer, pointer);
                // }
            }
            // else if (rep_type == "DArrayChecker"){
            //     to_run_ret pointer_pass_val;
            //     typedef std::remove_reference<decltype( *pointer_pass_val )>::type ret_originType;
            //     DArrayChecker<ret_originType, sizeof(valer) / sizeof(ret_originType)>(valer, pointer);
            // }

        };

        std::thread ddx(runner);
        ddx.detach();
        // ddx.join();
        // float checker = ((place + rand() % 3)) * ((float) 1 / 21);
        //cout << "FUNC EXCEEDED CHECK1 pointer: " << pointer << endl;
        int st = 0;  
        while(st < (time_limit_val*1000/15.6436)){
            std::this_thread::sleep_for(std::chrono::milliseconds(1));
            st = st + 1;
            if(pointer[1] == "1"){
                break;
            }        
        }
        std::chrono::time_point end_t = std::chrono::steady_clock::now();
        string diff = to_string((float)(std::chrono::duration_cast<std::chrono::milliseconds>(end_t-begin_t).count()) / 1000);
        //cout <<  "END SEC: " << st << " get pointer: " << pointer[0]  << "Time diff: " << diff<< endl;

        if(pointer[1] != "1" && pointer[1] == "0"){
            (*arr).at(place) = {FuncId, diff, "0"};/// )};
        } else { 
            (*arr).at(place) = {FuncId, diff, pointer[0]}; /// )};
        }
        //cout << pointer[1] << " " << pointer[0] << "out_res: " + (*arr).at(place)[0] << " " << (*arr).at(place)[1] << (*arr).at(place)[2];

    }catch(const exception &e){
        // make sense to notify it gone from outer codes not from templates
        //cout << "come to error" << endl;
        stream_gather << e.what();
        //cout << "errorcode: " << s.str() << endl;
        (*arr).at(place) = {FuncId, "<code_exec_error>" + stream_gather.str() + "</code_exec_error>", "0"}; /// )};  last zero only means its proper end 
    }
}

void CheckerDD(string(*checkerFunc)(string), string text){
    checkerFunc(text);
}

struct params {
    string  val1;
    string val2[1];
};

// I HAVE TO CREATE TRANSMIT FUNCTION TO PERFORM A REPRESENT FUNCTION
int main(){
    
    // THERE GOES CHECKER OF TIME LIITER LOGIC IF THERE COMES AT LEAST UNMATCHED TIME
    // IT WILL JUST STOP AND KILL ANY OTHER PROCESESS
    //std::thread(CheckerDD, Printer, "CheckerTest")

    // pre-forming values
    params values__lister[10] = {{"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}, {"1", {"2"}}};
    string FuncIds[10] = {"FuncId1", "FuncId2", "FuncId1", "FuncId2", "FuncId1", "FuncId2", "FuncId1", "FuncId2", "FuncId1", "FuncId2"};
    float given_time_limits[10] = {3.12, 3.12, 3.12, 3.12, 3.12, 3.12,3.12, 3.12,3.12, 3.12};
    vector<run_code_temp> objectsd;// = {CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter, CodePrinter};

    for(int v = 0; v < 10; v++){
        objectsd.push_back(CodePrinter2d);
    }
    
    // after forming values
    vector<vector<string>> catcher = {};
    vector<std::thread> threads;

    int Objects_size = objectsd.size();
    int valSent = 0;
    // promise<vector<string>> promiser;
    // std::future<vector<string>> ff = promiser.get_future();
    for(int v = 0; v < 1; v ++){
        // THERE IT GOES VIA RUNTIME AND 
        // EACH PASSING ARGUMENT WILL HAVE PARAM NAME AS WELL WHEN PASSING
        catcher.push_back({FuncIds[v], "-1", "-1"});
        /*  v - is place of func to  locate its value on a list
            given_time_limits - is unique time limit collection of each func
        */
        
        threads.push_back(std::thread(transmitFunc, objectsd[v], /*have to come in string way of func*/ "DDArrayChecker", &catcher, FuncIds[v], v, given_time_limits[v], /**//* further goes changed values */values__lister[valSent].val1, values__lister[valSent].val2));//

        if (v % 2 == 1) {
            valSent = valSent + 1;
        }
    }
    
    for(auto&& isd : threads) {
        isd.join();
    }
    
    std::cout << "FUNC GETTIME" << endl;
    vector<vector<string>> FinalCatcher;
    vector<int> OnPerformFuncs;
    std::vector<int>::iterator it;

    // fetching the results

    int conc = 0; //to handle finsihed funcs
    int conc2 = conc; // to make cycle on func check proceess
    bool istostop = false;
    string funcId = "";
    int reverser = 1;
    while (true) {
        
        if(istostop) {
            break; // stops when there is no need
        }

        if(OnPerformFuncs.size() >= Objects_size){
            break;
        }

        if(conc >= Objects_size){ // if there is loop to look out is ended
            if(reverser == 1 || int(OnPerformFuncs.size() >= Objects_size / 2)){
                conc = int(Objects_size / 2);
                reverser = 0;
            }else{
                conc = 0;
                reverser = 1;
            }
            // int min_to_start = 999; //it may change
            // OnPerformFuncs.clear();

            // for(int in = 0; in < catcher.size(); in++){
            //     it = std::find(OfPerformFuncs.begin(), OfPerformFuncs.end(), in);
            //     if (it != OfPerformFuncs.end()){    
            //         OnPerformFuncs.push_back(conc);
            //         if(min_to_start > conc){
            //             min_to_start = conc;
            //         }
            //     }
            // }        
            // conc = min_to_start;
        }
        
        it = std::find(OnPerformFuncs.begin(), OnPerformFuncs.end(), conc);
        if (it == OnPerformFuncs.end()){
            // vector<string> catcher = ff.get();
            funcId = catcher[conc][0];
            cout << "<rrn::"<< funcId << ">";
            //cout << "Size: " << sizeof((catcher[conc][0])) << "Div: " << sizeof(string) << sizeof(vector<string>);
            if(catcher[conc].size() >= 2 
                // && (catcher[conc][0] != "-1" 
                && catcher[conc][1] != "-1" 
                && catcher[conc][2] != "-1"
                ){
            
                //cout << "Run time: " << catcher[conc][0] << endl << catcher[conc][1] << sizeof(catcher) / sizeof(vector<string>);
                //cout << "Error is true: " << catcher[conc][1].find("code_exec_error");
                if(catcher[conc][1].find("code_exec_error") != std::string::npos){
                    istostop = true;
                    cout << catcher[conc][1]; // to print the error
                    //break; // no more will come any compilation
                }else if(stod(catcher[conc][1]) > TIME_LIMIT){
                    istostop = true;
                    // THERE HAVE TO BE LOGIC OF STOPPING THE PROCESS
                    cout << "<time_exceeded_err" << " time::" << catcher[conc][1] << " func::" << funcId << "/>";
                    //break; // no more will come any compilation
                }else{
                    if (catcher[conc].size() >= 3 || catcher[conc][2] == "-1"){
                        cout << "<res>" << catcher[conc][2] << "</res>"; // finally printing the fun result
                        cout  << endl << "<stat::run_time:" << catcher[conc][2] <<  "/>"; // there i should add value of mempor usage
                    }
                }
                // THERE IS JUST NO NEED OF PRINTING TWICE ALL WE NEED
                // IS GETTING ITS ERROR OR TIME LIMIT EXEED ERROR
                // cout << "<" + catcher.at(0) + ">" + catcher.at(1) + "<>";
                // FinalCatcher.push_back(catcher);
                conc2 = conc2 + 1;
                OnPerformFuncs.push_back(conc);
            }
            conc = conc + 1;
            cout << "</rrn::"<< funcId << ">\n";
        }
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
