package parse

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ParseMv(content *http.Response, params map[string]string) engine.ParseResult {

	doc, err := goquery.NewDocumentFromReader(content.Body)

	if err != nil {

	}

	// 相应的model类
	mvModel := model.Mv{}
	result := engine.ParseResult{}

	// 处理相应结果

	// 名称
	doc.Find("#minfo.clearfix div.info h1.font14w").Each(func(i int, selection *goquery.Selection) {
		mvModel.Name = selection.Text()
	})

	// 副标题
	doc.Find("#minfo.clearfix div.info > span:nth-child(3)").Each(func(i int, selection *goquery.Selection) {
		mvModel.Title = selection.Text()
	})

	fmt.Println(mvModel)

	// 把结果返回给engine处理
	result.Items = append(result.Items, mvModel)

	return result

}
