package job

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/segmentio/kafka-go"
)

// 评价数据流处理

// JobWorker 自定义执行job的结构体，实现 transport.Server
type JobWorker struct {
	kafkaReader *kafka.Reader              // kafka reader
	ESClient    *elasticsearch.TypedClient // ES Client
	log         *log.Helper
}

type ESClient struct {
	*elasticsearch.TypedClient
	Index string
}

// Start kratos程序启动之后会调用的方法
// ctx 是框架启动的时候传入的ctx， 是带有退出取消的
func (jw JobWorker) Start(ctx context.Context) error {
	jw.log.Debug("JobWorker start....")
	// 1. 从kafka中获取MySQL中的数据变更消息

	// 接收消息
	for {
		m, err := jw.kafkaReader.ReadMessage(ctx)
		if errors.Is(err, context.Canceled) {
			return nil
		}
		if err != nil {
			jw.log.Error("ReadMessage from kafka failed, err:%v", err)
			break
		}
		jw.log.Debug("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		// 2. 将完整的评价数据写入es

	}

	return nil
}

// Stop kratos结束后会调用的
func (jw JobWorker) Stop(context.Context) error {
	jw.log.Debug("JobWorker stop...")
	// 程序退出前关闭Reader
	return jw.kafkaReader.Close()
}

func readFromKafka() {
	// 创建一个reader，指定GroupID，从 topic-A 消费消息
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  "consumer-group-id", // 指定消费者组id
		Topic:    "topic-A",
		MaxBytes: 10e6, // 10MB
	})

	// 接收消息
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}

	// 程序退出前关闭Reader
	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}

}

func connES() {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	// 创建客户端连接
	client, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
		return
	}

}

// indexDocument 索引文档
func indexDocument(client *elasticsearch.TypedClient) {
	// 定义 document 结构体对象
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个好评！",
		Tags: []Tag{
			{1000, "好评"},
			{1100, "物超所值"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	// 添加文档
	resp, err := client.Index("my-review-1").
		Id(strconv.FormatInt(d1.ID, 10)).
		Document(d1).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}

// updateDocument 更新文档
func updateDocument(client *elasticsearch.TypedClient) {
	// 修改后的结构体变量
	d1 := Review{
		ID:      1,
		UserID:  147982601,
		Score:   5,
		Content: "这是一个修改后的好评！", // 有修改
		Tags: []Tag{ // 有修改
			{1000, "好评"},
			{9000, "有图"},
		},
		Status:      2,
		PublishTime: time.Now(),
	}

	resp, err := client.Update("my-review-1", "1").
		Doc(d1). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}
