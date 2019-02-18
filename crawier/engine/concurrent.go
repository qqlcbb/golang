package engine

type ConcurrentEngine struct{
	Scheducler Scheducler
	WorkCount int
	ItemChan chan Item
	RequestProcessor Processor
}

type Processor func(r Request) (ParseResult, error)

type Scheducler interface {
	ReadyNotifier
	// imterface 函数里面不需要名字
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(sends ...Request) {

	out :=  make(chan ParseResult)

	// 设置输入
	e.Scheducler.Run()

	// 创建10个worker
	for i := 0; i < e.WorkCount; i ++ {
		// 将worker放进worker队列
		e.createWorker(e.Scheducler.WorkChan(), out, e.Scheducler)
	}

	for _, r := range sends {
		// 将请求提交到request队列
		e.Scheducler.Submit(r)
	}
	for {
		// 从chan里接收
		result := <- out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			} ()
		}
		// url去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				// log.Printf("Duplicate request: %s", request.Url)
				continue
			}
			// 继续发送
			e.Scheducler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i m ready
			ready.WorkerReady(in)
			request := <- in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			// 发送出去
			out <- result
		}
	} ()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	// 没有出现过，不会挂掉，return false
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
