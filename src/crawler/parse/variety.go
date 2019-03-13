package parse

import (
	"crawler/engine"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ParseVariety(content *http.Response, params map[string]string) engine.ParseResult {

	doc, err := goquery.NewDocumentFromReader(content.Body)

	if err != nil {

	}

	result := engine.ParseResult{}

	doc.Find("#block3.left div.clearfix.noborder.block1 ul.clearfix li a").Each(func(i int, selection *goquery.Selection) {
		//hrefUrl,ret := selection.Attr("href")

		//if ret {
		//	url := host + hrefUrl
		//
		//}

	})

	return result

}
