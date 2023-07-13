# 发送附件

> 发送附件消息

!> 发文件有一定概率被屏蔽，建议使用发送文件链接的方式

## 请求参数

| 参数         | 类型      | 描述           | 建议值             |
|------------|---------|--------------|-----------------|
| para	      | Object  | 数据包          |                 |
| > id       | 	String | 	请求时间戳，自行构造  |                 |
| > type     | 	Int    | 	根据请求接口进行变更  | 固定值 5003        |
| > roomid   | 	String | 	群聊id，根据场景变换 | 空               |
| > wxid     | 	String | 	微信id，根据场景变换 | 接收人ID(群聊也填写这个值) |
| > content  | 	String | 	消息内容        | 发送文件本地地址        |
| > nickname | 	String | 	昵称          | 空               |
| > ext      | 	String | 	扩展字段        | 空               |

## 请求示例

```json
{
  "para": {
    "id": "1689039043",
    "type": 5003,
    "roomid": "",
    "wxid": "1111@chatroom",
    "content": "C:\\temp\\1.txt",
    "nickname": "",
    "ext": ""
  }
}
```

## 返回示例

```json
{
  "content": "send attatch msg:asm execution succsessed",
  "id": "1689039043",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-12 10:54:13",
  "type": 5003
}
```
