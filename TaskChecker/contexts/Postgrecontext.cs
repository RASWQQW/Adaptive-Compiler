using Microsoft.AspNetCore.Components;
using Microsoft.EntityFrameworkCore;
using Newtonsoft.Json;
using System.Text.Json.Serialization;
using TaskChecker.Models;

namespace TaskChecker.contexts
{
    public class Postgrecontext : DbContext
    {
        string Tokens { get; set; }
        string SourceName { get; set; } = "";

        public string SetTokens()
        {
            string curDirect = Directory.GetCurrentDirectory().ToString();
            string CurrentPath = Path.Combine(curDirect, "contexts\\BaseContent.json");
            using (var reader = new StreamReader(CurrentPath)) {
                string json = reader.ReadToEnd();
                return JsonConvert.DeserializeObject<Dictionary<string, object>>(json)["BaseValue"].ToString();
            }
        }

        public Postgrecontext(DbContextOptions<Postgrecontext> options)  : base(options) {
            this.Tokens = this.SetTokens();
            Database.EnsureCreated();
        }
        
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.UseSerialColumns();
            base.OnModelCreating(modelBuilder);
        }
        
        protected override void OnConfiguring(DbContextOptionsBuilder builder) {

            builder.UseNpgsql(this.Tokens);
        }

        public DbSet<Tasks> Tasks { get; set; }
        public DbSet<Functions> Functions { get; set; }
        public DbSet<ProperCode> ProperCode { get; set; }
        public DbSet<Arguments> Arguments { get; set; }
        public DbSet<Checker> Checker { get; set; }
    }
}
