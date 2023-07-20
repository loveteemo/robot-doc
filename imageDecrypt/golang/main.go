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
