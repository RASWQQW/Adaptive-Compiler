using Confluent.Kafka;

namespace TaskChecker.BG
{
    public class KafkaProducer : BackgroundService
    {
        IProducer<Null, string> built;
        public KafkaMain kafkaMain;
        public KafkaProducer()
        {
            kafkaMain = new KafkaMain();
            built = new ProducerBuilder<Null, string>(config: kafkaMain.ConsConf).Build();
        }
        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            var val = built.ProduceAsync(KafkaMain.topicName, new Message<Null, string>()).Result;
            //throw new NotImplementedException();
            return Task.CompletedTask;
        }
        public override Task StopAsync(CancellationToken stoppingToken) {
            return Task.CompletedTask;
        }
    }
}
