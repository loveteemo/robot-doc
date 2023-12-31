# 获取自己信息

> 获取自己信息

## 请求参数

| 参数         | 类型      | 描述           | 建议值      |
|------------|---------|--------------|----------|
| para	      | Object  | 数据包          |          |
| > id       | 	String | 	请求时间戳，自行构造  |          |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 6500 |
| > roomid   | 	String | 	群聊id，根据场景变换 | 空        |
| > wxid     | 	String | 	微信id，根据场景变换 | 空        |
| > content  | 	String | 	消息内容        | 空        |
| > nickname | 	String | 	昵称          | 空        |
| > ext      | 	String | 	扩展字段        | 空        |

## 请求示例

```json
{
  "para": {
    "id": "1689039043",
    "type": 6500,
    "roomid": "",
    "wxid": "",
    "content": "",
    "nickname": "",
    "ext": ""
  }
}
```

## 返回示例

```json
{
  "content": "{\"wx_code\":\"\",\"wx_head_image\":\"https://wx.qlogo.cn\",\"wx_id\":\"wxid_1\",\"wx_name\":\"机器人昵称\"}",
  "id": "1689039043",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-12 10:35:41",
  "type": 6500
}
```
