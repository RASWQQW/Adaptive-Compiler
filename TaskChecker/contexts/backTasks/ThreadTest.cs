using System.Diagnostics;

namespace TaskChecker.contexts.backTasks
{
    public class ThreadTest : IHostedService
    {
        public static void BackMethod(int san) {
            Thread.Sleep(san);
            Debug.WriteLine($"is Done after, {san}");
        }
        public static void BackMethod() { }
        public static async Task<string> Baker() {
            return "Baker is pretty small";
        }

        public Task StartAsync(CancellationToken cancellationToken)
        {
            int val = 10;
            Thread nt = new Thread(new ThreadStart(BackMethod));
            nt.Start(val);

            Task.WaitAll(new Task[]{ Baker() });

            Task istask = Task.Run(action: () => BackMethod(12));
            istask.Start();

            Debug.WriteLine(istask + "" + istask.IsCompleted + "" + Task.CurrentId);
            return Task.CompletedTask;
        }

        public Task StopAsync(CancellationToken cancellationToken)
        {
            throw new NotImplementedException();
        }
    }
}
