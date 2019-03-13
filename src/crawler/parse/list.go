package parse

import (
	"crawler/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func ParseList(content *http.Response, params map[string]string) engine.ParseResult {
	doc, err := goquery.NewDocumentFromReader(content.Body)
	result := engine.ParseResult{}
	if err != nil {
		return result
	}

	var funcStr func(*http.Response, map[string]string) engine.ParseResult
	var dup int
	var nextList string

	if params["type"] == "1" {
		funcStr = ParseMovie
		dup = 1
		nextList = movieHost
	} else if params["type"] == "2" {
		funcStr = ParseTv
		dup = 0
		nextList = tvHost
	} else if params["type"] == "3" {
		funcStr = ParseComic
		dup = 0
		nextList = comicHost
	} else if params["type"] == "4" {
		funcStr = ParseVariety
		dup = 0
		nextList = varietyHost
	} else if params["type"] == "5" {
		funcStr = ParseMv
		dup = 1
		nextList = mvHost
	}

	doc.Find("#block3.left div.clearfix.noborder.block1 ul.clearfix li a").Each(func(i int, selection *goquery.Selection) {
		hrefUrl, ret := selection.Attr("href")
		if ret {
			url := host + hrefUrl
			result.Requests = append(result.Requests, engine.Request{
				Url:         url,
				IsDuplicate: dup,
				Params:      params,
				ParserFunc:  funcStr,
			})
		}

	})

	// 判断是否有下一页
	lastPage, ret := doc.Find(".pager a").Last().Attr("href")

	if ret {
		pageArr := strings.Split(lastPage, "/")
		maxPage := pageArr[len(pageArr)-1]
		maxPageInt, _ := strconv.Atoi(maxPage)
		page, _ := strconv.Atoi(params["page"])

		if maxPageInt > page {
			nextPage := strconv.Itoa(page + 1)
			nextUrl := nextList + string(nextPage)
			params["page"] = nextPage
			fmt.Println(nextUrl)
			result.Requests = append(result.Requests, engine.Request{
				Url:         nextUrl,
				IsDuplicate: 0,
				Params:      params,
				ParserFunc:  ParseList,
			})
		}

	}

	return result

}
