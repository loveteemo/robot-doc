# 收到图片消息

> 收到图片消息

## 请求参数

| 参数        | 类型     | 描述            | 
|-----------|--------|---------------|
| content	  | Object | 内容            |
| > content | String | 消息内容 xml 加密格式 |
| > detail  | String | 文件存放位置(加密)    |
| > id1     | String | 群聊ID          |
| > id2     | String | 发消息人ID        |
| > thumb   | String | 缩略图(加密)       |
| id        | String | 时间戳           |
| receiver  | String | 接收方           |
| sender    | String | 发送方           |
| srvid     | String | 服务ID          |
| status    | String | 状态            |
| time      | String | 时间            |
| type      | String | 3             |

## 示例

```json
{
  "content": {
    "content": "{{xml}}",
    "detail": "wxid_2\\FileStorage\\Image\\2023-07\\d014baf0b74de7ebebbe609756db9caf.dat",
    "id1": "1111@chatroom",
    "id2": "wxid_1",
    "thumb": "wxid_2\\FileStorage\\Image\\Thumb\\2023-07\\81f390d75949c7c1531e0212c0fc3dab_t.dat"
  },
  "id": "20230713152041",
  "receiver": "CLIENT",
  "sender": "SERVER",
  "srvid": 1,
  "status": "SUCCSESSED",
  "time": "2023-07-13 15:20:41",
  "type": 3
}
```

xml 解析出来

```json
{
  "img": {
    "@aeskey": "4a4701b190af8d89a2aac44a49fed952",
    "@encryver": "1",
    "@cdnthumbaeskey": "4a4701b190af8d89a2aac44a49fed952",
    "@cdnthumburl": "3057020100044b304902010002045c1e9b5002032e6bfd02043b48aa3d020464afa5c8042437363361376663662d386533392d346362362d623438382d6264376235336163343939610204011418020201000405004c54a100",
    "@cdnthumblength": "9579",
    "@cdnthumbheight": "67",
    "@cdnthumbwidth": "120",
    "@cdnmidheight": "0",
    "@cdnmidwidth": "0",
    "@cdnhdheight": "0",
    "@cdnhdwidth": "0",
    "@cdnmidimgurl": "3057020100044b304902010002045c1e9b5002032e6bfd02043b48aa3d020464afa5c8042437363361376663662d386533392d346362362d623438382d6264376235336163343939610204011418020201000405004c54a100",
    "@length": "45777",
    "@md5": "caba490c74315184523d264273aa9fef"
  },
  "platform_signature": "",
  "imgdatahash": ""
}
```
