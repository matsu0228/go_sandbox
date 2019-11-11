# Cloud Pub/Sub

## 事前準備

- GCP Console > [Cloud Pub/Sub](https://console.cloud.google.com/cloudpubsub/topic/list?hl=ja&project=spolive-dev) から、下記を作成しておく
  - Topic
  - Subscription

- それぞれのidを控えておく

- 権限のあるSAのcredential JSONを用意する ( firebase credential JSONでも利用できそう)


## 実行


```
# message送信
go run pubsub.go publisher.go  

# message受信
go run pubsub.go subscriber.go  
```


## 参考

- 概要: https://cloud.google.com/pubsub/docs/overview?hl=ja
- 対応言語: https://cloud.google.com/pubsub/docs/tutorials?hl=ja
  - Go, Python, Node.jsで、利用可能