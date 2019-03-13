package engine

import (
	"crawler/fetch"
	"fmt"
	"golang.org/x/net/context"
	"log"
	"time"
)

type Engine struct {
	Scheduler   Scheduler        // 适配器
	WorkerCount int              // 工作协程
	ItemChan    chan interface{} // 爬取内容
}

type Scheduler interface {
	ReadyNotifilter
	//传入request方法
	Submit(Request)

	Run()
	WorkerChan() chan Request
}

type ReadyNotifilter interface {
	//传入线程对应的接收request 的channel
	WorkerReady(chan Request)
}

func (e *Engine) Run(seed ...Request) {

	// 管道
	out := make(chan ParseResult)

	// 任务适配器
	e.Scheduler.Run()

	//启动worker
	for i := 0; i < e.WorkerCount; i++ {
		//
		createWork(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//把resquest放入到channel里面(去重)
	for _, r := range seed {
		if isDuplicate(r) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	//循环遍历结果（处理结果）
	for {

		cxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)

		defer cancel()

		select {
		case <-cxt.Done():
			fmt.Println("tidy result is over")

			goto Loop

		case result := <-out:

			for _, item := range result.Items {
				//log.Printf("Got item:#%d,%v\n",itemCount,item)
				//itemCount++

				go func() {
					// 将爬取结果写入管道
					e.ItemChan <- item
				}()
			}

			// 把下一次采集的页面放到适配器中
			for _, request := range result.Requests {

				if isDuplicate(request) {
					continue
				}

				e.Scheduler.Submit(request)
			}
		}

	}

Loop:
}

/**
（创建工作协程，开始采集任务）
启动工作任务
*/
func createWork(in chan Request, out chan ParseResult, s ReadyNotifilter) {
	go func() {
		for {

			// 把任务分到工作管道进行分配
			s.WorkerReady(in)

			cxt, cancel := context.WithTimeout(context.Background(), 60*time.Second)

			defer cancel()

			select {
			case <-cxt.Done():
				goto Loop
			case request := <-in:
				// 开始采集处理
				result, err := worker(request)
				if err != nil {
					continue
				}

				// 返回的结果
				out <- result
			}
		}

	Loop:
	}()

}

/**
 * 	工作协程（获取网页内容，解析相应的结果）
 *
 */
func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s", r.Url)

	//调用fetch 去获取网页内容
	body, err := fetch.Fetch(r.Url)

	if err != nil {
		log.Printf("fetcher:error"+"fetcher.url %s,%v", r.Url, err)
		return ParseResult{}, err
	}
	//把requests加入 队列里面
	return r.ParserFunc(body, r.Params), err
}
