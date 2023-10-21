using Confluent.Kafka;

namespace TaskChecker.BG
{


    // There stays group name and other constants
    public class KafkaMain
    {
        static public string host = "localhost:9092";
        static public string topicName = "Topicer";
        public ProducerConfig ProducerConf { get; set; }
        public ConsumerConfig ConsConf { get; set; }

        public AdminClientConfig AdminClientConf { get; set; }

        public KafkaMain()
        {
            ProducerConf = new ProducerConfig() { 
                BootstrapServers = host,
            };

            ConsConf = new ConsumerConfig()
            {
                GroupId = "GrId", // I have to alose find it
                BootstrapServers = host,
                AutoOffsetReset = AutoOffsetReset.Earliest
            };

            AdminClientConf = new AdminClientConfig()
            {
                BootstrapServers = host,
            };
    }
}
}
