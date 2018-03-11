# API

提供快捷1对N接入方案，中心化数据的API服务

## Usage
```bash
./run xxx ../data/xxx.db
```

- 第一参数：数据库表名
- 第二参数：数据文件位置

## API
非常简单，只需两种接口：读/写

### Ask(提问)
- 入参：`question`(问题)和`options`(选项)
- 出参：推荐的`answer`(答案)

请求：
```bash
curl -X POST \
  http://127.0.0.1:8080/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"question":"手机生产商诺基亚最初是以生产什么为主？","options":["耳机", "纸", "杂货"]}'
```

响应：
```json
{
    "errcode": 0,
    "errmsg": "",
    "data": {
        "question": "手机生产商诺基亚最初是以生产什么为主？",
        "options": [
            "耳机",
            "纸",
            "杂货"
        ],
        "answer": "纸",
        "accuracy": 0.8399749788424035
    }
}
```

### Answer(记忆)
- 入参：`question`(问题)、`options`(选项)和`answer`(答案)
- 出参：存储结果

请求：
```bash
curl -X POST \
  http://127.0.0.1:8080/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"question":"手机生产商诺基亚最初是以生产什么为主？","options":["耳机", "纸", "杂货"],"answer": "纸"}'
```

响应：
```json
{
    "errcode": 0,
    "errmsg": "",
    "data": {
        "question": "手机生产商诺基亚最初是以生产什么为主？",
        "options": [
            "耳机",
            "纸",
            "杂货"
        ],
        "answer": "纸",
        "accuracy": 1
    }
}
```
