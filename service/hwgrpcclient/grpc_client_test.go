/*
   Author: Mr.Huang
*/
package client

import (
	"fmt"
	"testing"
)

func Test_callhwgrpc(t *testing.T) {
	res, err := CallCommonLogic()
	fmt.Println("res:", res, ",err:", err)
}
