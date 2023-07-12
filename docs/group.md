# 获取群聊

> 获取群聊列表


!> 接口待验证

## 请求参数

| 参数         | 类型      | 描述           | 建议值      |
|------------|---------|--------------|----------|
| para	      | Object  | 数据包          |          |
| > id       | 	String | 	请求时间戳，自行构造  |          |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 5000 |
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
    "type": 5000,
    "roomid": "",
    "wxid": "",
    "content": "",
    "nickname": "",
    "ext": ""
  }
}
```

## 返回示例

!> 需要关注的是 `wxid` 这个就行

```json
{
  "content": [
    {
      "headimg": "",
      "name": "朋友推荐消息",
      "node": 184644872,
      "remarks": "",
      "wxcode": "fmessage",
      "wxid": "fmessage"
    }
  ],
  "id": "1689039043",
  "type": 5000
}
```
