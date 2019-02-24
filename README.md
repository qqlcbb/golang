# golang学习代码
> 20190101

## 启用
1. 使用docker启用ElatisSearch,并使用9200端口
2. 启用worker进行爬取页面，可启动n个worker
```
go run crawier_distributed/worker/server/worker.go --port 9001
go run crawier_distributed/worker/server/worker.go --port 9002
go run crawier_distributed/worker/server/worker.go --port 9003
go run crawier_distributed/worker/server/worker.go --port 9004
```
3. 启用saver服务，进行存取数据
```
go run crawier_distributed/persist/server/itemsaver.go --port 1234
```
4. 启动主程序，进行分发爬虫任务
```
 go run crawier_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9001,:9002,:9003,:9004"
```