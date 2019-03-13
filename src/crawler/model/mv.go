package model

// 音乐Mv相关model
type Mv struct {
	Name       string            //名称
	Alias      []string          //别名
	Title      string            // 标题
	Thumb      string            // 缩略图
	Performer  []interface{}     // 演员
	MovieType  []interface{}     // 类型
	Area       string            // 地区
	Language   []string          // 语言
	Director   string            // 导演
	FilmLength int               // 片长
	Intro      string            // 简介
	ScreenShot string            // 截图
	MVUrl      map[string]string // 电脑版下载地址
	FlatUrl    map[string]string // 平板下载地址
	MobileUrl  map[string]string // 手机下载地址
	CreateTime int               // 创建时间
	UpdateTime int               // 修改时间
}
