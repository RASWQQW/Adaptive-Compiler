using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Rendering;
using Microsoft.EntityFrameworkCore;
using System;
using System.Linq;
using System.Text.RegularExpressions;
using TaskChecker.contexts;
using TaskChecker.Controllers.Methods.Others;
using TaskChecker.Models;
using TaskChecker.Models.Views;

namespace TaskChecker.Properties.Controllers
{
    public class HomeController : Controller
    {
        Postgrecontext context { get; set; }

        public HomeController(ServiceProvider services)
        {
            context = services.GetService<Postgrecontext>();
        }
        [Route("[controller]/Testing")]
        public IActionResult Index()
        {
            return View("Index");
        }

        // here have to be topic and therefore it sends to its full
        // homework page where whom chooses and sends its work

        [Route("[controller]/run-code/{topic_id}/{lang?}")]
        public IActionResult TopicId(int topic_id, string lang)
        {    
            var result = context.Tasks.FromSql($"SELECT * FROM Tasks WHERE Tasks.task_name_id = '{topic_id}';`);").ToList();
            if (result.Count() > 0)
            {
                var GetFunc = context.Functions.FromSql($"SELECT * FROM Functions WHERE task_id = {result[0].id}").ToList();
                if (GetFunc.Count() > 0) {
                    Validation valClass = new Validation(context);
                    var page = new editorPage { };
                    // here goes all list of langs on ListForm
                    var SelectVals = new List<SelectListItem> { };
                    foreach (var inner in new List<string> { "java", "csharp", "python", "cpp" })
                    {
                        SelectVals.Add(new SelectListItem { Value = inner, Text = inner });
                    }
                    page.langs = SelectVals;

                    if (lang == null || lang == "cpp")
                    {
                        // here goes default func cunstraction
                        // namely cpp and its values
                        string args = valClass.ArgsMaker(GetFunc[0].id, "cpp");
                        page.func = $"static {GetFunc[0].return_value} {GetFunc[0].func_name}({args})" + "{\n\n}";
                        page.lang = "cpp";
                    }

                    else if (lang == "python") {
                        string args = valClass.ArgsMaker(GetFunc[0].id, "python");
                        page.lang = $"def {GetFunc[0].func_name}({args}) -> {Validation.DictContainsKey(GetFunc[0].return_value)}";
                    }
                    else if (lang == "java") {
                        string args = valClass.ArgsMaker(GetFunc[0].id, "java");
                        page.lang = $"public static {GetFunc[0].return_value} {GetFunc[0].func_name}({args})" + "{\n\n}";
                    }
                    else if (lang == "csharp") {
                        string allargs = valClass.ArgsMaker(GetFunc[0].id, "csharp");
                        page.lang = $"public static {GetFunc[0].return_value} {GetFunc[0].func_name}({allargs})" + "{\n\n}";
                    }
                    return View(page);
                }
            }
            return (base.Content("<body>There is no content which id is approach</body>"));
        }

        // to get and compile given code with a background response
        // and waiting without udating a page
        [Route("[controller]/Get-Proper-Code/{topic-id}/{lang?}")]
        [HttpPost]
        public IActionResult GetCode(string topic_id, string lang) {

            string Query = $"SELECT code FROM \"ProperCode\" WHERE topic_id = {topic_id} AND lang = {lang}";
            var ProperCode = context.Database.SqlQuery<List<string>>($"{Query}").First();
            return base.Content($"<h2>Proper Code<h3> <br><br> <input type=\"text\" value=\"{ProperCode}\" readonly>");

        }

        [Route("[controller]/addtopictask")]
        [HttpGet]
        [HttpPost]
        public IActionResult AddTopic(TaskCreator? saveObject = null) {

            if (HttpContext.Request.Method == "POST")
            {
                if (saveObject != null)
                {
                    if (saveObject.task_id.Count() > 0)
                    {
                        return View(base.Content("You cannot pick two task ids at one code"));
                    }
                    var current_task_id = saveObject.task_id.FirstOrDefault();
                    int c_task_id = Convert.ToInt32(current_task_id.Value);
                    Tasks task = context.Tasks.FirstOrDefault(w => w.id == c_task_id);
                    context.Functions.Add(
                        new Functions
                        {
                            task_id = task,
                            func_name =  saveObject.MethodName,
                            return_value = saveObject.returnType
                        });

                    List<string> argsName = new(), argsType = new();
                    foreach (Argument arg in saveObject.args)
                    {
                        argsName.Add(arg.Name);
                        argsType.Add(arg.Type);
                    }

                    context.ProperCode.Add(new ProperCode() { task_id = task, lang = saveObject.lan, code = saveObject.Code });
                    context.Arguments.Add(new Arguments() { task_id = task, args_name = argsName.ToArray(), args_types = argsType.ToArray() });
                    context.SaveChanges();
                }
                return View();
            }
            else {
                var tasker = new List<SelectListItem>() { };
                var AllTasks = context.Tasks.FromSql($"SELECT * FROM \"Tasks\"").ToList();
                foreach (Tasks task in AllTasks)
                { 
                    tasker.Add(new SelectListItem { Value = task.id.ToString(), Text = task.task_name_id });
                }
                var Creator = new TaskCreator()
                {
                    task_id = tasker,
                };

                return View(Creator);
            }
        }

        [Route("[controller]/Create-topic")]
        public IActionResult CreateTopic()
        {
            if (Request.Method == "GET")
            {
                return View();
            }
            else {
                if(Request.Form.ContainsKey("task_name_id"))
                {
                    string task_nid = Request.Form["task_name_id"][0];
                    string dscr = Request.Form["DSCR"][0];
                 
                    context.Database.ExecuteSql($"INSERT INTO \"Tasks\"(task_name_id, text) VALUES({task_nid}, {dscr ?? "No any descr"});");
                    var lastsaving = context.Tasks.FromSql($"SELECT * FROM \"Tasks\" Order by id desc LIMIT 1").ToList();
                    return base.Content($"<h1 style = \"color: green;\">Last value succesfully added<h1/>id: {lastsaving[0].id}, task_name: {lastsaving[0].task_name_id}", "text/html");
                }
            }
            return base.Content("Your value is not proper cause of not filling a form", "text/html");
        }
        [Route("[controller]/run-code")]
        public IActionResult RunCodeOnGO() {

            return View("Sender");
        }
    }
}
