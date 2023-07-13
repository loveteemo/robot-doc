# 群聊用户信息

> 群聊用户信息

## 请求参数

| 参数         | 类型      | 描述           | 建议值      |
|------------|---------|--------------|----------|
| para	      | Object  | 数据包          |          |
| > id       | 	String | 	请求时间戳，自行构造  |          |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 5020 |
| > roomid   | 	String | 	群聊id，根据场景变换 | 群聊ID     |
| > wxid     | 	String | 	微信id，根据场景变换 | 用户ID     |
| > content  | 	String | 	消息内容        | 空        |
| > nickname | 	String | 	昵称          | 空        |
| > ext      | 	String | 	扩展字段        | 空        |

## 请求示例

```json
{
  "para": {
    "id": "1689039043",
    "type": 5020,
    "roomid": "11111@chatroom",
    "wxid": "wxid_1",
    "content": "",
    "nickname": "",
    "ext": ""
  }
}
```

## 返回示例

```json
{
  "content": "{\"nick\":\"wxid_1\",\"roomid\":\"1111@chatroom\",\"wxid\":\"wxid_1\"}",
  "id": "1689039043",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-12 11:05:35",
  "type": 5020
}
```
