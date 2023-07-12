# 发送@文本消息

> 发送@文本消息

## 请求参数

| 参数         | 类型      | 描述           | 建议值                       |
|------------|---------|--------------|---------------------------|
| para	      | Object  | 数据包          |                           |
| > id       | 	String | 	请求时间戳，自行构造  |                           |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 550                   |
| > roomid   | 	String | 	群聊id，根据场景变换 | 群聊ID                      |
| > wxid     | 	String | 	微信id，根据场景变换 | @谁写谁 <br/>所有人写 notify@all |
| > content  | 	String | 	消息内容        | 发送文本消息                    |
| > nickname | 	String | 	昵称          | @的昵称                      |
| > ext      | 	String | 	扩展字段        | 空                         |

## 请求示例

```json
{
  "para": {
    "id": "1689039043",
    "type": 550,
    "roomid": "21200044202@chatroom",
    "wxid": "notify@all",
    "content": "hello",
    "nickname": "所有人",
    "ext": ""
  }
}
```

## 返回示例

```json
{
  "content": "send at msg : execution succeeded",
  "id": "1689039043",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-12 11:00:18",
  "type": 550
}
```
