that project for json-rpc 2.0
package main

import (
	"fmt"
	"json-rpc"
	"net/http"
)
type CalculationService struct {}

func (c * CalculationService)Sum(a,b int)(int) {

	return  a + b
}

func (c * CalculationService)Multiply(a,b int)(int) {

	return a * b
}

func (c * CalculationService)Divide(a,b int)(int) {

	return a / b
}

func CalculationMinus(a,b int)(int) {

	return a - b
}



func main() {
	Json := json_rpc.JsonRpcService{}
	Json.RegisterService("cal",new(CalculationService))
	if err := http.ListenAndServe(":8090",&Json);err != nil {
		fmt.Println("http server exception %s \n",err.Error())
	}

}

test 1 
       request :
               {"jsonrpc":"2.0","id":"222222","method":"cal_divide","params":[100,10]}
       response :
                {
                    "Id": "222222",
                    "JsonRpc": "2.0",
                    "Result": 10
                }
