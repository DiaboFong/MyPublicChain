package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
	"encoding/json"
)

/*
将一个int64的整数：转为二进制后，每8bit一个byte。转为[]byte
 */
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	//将二进制数据写入w
	//func Write(w io.Writer, order ByteOrder, data interface{}) error
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	//转为[]byte并返回
	return buff.Bytes()
}

/*
json解析的的函数

*/
func JSONToArray(jsonString string) []string {
	var arr [] string
	err := json.Unmarshal([]byte(jsonString), &arr)
	if err != nil {
		log.Panic(err)
	}
	return arr
}


//字节数组反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}