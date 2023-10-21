using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Rendering;

namespace TaskChecker.Models.Views
{
    public class TaskCreator
    {
        public List<SelectListItem> task_id { get; set; }

        [FromForm(Name = "Language")]
        public string? lan { get; set; } = "cpp"; // it default gets c++

        [FromForm(Name = "Method Name")]
        public string? MethodName { get; set; }

        [FromForm(Name = "Proper Code")]
        public string? Code { get; set; }

        [FromForm(Name = "Arguments")]
        public List<Argument> args { get; set; }

        [FromForm(Name = "Return value type")]
        public string returnType { get; set; }
    }
}
