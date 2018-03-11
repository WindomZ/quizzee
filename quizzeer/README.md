# quizzeer

> quizzee的推荐实现方案 - 开箱即用

## Usage
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
