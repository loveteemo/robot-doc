# 收到文本消息

> 收到文本消息

## 请求参数

| 参数        | 类型     | 描述      | 
|-----------|--------|---------|
| content	  | Object | 内容      |
| > content | String | 消息内容    |
| > id      | String | 	群聊ID   |
| id        | String | 	时间戳    |
| srvid     | String | 	服务ID   |
| status    | String | 	状态     |
| time      | String | 	昵称     |
| type      | String | 	10000  |
| wxid      | String | 	发生微信ID |

## 示例

```json
{
  "content": {
    "content": "a邀请你和b加入了群聊",
    "id1": "11112@chatroom"
  },
  "id": "20230713152537",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-13 15:25:37",
  "type": 10000
}
```
