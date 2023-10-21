using Microsoft.AspNetCore.Mvc;
using text = System.Text;
using TaskChecker.BG;
using TaskChecker.contexts;
using TaskChecker.Models.Views;

namespace TaskChecker.Controllers
{
    public class CompillerController : Controller
    {

        public Postgrecontext _context { get; set; }
        public CompillerController(Postgrecontext contextis)
        {
            _context = contextis;
        }


        // here goes check of code and returning whether it 
        // correct or not

        // Here goes checking of the code when it arrives
        // and returns value as request response back
        public IActionResult Checker(CodeChecking data)
        {
            KafkaMain compiler = new KafkaMain();
            KafkaProducer producer = new KafkaProducer();
            // here goes that is proper or not 
            return base.Content("");

        }

        public IActionResult CheckerApi() {

            // There have to be api checker
            // which sends respond and gets correspoding result


            using(HttpClient client = new HttpClient())
            {
                var GivenLink = "";
                var args = new Dictionary<string, string>() { };

                var formContent = new MultipartFormDataContent();

                formContent.Add(new MultipartContent()) ;
                var bytes = new ByteArrayContent(text.Encoding.UTF8.GetBytes(GivenLink));
                var argsConv = new StringContent(args.ToString(), text.Encoding.UTF8, "application/json");
                client.PostAsync(GivenLink, argsConv);
            }

                return View();
        }
    }
}
