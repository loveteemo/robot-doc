# 环境

## 启动 serve.js

作者源码可以在github上获取，也可以用下面简单的代码来启动一个ws服务器。

```js
// 配置监听端口 配合非注入工具使用 
// 如需提供给外部调用，要服务器开通指定端口
const Host = '127.0.0.1'
const Port = 5555

// 自定义逻辑，用于处理接收到的消息和发送消息
const Notify = 'https://baidu.com/webhook'

const WebSocket = require('ws');
const http = require('http');
const ws = new WebSocket('ws://' + Host + ':' + Port);

// 启动 ws 服务
ws.on('open', function open() {
    console.log('connected');
});

// 关闭 ws 服务
ws.on('close', function close() {
    console.log('disconnected');
});

// 接收到消息 data 为二进制，需要转换成字符串
ws.on('message', function incoming(data) {
    const receiveData = JSON.parse(data);
    //也可以自己实现消息处理，这里因为考虑到扩展，所以我转发到了webhook上
    httpRequest(Notify, 'POST', receiveData).then(r => {
        console.log(r)
    }).catch(e => {
        console.log(e)
    })
    console.log(receiveData);
});

function httpRequest(url, method, data) {
    return new Promise((resolve, reject) => {
        const options = {
            method: method, headers: {
                'Content-Type': 'application/json',
            },
        };

        const req = http.request(url, options, (res) => {
            let responseBody = '';

            res.on('data', (chunk) => {
                responseBody += chunk;
            });

            res.on('end', () => {
                resolve(responseBody);
            });
        });

        req.on('error', (error) => {
            reject(error);
        });

        if (data) {
            req.write(JSON.stringify(data));
        }

        req.end();
    });
}
```

## 启动ws

```bash
node run serve.js
```

项目是node环境启动，也可以启动其他语言的ws来实现。


