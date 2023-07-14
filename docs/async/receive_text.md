# 收到文本消息

> 收到文本消息

## 使用自定义回复示例

前面提到过，当开启 websocket 服务时，可以收到消息的异步转发。

当异步转发的环境收到自定义消息，可以根据开发对应的自定义回复内容。

这里提供一个PHP示例，当收到文本消息时，如果内容为“签到”，则回复“签到成功”。

可以适当添加 nosql,mysql 作为配置项，实现更多的自定义回复。

!> 请勿用作其他用途，仅供学习参考。

```php
<?php

$data = $_POST;
if (isset($data["type"]) && $data["type"] == 1) { //私聊信息
    if ($data["content"] == "签到") {
        $sendData = [
            "para" => [
                "id" => (string)time(),
                "type" => 555,
                "roomid" => $data["wxid"],//回复的人
                "wxid" => $data["wxid"],//回复的人
                "content" => "签到成功",//回复的消息
                "nickname" => "null",
                "ext" => "null"
            ]
        ];
    }
}

function sendRequest($api, $data): array
{
    $curl = curl_init();
    curl_setopt_array($curl, array(
        CURLOPT_URL => $api,
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_ENCODING => '',
        CURLOPT_MAXREDIRS => 10,
        CURLOPT_TIMEOUT => 0,
        CURLOPT_FOLLOWLOCATION => true,
        CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
        CURLOPT_CUSTOMREQUEST => 'POST',
        CURLOPT_POSTFIELDS => json_encode($data, JSON_UNESCAPED_UNICODE),
        CURLOPT_HTTPHEADER => array(
            'Content-Type: application/json',
        ),
    ));
    $response = curl_exec($curl);
    curl_close($curl);
    return json_decode($response, true) ?? [];
}
```

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
| type     | String | 	1      |
| wxid     | String | 	发生微信ID |

## 示例

```json
{
  "content": "hello",
  "id": "20230711154020",
  "id1": "",
  "id2": "",
  "id3": "",
  "srvid": 1,
  "time": "2023-07-11 15:40:20",
  "type": 1,
  "wxid": "wxid_1"
}
```
