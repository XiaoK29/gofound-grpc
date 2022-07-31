package pagination

import (
	"math"
)

type Pagination struct {
	Limit     int32 // 限制大小
	PageCount int32 // 总页数
	Total     int32 // 总数据量
}

func (p *Pagination) Init(limit, total int32) {
	p.Limit = limit

	//计算总页数

	p.Total = total

	pageCount := math.Ceil(float64(total) / float64(limit))
	p.PageCount = int32(pageCount)

}

func (p *Pagination) GetPage(page int32) (s, e int32) {
	//获取指定页数的数据
	if page > p.PageCount {
		page = p.PageCount
	}
	if page < 0 {
		page = 1
	}

	//从1开始
	page -= 1

	//计算起始位置
	start := page * p.Limit
	end := start + p.Limit

	if start > p.Total {
		return 0, p.Total - 1
	}
	if end > p.Total {
		end = p.Total
	}

	return start, end

}
