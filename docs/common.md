# 请求地址

因为每个接口的请求地址都是相同的，所以在这里统一说明一下。请求服务器IP+端口号

如果是本地开发，IP为`localhost`，或者 `127.0.0.1`, 端口号为`5555`，则请求地址为`http://localhost:5555`

## 请求参数

!> 所有接口都需要传递的参数,可以为 `null`,但是不能不传递

!> id 类型为 `string`

| 参数         | 类型      | 描述           |
|------------|---------|--------------|
| para	      | Object  | 数据包          |
| > id       | 	String | 	请求时间戳，自行构造  |
| > type     | 	Int    | 	根据请求接口进行变更  |
| > roomid   | 	String | 	群聊id，根据场景变换 |
| > wxid     | 	String | 	微信id，根据场景变换 |
| > content  | 	String | 	消息内容        |
| > nickname | 	String | 	昵称          |
| > ext      | 	String | 	扩展字段        |

## 返回参数

| 参数       | 类型      | 描述      |
|----------|---------|---------|
| content  | 	Object | 	返回数据   |
| id       | 	String | 	时间戳    |
| type     | 	Int    | 	根据请求返回 |
| receiver | String  | 接收者     |
| sender   | String  | 发送者     |
| srvid    | Int     | 服务器id   |
| status   | String  | 状态      |
| time     | String  | 时间      |
