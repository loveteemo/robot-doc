# 心跳

> 收到心跳

## 请求参数

| 参数       | 类型     | 描述      | 
|----------|--------|---------|
| content	 | String | 内容      |
| id       | String | 	时间戳    |
| id1      | String | 	根据场景变换 |
| id2      | String | 	根据场景变换 |
| id3      | String | 根据场景变换  |
| srvid    | String | 	服务ID   |
| time     | String | 	昵称     |
| type     | String | 	5005   |
| wxid     | String | 	发生微信ID |

## 示例

```json
{
  "content": "heart beat",
  "id": "20230713152810",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-13 15:28:10",
  "type": 5005
}
```
