# 发送文本消息

> 发送文本消息

## 请求参数

| 参数         | 类型      | 描述           | 建议值             |
|------------|---------|--------------|-----------------|
| para	      | Object  | 数据包          |                 |
| > id       | 	String | 	请求时间戳，自行构造  |                 |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 555         |
| > roomid   | 	String | 	群聊id，根据场景变换 | 空               |
| > wxid     | 	String | 	微信id，根据场景变换 | 接收人ID(群聊也填写这个值) |
| > content  | 	String | 	消息内容        | 发送文本消息          |
| > nickname | 	String | 	昵称          | 空               |
| > ext      | 	String | 	扩展字段        | 空               |

## 请求示例

```json
{
  "para": {
    "id": "1689039043",
    "type": 555,
    "roomid": "",
    "wxid": "wxid_1",
    "content": "hello",
    "nickname": "",
    "ext": ""
  }
}
```

## 返回示例

```json
{
  "content": "send text message : execution succsessed",
  "id": "1689039043",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-12 10:56:33",
  "type": 555
}
```
