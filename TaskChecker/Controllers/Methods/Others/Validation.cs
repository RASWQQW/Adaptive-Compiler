using Microsoft.EntityFrameworkCore;
using System.Text.RegularExpressions;
using TaskChecker.contexts;

namespace TaskChecker.Controllers.Methods.Others
{
    public class Validation
    {
        public Postgrecontext _context { get; set; }


        // Here a exceptions in types 
        public static Dictionary<string, string> pyList = new() { { "string", "str" }, { "decimal", "float" } }; // its for converting type specially for another
        public static Dictionary<string, string> javaList = new() { {"string", "String" } };

        public Validation(Postgrecontext context) {
            _context = context;
        }

        public static string DictContainsKey(string searchValue, Dictionary<string, string> dict = null) {
            if (dict == null) { dict = pyList; }
            return pyList.ContainsKey(searchValue) ? dict[searchValue] : searchValue;
        }

        public static string RemoveDigits(string source2) {
            string source = "";
            for (int v = 0; v < source2.Length; v++) {
                if (!char.IsDigit(source2[v])){
                    source += source2[v];
                }
            }
            return source;
        }

        public string ArgsMaker(int FuncId, string lang)
        {
            Func<string, List<string>?, string> RemoveAll = (val, replacers) => {
                
                replacers.AddRange(new List<string>{ "[", "]"});
                foreach (string sl in replacers) {
                    val = val.Replace(sl, "");
                }
                return val;
            };

            Func<List<List<string>>, string, List<string>> Cycler = (ArgsFr, lefter) =>
            {
                List<string> args = new List<string>();
                var FinderReg = new Regex(@"\[\d*\]+");
                for (int i = 0; i < ArgsFr[0].Count(); i++)
                {
                    var ArgTypeG = ArgsFr[1][i];
                    var ArgNameG = ArgsFr[0][i];

                    if (lefter == "cpp" && ArgsFr[1][i].Contains("[")){
                        string InnerPiece = FinderReg.Match(ArgsFr[1][i]).Value;
                        string ArgName = ArgsFr[0][i] + InnerPiece;
                        string ArgType = RemoveAll(RemoveDigits(ArgTypeG), new List<string> { });
                        args.Add($"{ArgType} {ArgName}");
                    }
                    else if (lefter == "python"){
                        if (ArgTypeG.Contains("[")){
                            
                            var WholeValue = RemoveAll(RemoveDigits(ArgTypeG), new List<string> { });
                            WholeValue = DictContainsKey(searchValue: WholeValue);
                            
                            var Matches = FinderReg.Matches(ArgTypeG);
                            for (int v = 0; v < Matches.Count(); v++) {
                                WholeValue = "list[" + WholeValue + "]";
                            }
                            args.Add($"{ArgsFr[0][i]} : {WholeValue}");
                            
                            // Dictionary<int, object>[,] app = new Dictionary<int, object>[2, 2]{ { 1, 2 }, { 1, 2 } };
                        }else{
                            args.Add($"{ArgsFr[0][i]}: { DictContainsKey(searchValue: ArgTypeG)}");
                        }
                    }
                    else if(lefter=="java"){
                        // there all other things gets its usual main form
                        args.Add($"{javaList[ArgTypeG]} {ArgsFr[0][i]}");
                    }

                    else if (lefter == "csharp"){
                        args.Add($"{ArgTypeG} {ArgsFr[0][i]}");

                    }
                }
                return args;
            };

            string searchString =
                                $"SELECT FunctionArgs.args_name, FunctionArgs.args_types" +
                                $"FROM FunctionArgs" +
                                $"Join Functions on FunctionArg.task_id = Functions.task_d" +
                                $"WHERE Functions.id = {FuncId}";

            List<List<string>> GetArgs = _context.Database.SqlQuery<List<string>>($"{searchString}").ToList();
            if (GetArgs[0].Count() == GetArgs[1].Count())
            {
                List<string> args = Cycler(GetArgs, lang);
                // adding a type to string to show in func braces
                string newArgs = string.Join(", ", args);
                return newArgs;

            }
            else {
                throw new Exception("Type and Name of current func values are not mathces");
            }
            
        }
    }
}
