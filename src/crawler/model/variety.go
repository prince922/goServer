package model

// 综艺相关model
type Variety struct {
	Name         string              //名称
	Alias        []string            //别名
	Thumb        string              // 缩略图
	Performer    []interface{}       // 演员
	MovieType    []interface{}       // 类型
	UpdateIngStr string              // 更新到的集数
	Area         string              // 地区
	Director     string              // 导演
	ReleaseTime  int                 // 上映时间
	FilmLength   int                 // 片长
	BeanScore    float32             // 豆瓣评分
	BeanUrl      string              // 豆瓣相关评论地址
	Intro        string              // 简介
	ScreenShot   string              // 截图
	MVUrl        []map[string]string // 电脑版下载地址
	FlatUrl      []map[string]string // 平板下载地址
	MobileUrl    []map[string]string // 手机下载地址
	CreateTime   int                 // 创建时间
	UpdateTime   int                 // 修改时间
}
