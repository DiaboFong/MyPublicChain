package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/json"
)

//将一个Int64的整数:转为二进制后，每8bit一个byte,转为[]byte

func IntToHex(num int64) []byte { //unit8 -->0-255
	buff := new(bytes.Buffer)
	//将二进制数据写入w
	//func Write(w io.Writer,order ByteOrder,data interface{})error
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	//转为[]byte并返回
	return buff.Bytes()
}

/*
JSON解析函数，JSON字符串转成数组
 */

func JSONToArray(jsonString string) []string {
	var arr []string
	err := json.Unmarshal([]byte(jsonString), &arr)
	if err != nil {
		log.Panic(err)
	}
	return arr
}
