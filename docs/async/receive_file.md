# 收到文件消息

> 收到文件消息

## 请求参数

| 参数        | 类型     | 描述            | 
|-----------|--------|---------------|
| content	  | Object | 内容            |
| > content | String | 消息内容 xml 加密格式 |
| > id1     | String | 群聊ID          |
| > id2     | String | 发消息人ID        |
| id        | String | 时间戳           |
| receiver  | String | 接收人ID         |
| sender    | String | 发送人ID         |
| srvid     | String | 服务ID          |
| status    | String | 状态            |
| time      | String | 时间            |
| type      | String | 49            |
| wxid      | String | 发生微信ID        |

## 示例

```json
{
  "content": {
    "content": "{{xml数据}}",
    "id1": "39036759383@chatroom",
    "id2": "long77591"
  },
  "id": "20230713153831",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-13 15:38:31",
  "type": 49
}
```

xml 转义出来后是下面的内容

```json
{
  "fromusername": "wxid_1",
  "scene": "0",
  "commenturl": "",
  "appmsg": {
    "title": "file_name",
    "des": "",
    "action": "view",
    "type": "6",
    "showtype": "0",
    "content": "",
    "url": "",
    "dataurl": "",
    "lowurl": "",
    "lowdataurl": "",
    "recorditem": "",
    "thumburl": "",
    "messageaction": "",
    "md5": "6dcd6ca7562c6bf1bb0348340f497dd1",
    "extinfo": "",
    "sourceusername": "",
    "sourcedisplayname": "",
    "commenturl": "",
    "appattach": {
      "totallen": "89",
      "attachid": "",
      "emoticonmd5": "",
      "fileext": "json",
      "fileuploadtoken": "v1_u+9RA08RB2TpvXausqMZzsMQOvspAYp43/1W23BuxiSf/GJuuLROjh/+KSjkrbt6tyVFTlWGitlia5c7bmP6DEMiV1q134F9Heg1o238LNJZmcrJTrQ3t26v3dTexpP4npU4UnXC7ik1KVDqi1df8e5JeZJadF8rC9IGMxCz84fpYK64EmtEEirINMjSRmfx07DCoLZymA==",
      "aeskey": ""
    },
    "weappinfo": {
      "pagepath": "",
      "username": "",
      "appid": "",
      "appservicetype": "0"
    },
    "websearch": "",
    "@appid": "",
    "@sdkver": "0"
  },
  "appinfo": {
    "version": "1",
    "appname": "Window wechat"
  }
}
```
