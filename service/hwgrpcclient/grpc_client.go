/*
   Author: Mr.Huang
*/
package client

import (
	"context"
	"dbservice/service/hwgrpc"
	"fmt"
	"google.golang.org/grpc"
)

/*
	GRCP 客户端调用

	select 方法 rows 带数组
	{"status":"success","rows":"[{"d_desc":"Hx科技"  ,"d_height":40,"d_id":"1681829541653","d_ip":"192.168.0.110","d_name":"路由器01","d_setup_address":"北京市大安小区15栋2单101室","d_width":200,"d_writetime":"2023-04-18 22:52:21"},{"d_desc":"Hx科技","d_height":4  0,"d_id":"1681830496354","d_ip":"192.168.0.110","d_name":"路由器01",   "d_setup_address":"北京市大安小区15栋2单101室","d_width":120,"d_writetime ":"2023-04-18 23:08:16"}]"}
	insert/delete/update 方法 rows 是对象
	{"status":"success","rows":{"lastInsertId":0,"rowsAffected":1}}

*/
func CallCommonLogic() (string, error) {
	//1、Dail连接
	conn, err := grpc.Dial(":8089", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	serviceClient := hwgrpc.NewCommonServiceClient(conn)
	example := `
{
				"id":"1",
				"method":"sel",
				"cols":"*",
				"where":" and d_ip like '192.168.%' ",
				
				"pageSize":10,
				"currentPage":0
			}
`
	example = `
	{
					"id":"3",
					"method":"ins",
					"cols":"d_id,d_name,d_setup_address,d_ip,d_desc,d_width,d_height,d_writetime",
					"values":"'${timestamp13}','路由器04','天津市满源小区1栋12单101室','192.168.1.110','Hx科技A',230,140,'${YYYY-MM-DD hh:mm:ss}' ",
	
				}
	`
	//inputRequest := &hwgrpc.CommonInputRequest{InputJsonString: "{\"name\": \"201907300002\", \"age\":\"" + strconv.Itoa(int(time.Now().Unix())) + "\",\"address\":\"北京\"}"}
	inputRequest := &hwgrpc.CommonInputRequest{InputJsonString: example}
	outputResponse, err := serviceClient.CallCommonMethod(context.Background(), inputRequest)
	if err != nil {
		fmt.Println("调用出错：", err)
		return "", err
	} else {
		fmt.Println("服务器返回：", outputResponse)
	}
	//fmt.Println("rows:", gjson.Parse(outputResponse.OutputJsonString).Get("rows").String())
	return outputResponse.OutputJsonString, nil
}
