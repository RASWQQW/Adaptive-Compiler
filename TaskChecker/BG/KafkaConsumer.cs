using Confluent.Kafka;
using System.Diagnostics;

namespace TaskChecker.BG
{
    public class KafkaConsumer : BackgroundService
    {

        KafkaMain kafkaMain;
        IConsumer<Ignore, string> build;
        public KafkaConsumer()
        {
            kafkaMain = new KafkaMain();
            build = new ConsumerBuilder<Ignore, string>(kafkaMain.ConsConf).Build();
        }

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            build.Subscribe(KafkaMain.topicName);
            while (stoppingToken.IsCancellationRequested) {
                var values = build.Consume(new CancellationTokenSource().Token);
                var Runner = Task.Run(() => CodeRunnerByConsume());
                Debug.WriteLine("The message is gained: " + Runner.Result + "Values: " + values.Message + "Add value: " + values.Value);
            }
            return Task.CompletedTask;
        }

        public async Task<List<string>> CodeRunnerByConsume() {
            return new List<string> { };
        }

        public List<string> CompilerPro(string genPath) {

            Process cmd = new Process();
            cmd.StartInfo.FileName = "powershell.exe";
            cmd.StartInfo.RedirectStandardInput = true;
            cmd.StartInfo.RedirectStandardOutput = true;
            cmd.StartInfo.CreateNoWindow = true;
            cmd.StartInfo.UseShellExecute = false;
            cmd.Start();

            cmd.StandardInput.WriteLine($"cd {genPath}");
            cmd.StandardInput.WriteLine($"cd {genPath}");

            cmd.StandardInput.Flush();
            cmd.StandardInput.Close();

            cmd.WaitForExit();
            var result = cmd.StandardOutput.ReadToEnd();
            var errors = cmd.StandardError.ReadToEnd();

            return new List<string> { result, errors};
        }
    }
}
