package utility

import (
	"2408_middleware-stress-test/global"
	"fmt"

	"github.com/moov-io/iso8583"
)

func GetSampleLogon() ([]byte, string) {
	tid := getRandomNum(8, 99999998)
	message := iso8583.NewMessage(Spec87RK)
	err := message.Unpack(global.DataLogon[2:])
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}, tid
	}
	_ = message.Field(11, getRandomNum(6, 999998))
	_ = message.Field(41, tid)
	_ = message.Field(42, getRandomNum(15, 99999998))
	msg, err := message.Pack()
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}, tid

	}

	newMsg := append([]byte{}, global.DataLogon[:2]...)
	newMsg = append(newMsg, msg...)

	return newMsg, tid
}
