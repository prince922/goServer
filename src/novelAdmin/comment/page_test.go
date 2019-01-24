package comment

import (
	"testing"
)

func TestPage_ShowPage(t *testing.T) {
	page := &Page{100, 5}

	for i := 6; i < 20; i++ {
		page.ShowPage(i, 10)
	}

}
