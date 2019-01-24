package comment

import (
	"math"
)

/**
分类处理类
*/

type Page struct {
	Count int // 总条数
	Size  int // 每页显示条数
}

type PageShow struct {
	PageNow   int   // 当前页
	PageSize  int   // 显示条数
	PageCount int   // 总页数
	Index     []int // 页数列表
}

var (
	pageInfo *PageShow
)

// 获取分页数据
func (this *Page) ShowPage(pageNow, indexNum int) (pageInfo PageShow) {
	pageInfo.PageNow = pageNow
	pageInfo.PageSize = this.Size

	// 计算总数页
	pageInfo.PageCount = this.getTotalPage()
	pageInfo.Index = this.getIndex(pageNow, indexNum, pageInfo.PageCount)

	return
}

func (this *Page) getTotalPage() (count int) {
	// 总页数
	count = int(math.Ceil(float64(this.Count) / float64(this.Size)))
	return
}

func (this *Page) getIndex(pageNow, indexNum, pageCount int) (pages []int) {

	mediuIndex := int(math.Ceil(float64(indexNum) / float64(2)))

	switch {
	case pageCount > indexNum && pageNow >= mediuIndex:
		pageStart := pageNow - mediuIndex + 1
		countStart := pageCount - indexNum + 1
		start := int(math.Min(float64(pageStart), float64(countStart)))
		pages = make([]int, indexNum)
		for i, _ := range pages {
			pages[i] = start + i
		}
	default:
		pages = make([]int, int(math.Min(float64(indexNum), float64(pageCount))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
	}
	return
}
