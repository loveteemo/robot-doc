# 收到@文本消息

> 收到@文本消息


!> 只能通过昵称来判断是否@你，等作者扩展

## 请求参数

| 参数       | 类型     | 描述      | 
|----------|--------|---------|
| content	 | String | 内容      |
| id       | String | 	时间戳    |
| id1      | String | 	发消息人ID |
| id2      | String | 	根据场景变换 |
| id3      | String | 根据场景变换  |
| srvid    | String | 	服务ID   |
| time     | String | 	昵称     |
| type     | String | 	1      |
| wxid     | String | 	发生微信ID |

## 示例

```json
{
  "content": "@robot🎉 你 好",
  "id": "20230713152538",
  "id1": "wxid_1",
  "id2": "",
  "id3": "",
  "srvid": 1,
  "time": "2023-07-13 15:25:38",
  "type": 1,
  "wxid": "11112@chatroom"
}
```
