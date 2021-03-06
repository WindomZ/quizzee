# quizzee

> 问答题辅助方案 - 开箱即用、推荐引擎、嵌入式存储、RESTful API

[![Build Status](https://travis-ci.org/WindomZ/quizzee.svg?branch=master)](https://travis-ci.org/WindomZ/quizzee)

目的是为各类问答题辅助项目，统一提供`推荐`与`记忆`两个核心服务解决方案，高效且稳定。

提供`两种`接入方式：
- RESTful [API](#api)
- [原生Go用法](#usage)

## Features
- [x] `1问题+N选项`的答题模式
- [x] 推荐`择优`+`答案`记忆
- [x] RESTful [API](#api)
- [x] 4种国内主流[`搜索引擎`](#search-engine)支持
- [x] 5种跨平台嵌入式[`数据库`](#database)支持

## Install
```bash
go get github.com/WindomZ/quizzee/...
```

## API
跨语言方案，详见[API服务和文档](./server)

推荐下面Docker方案

### Docker
构建镜像
```bash
docker build -t quizzee/server .
```

运行容器
```bash
docker run -d -p 8080:8080 --name quizzee-server quizzee/server
```

## Usage
原生Go方案，两种模式支持：
1. [开箱即用](#zero-configuration-usage) - 纯Go实现，快速应用各平台
1. [自行配置](#advanced-usage) - 配置方案，可能需要cgo编译

### Zero configuration usage
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

### Advanced usage
```
import (
	_ "github.com/WindomZ/gcws/jieba"
	"github.com/WindomZ/quizzee"
	_ "github.com/WindomZ/quizzee-db/bolt"
)

// 初始化搜索
quizzee.RegisterCWS("jieba")
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

## Related
- [tools-weight](./tools/weight) - 搜索权重计算工具

## Contributing
欢迎你提交PR。

也可以在[issues](https://github.com/WindomZ/quizzee/issues)汇报Bugs，提出新想法、新要求或者讨论问题。

如果你喜欢这个项目，可以点下 :star: 予以支持！
