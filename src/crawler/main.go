package main

import (
	"crawler/engine"
	"crawler/item"
	"crawler/parse"
	"crawler/scheduler"
	"time"
)

var Restart = make(chan int)

func main() {

	// 定义相应适配器及工作协程数
	e := engine.Engine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    item.ItemSave(),
	}

	//go func() {
	//
	//	for{
	//		// 开启采集 电影
	//		movieParams := make(map[string]string)
	//		movieParams["type"] = "1"
	//		movieParams["page"] = "1"
	//		e.Run(engine.Request{
	//			Url:"https://www.80s.tw/movie/list/-----p/1",   	//采集地址
	//			IsDuplicate:0,					  	// 是否添加到过滤器中
	//			Params:movieParams,
	//			ParserFunc:parse.ParseList,    	// 解析器
	//		})
	//
	//		time.Sleep(86400 * time.Second)
	//	}
	//}()
	//
	//
	//go func() {
	//	for{
	//		// 电视剧
	//		tvParams := make(map[string]string)
	//		tvParams["type"] = "2"
	//		tvParams["page"] = "1"
	//		e.Run(engine.Request{
	//			Url:"https://www.80s.tw/ju/list/----0--p/1",   	//采集地址
	//			IsDuplicate:0,					  	// 是否添加到过滤器中
	//			Params:tvParams,
	//			ParserFunc:parse.ParseList,    	// 解析器
	//		})
	//
	//		time.Sleep(86400 * time.Second)
	//	}
	//}()
	//
	//
	//go func() {
	//	for{
	//		// 动漫
	//		comicParams := make(map[string]string)
	//		comicParams["type"] = "3"
	//		comicParams["page"] = "1"
	//		e.Run(engine.Request{
	//			Url:"https://www.80s.tw/dm/list/----14--p/1",   	//采集地址
	//			IsDuplicate:0,					  	// 是否添加到过滤器中
	//			Params:comicParams,
	//			ParserFunc:parse.ParseList,    	// 解析器
	//		})
	//
	//		time.Sleep(86400 * time.Second)
	//	}
	//}()
	//
	//go func() {
	//	for{
	//		// 综艺
	//		varietyParams := make(map[string]string)
	//		varietyParams["type"] = "4"
	//		varietyParams["page"] = "1"
	//		e.Run(engine.Request{
	//			Url:"https://www.80s.tw/zy/list/----4--p/1",   	//采集地址
	//			IsDuplicate:0,					  	// 是否添加到过滤器中
	//			Params:varietyParams,
	//			ParserFunc:parse.ParseList,    	// 解析器
	//		})
	//
	//		time.Sleep(86400 * time.Second)
	//	}
	//}()

	for {

		// 音乐MV
		mvParams := make(map[string]string)
		mvParams["type"] = "5"
		mvParams["page"] = "1"
		e.Run(engine.Request{
			Url:         "https://www.80s.tw/mv/list/-----p/1", //采集地址
			IsDuplicate: 0,                                     // 是否添加到过滤器中
			Params:      mvParams,
			ParserFunc:  parse.ParseList, // 解析器
		})

		time.Sleep(86400 * time.Second)
	}
}
