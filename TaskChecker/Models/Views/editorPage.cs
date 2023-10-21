using Microsoft.AspNetCore.Mvc.Rendering;

namespace TaskChecker.Models.Views
{
    public class editorPage
    {
       public string func { get; set; }
       public List<SelectListItem> langs { get; set; }
       public string lang { get; set; }
    }
}
