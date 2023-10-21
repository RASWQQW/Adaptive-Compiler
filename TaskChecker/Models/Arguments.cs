namespace TaskChecker.Models
{
    public class Arguments
    {
        public int id { get; set; }
        public Tasks task_id { get; set; }
        public string[] args_name { get; set; }
        public string[] args_types { get; set; }
    }
}

