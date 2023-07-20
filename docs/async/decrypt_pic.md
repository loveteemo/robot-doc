# 图片解密

## 说明

收到的微信图片是 `.dat` 格式的文件，需要进行异或解密，下面是各种语言的版本

## php 版本

```php
<?php

$file_path = 'file.dat'; // 替换为实际的.dat文件路径

// 读取.dat文件内容到缓冲区
$buffer = file_get_contents($file_path);

// 根据buffer内容获取文件后缀
$SuffixMap = [
    'ffd8ffe000104a464946' => 'jpg',
    '89504e470d0a1a0a0000' => 'png',
    '47494638396126026f01' => 'gif',
    '49492a00227105008037' => 'tif',
    '424d228c010000000000' => 'bmp',
    '424d8240090000000000' => 'bmp',
    '424d8e1b030000000000' => 'bmp'
];

// 确保读取成功
if ($buffer !== false) {
    // 打印文件流 二进制
    // var_dump($buffer);

    //异或运算拿到值 和 后缀
    $xor1 = getXOR($buffer, $SuffixMap);
    if (!$xor1) {
        echo "失败";
        die;
    }

    //转换之后的文件流
    $newBuffer = '';
    //遍历文件流
    $bufferArray = str_split($buffer);
    foreach ($bufferArray as $value) {
        //异或运算
        $newBuffer .= chr(ord($value) ^ hexdec('0x' . $xor1['value']));
    }
    // var_dump($newBuffer);

    //保存文件
    $pathInfo = pathinfo($file_path);
    $saveFileName = __DIR__ . '/' . $pathInfo['filename'] . '.' . $xor1['suffix'];
    file_put_contents($saveFileName, $newBuffer);

} else {
    // 处理读取失败的情况
    echo '读取文件失败';
    die;
}

function getXOR($fileBuffer, $SuffixMap): ?array
{
    $sMapList = array_keys($SuffixMap);
    foreach ($sMapList as $key) {
        $suffix = $SuffixMap[$key];
        $hex = [substr($key, 0, 2), substr($key, 2, 2), substr($key, 4, 2)];
        $map = [];
        for ($a = 0; $a < 3; $a++) {
            $byte = ord($fileBuffer[$a]);
            $value = $byte ^ hexdec($hex[$a]);
            $map[$a] = $value;
        }
        if ($map[0] == $map[1] && $map[1] === $map[2]) {
            return ["value" => dechex($map[0]), "suffix" => $suffix];
        }
    }
    return null;
}

```

## python 版本

QQ群友：`QQ 364345866` 提供

```python
import os
import binascii

# 替换为实际的.dat文件路径
file_path = 'G:/WeChatFiles//WeChat Files/xxxwxid/FileStorage/MsgAttach/983e528fb31e20bd89d8f1bd531a6d5f/Thumb/2022-07/2022-07-13 14_09_21_852.dat'

# 后缀映射
suffix_map = {
    'ffd8ffe000104a464946': 'jpg',
    '89504e470d0a1a0a0000': 'png',
    '47494638396126026f01': 'gif',
    '49492a00227105008037': 'tif',
    '424d228c010000000000': 'bmp',
    '424d8240090000000000': 'bmp',
    '424d8e1b030000000000': 'bmp'
}

def get_xor(file_buffer, suffix_map):
    for key in suffix_map.keys():
        suffix = suffix_map[key]
        hex_values = [key[i:i+2] for i in range(0, len(key), 2)]
        map_values = []
        for a in range(3):
            byte = file_buffer[a]
            value = byte ^ int(hex_values[a], 16)
            map_values.append(value)
        if map_values[0] == map_values[1] == map_values[2]:
            return {"value": format(map_values[0], 'x'), "suffix": suffix}
    return None

# 读取.dat文件内容到缓冲区
with open(file_path, 'rb') as file:
    buffer = file.read()

# 确保读取成功
if buffer:
    xor1 = get_xor(buffer, suffix_map)
    if not xor1:
        print("失败")
        exit()

    # 转换之后的文件流
    new_buffer = bytearray()
    # 遍历文件流
    for value in buffer:
        # 异或运算
        new_buffer.append(value ^ int('0x' + xor1['value'], 16))

    # 保存文件
    save_file_name = os.path.join(os.path.dirname(file_path), os.path.splitext(os.path.basename(file_path))[0] + '.' + xor1['suffix'])
    with open(save_file_name, 'wb') as file:
        file.write(new_buffer)
else:
    # 处理读取失败的情况
    print('读取文件失败')

```

## golang 版本

QQ群友：`QQ 364345866` 提供

```go
package main

import (
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
)

var suffixMap = map[string]string{
	"ffd8ffe000104a464946": "jpg",
	"89504e470d0a1a0a0000": "png",
	"47494638396126026f01": "gif",
	"49492a00227105008037": "tif",
	"424d228c010000000000": "bmp",
	"424d8240090000000000": "bmp",
	"424d8e1b030000000000": "bmp",
}

func getXOR(fileBuffer []byte, suffixMap map[string]string) (string, string, bool) {
	for key, suffix := range suffixMap {
		hexValues := []byte(key)
		mapValues := make([]byte, 3)
		for a := 0; a < 3; a++ {
			value, _ := hex.DecodeString(string(hexValues[a*2 : a*2+2]))
			mapValues[a] = fileBuffer[a] ^ value[0]
		}
		if mapValues[0] == mapValues[1] && mapValues[1] == mapValues[2] {
			return hex.EncodeToString(mapValues[:1]), suffix, true
		}
	}
	return "", "", false
}

func main() {
	// 替换为实际的.dat文件路径 wxid替换自己的
	filePath := "G:/WeChatFiles//WeChat Files/wxid/FileStorage/MsgAttach/983e528fb31e20bd89d8f1bd531a6d5f/Image/2022-08/7a1f17ebccf3c61e17ccdecd971d3b11.dat"

	// 读取.dat文件内容到缓冲区
	fileBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		// 处理读取失败的情况
		panic("读取文件失败")
	}

	// 确保读取成功
	xorValue, suffix, ok := getXOR(fileBuffer, suffixMap)
	if !ok {
		panic("失败")
	}

	// 转换之后的文件流
	newBuffer := make([]byte, len(fileBuffer))
	// 遍历文件流
	xorByte, _ := hex.DecodeString(xorValue)
	for i, value := range fileBuffer {
		// 异或运算
		newBuffer[i] = value ^ xorByte[0]
	}

	// 保存文件
	dir, filename := filepath.Split(filePath)
	saveFileName := filepath.Join(dir, filename+"."+suffix)
	err = ioutil.WriteFile(saveFileName, newBuffer, 0644)
	if err != nil {
		panic("保存文件失败")
	}
}
```

## 参考项目

[js解密图片](https://github.com/xiaotuni/node-weixin-image-dat/)
