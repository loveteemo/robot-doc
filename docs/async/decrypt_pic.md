# 图片解密

## 说明

收到的微信图片是 .dat 格式的文件，需要进行异或解密，下面是 PHP 版本的

## PHP版本

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

## 参考项目

[js解密图片](https://github.com/xiaotuni/node-weixin-image-dat/)
