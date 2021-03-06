# API

> quizzee - 高效的**1对N**问答题API服务

已在`server.go`中实现了一套Go服务。如何运行下面有说明。

## API
用法非常简单，客户端只需实现三种：[检测](#ping检测)/[提问](#ask提问)/[记忆](#answer(记忆))

### Ping(检测)
用于检测发现服务

请求：
```bash
curl -X POST \
  http://127.0.0.1:8080/ping \
  -H 'cache-control: no-cache'
```

响应：
```json
{
    "errcode": 0,
    "errmsg": ""
}
```

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

## Usage
- 虚拟容器[Docker](#docker)(推荐)
- 命令行([CLI](#cli))

### Docker
构建镜像
```bash
docker build -t quizzee/server .
```

运行容器
```bash
docker run -d -p 8080:8080 --name quizzee-server quizzee/server
```

### CLI
命令行格式
```bash
./run $1 $2
```
例如：
```bash
./run xxx ../data/xxx.db
```

- 第一参数：数据库表名
- 第二参数：数据文件位置
