package engine

import "net/http"

type ParseResult struct {
	Requests []Request     //解析器返回一个url对应一个解析器名称
	Items    []interface{} //解析器对应的结果值
}

type Request struct {
	Url         string //下一次要爬虫的地址
	IsDuplicate int    // 是否要添加那到去重里
	Params      map[string]string
	ParserFunc  func(*http.Response, map[string]string) ParseResult //地址内容对应的解析器
}

func NilParser([]byte, map[string]interface{}) ParseResult {
	return ParseResult{}
}
