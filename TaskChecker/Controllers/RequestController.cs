using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Primitives;
using TaskChecker.contexts;

namespace TaskChecker.Properties.Controllers
{
    public class RequestController : Controller
    {
        public Postgrecontext _context { get; set; }

        public RequestController(Postgrecontext context){

            _context = context;
        }


        [Route("[controller]/TgCreate")]
        public IActionResult Get() {
            var res = _context.Database.SqlQuery<string>(
                $"CREATE OR REPLACE TRIGGER Wonder BEFORE INSERT OF name ON Users FOR EACH ROW EXECUTE FUNCTION Checker();").ToList();
            return View();
        }

        [Route("[controller]/Saver")]
        public IActionResult Printer()
            {
            string username = "Akella";
            //_context.Database.ExecuteSql($"create table if not exists checker(id serial primary key, name varchar);");
            //_context.Database.ExecuteSql($"Insert INTO checker(name) VALUES({username})");
            string getUserName = _context.Database.ExecuteSql($"SELECT name FROM checker").ToString();
            List<string> getusername = _context.Database.SqlQuery<string>($"SELECT name FROM checker").ToList();
            string html = $"<p>Create new value</p><div><h1>{getusername[0]}</h1></div>";
            return base.Content(html, "text/html");
        }

        
        [Route("[controller]/AjaxCheck")]
        [HttpPost]
        [HttpGet]
        public IActionResult AjaxCheck() {
            if (HttpContext.Request.Method == "POST")
            {
                var json = "";
                // var formvals = Request.Form.Keys;
                if (Request.HasJsonContentType()){
                   json = Request.ReadFromJsonAsync<Dictionary<string, string>>().Result["val"];

                }

                string newHtml = $"<h1>isval {json}</h1>";
                return base.Content(newHtml, "text/html");

            }
            else if (HttpContext.Request.Method == "GET") {
                return View("Checker");
            }
            return base.Content("There is no post, only get", "text/html");
        }
    }
}
