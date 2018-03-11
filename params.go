package quizzee

import "time"

const (
	weightBaidu = 0.969039 // Baidu搜索权重
	weightBing  = 1.000000 // Bing搜索权重
	weightSogou = 0.836369 // Sogou搜索权重
	weight360   = 0.841272 // 360搜索权重
)

const Timeout = time.Second * 3 // 搜索超时时间
