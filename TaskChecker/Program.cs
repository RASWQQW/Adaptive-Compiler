using Microsoft.AspNetCore.Mvc.Razor;
using TaskChecker.contexts;
using TaskChecker.contexts.backTasks;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddRazorPages();
builder.Services.AddDbContext<Postgrecontext>();
builder.Services.AddControllers();
builder.Services.AddMvc(option => option.EnableEndpointRouting = false);

//builder.Services.AddHostedService<ThreadTest>();

//builder.Services.Configure<RazorViewEngineOptions>(o => {
//    o.ViewLocationFormats.Clear();
//    o.ViewLocationFormats.Add(
//        "MVC/Views/{1}/{0}" + RazorViewEngine.ViewExtension);
//});



var app = builder.Build();


app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.UseMvc();
app.UseHttpsRedirection();
app.UseStaticFiles();
app.UseRouting();
app.UseAuthorization();
app.MapRazorPages();
app.Run();
