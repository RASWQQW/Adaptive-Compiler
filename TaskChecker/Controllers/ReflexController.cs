using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Primitives;
using System.Reflection;
using TaskChecker.contexts;

namespace TaskChecker.Controllers
{
    public class ReflexController : Controller
    {
        PropertyInfo Rcontext;
        Postgrecontext context;
        public ReflexController(Postgrecontext newcontext)
        {
            context = newcontext; 
        }

        [Route("[controller]/MakeInsert/{ObjectType?}")]
        [HttpGet]
        [HttpPost]
        public IActionResult MakeInsert(string ObjectType) {
            Rcontext = typeof(Postgrecontext).GetProperty(ObjectType);
            if (Rcontext != null) {
                if (Rcontext.PropertyType.FullName.Contains("DbSet"))
                {
                    var ObjectName = Rcontext.PropertyType.GetGenericArguments()[0];
                    var NewObject = Activator.CreateInstance(Type.GetType(ObjectName.FullName));
                    PropertyInfo[] properties = Type.GetType(ObjectName.FullName).GetProperties();

                    return View(new Tuple<PropertyInfo[], object>(properties, NewObject));
                }
            }
            string message = "The given object is not in the dbcontext";
            return base.Content($"<h1 style=\"color: red\">{message}<h1\\>", "text/html");
            throw new Exception("The given object is not in the dbcontext");
        }

        [Route("[controller]/InsertObject")]
        public IActionResult InsertObject() {
            // Selects all from object by using reflex
            HttpContext.Request.Form.TryGetValue("ClassName", out StringValues ClassNameValue);
            var InsType = Type.GetType(ClassNameValue);
            var ClassInstance = Activator.CreateInstance(InsType);
            IFormCollection FormVals = HttpContext.Request.ReadFormAsync().Result;
            foreach(var keys in FormVals) {
                if(keys.Key != "ClassName"){
                    var QueryfindVal = InsType.GetProperty(keys.Key);
                    if (QueryfindVal != null){
                        QueryfindVal.SetValue(ClassInstance, keys.Value.ToString());
                    }
                }
            }
            context.Add(ClassInstance);
            context.SaveChanges();
            return View();
        }
    }

}
