# 图片解密

## 说明

收到的微信图片是 .dat 格式的文件，需要进行异或解密，下面是 PHP 版本的

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

## 参考项目

[js解密图片](https://github.com/xiaotuni/node-weixin-image-dat/)
