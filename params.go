package quizzee

import "time"

const (
	weightBaidu = 0.3 // Baidu搜索权重
	weightBing  = 0.4 // Bing搜索权重
	weightSogou = 0.6 // Sogou搜索权重
	weight360   = 0.5 // 360搜索权重
)

const Timeout = time.Second * 3 // 搜索超时时间
