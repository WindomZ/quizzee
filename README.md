# quizzee

> 问答题辅助Go方案 - 开箱即用、推荐引擎、嵌入式存储

[![Build Status](https://travis-ci.org/WindomZ/quizzee.svg?branch=master)](https://travis-ci.org/WindomZ/quizzee)

## Features
- [x] `1问题+N选项`的答题模式
- [x] [`quizzeer`](#usage(zero-configuration))零配置支持
- [x] 4种国内[`搜索引擎`](#search-engine)支持
- [x] 5种嵌入式[`数据库`](#database)支持

## Install
```bash
go get github.com/WindomZ/quizzee/...
```

两种模式支持：
1. [开箱即用](#usage(zero-configuration)) - 纯Go实现，快速应用
1. [自行配置](#advanced-usage) - 配置方案，可能需要cgo支持

## Usage(zero configuration)
```
import "github.com/WindomZ/quizzee/quizzeer"

// 初始化数据库
quizzeer.RegisterDB("testing", "../data/data.db") // 配置表名和文件路径
...

// 推荐答案
recommend, rates := quizzeer.Recommend(
    "手机生产商诺基亚最初是以生产什么为主？", // 问题
    []string{"耳机", "纸", "杂货"},      // 回答选项
)
// recommend int       最佳选项序号
// rates     []float64 各选项权重比
...

// 存储结果
quizzeer.Mark(
    "手机生产商诺基亚最初是以生产什么为主？",
    []string{"耳机", "纸", "杂货"},
    "纸",
)
```

## Advanced usage
```
import (
	_ "github.com/WindomZ/gcws/sego"
	"github.com/WindomZ/quizzee"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

// 初始化搜索
quizzee.RegisterCWS("sego")
// 初始化数据库
quizzee.RegisterDB("testing", "../data/data.db") // 配置表名和文件路径
...

// 推荐答案
recommend, rates := quizzee.Recommend(
    "手机生产商诺基亚最初是以生产什么为主？", // 问题
    []string{"耳机", "纸", "杂货"},      // 回答选项
)
// recommend int       最佳选项序号
// rates     []float64 各选项权重比
...

// 存储结果
quizzee.Mark(
    "手机生产商诺基亚最初是以生产什么为主？",
    []string{"耳机", "纸", "杂货"},
    "纸",
)
```

## Search engine
- Baidu
- Bing
- Sogou
- 360

## Database
详见[quizzee-db](https://github.com/WindomZ/quizzee-db) - 相关数据库支持

## Contributing
欢迎你提交PR。

也可以在[issues](https://github.com/WindomZ/quizzee/issues)汇报Bugs，提出新想法、新要求或者讨论问题。

如果你喜欢这个项目，可以点下 :star: 予以支持！

## License
[MIT](https://github.com/WindomZ/quizzee/blob/master/LICENSE)
