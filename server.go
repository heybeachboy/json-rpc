package json_rpc

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const SERVICEANDMETHODSEPARATOR = `_`

/**
 * json rpc service implements
 */

type JsonRpcService struct {
	services ServiceMap
	//methods   MethodMap
	callbacks CallbackMap
}

/**
 * register service
 */
func (j *JsonRpcService) RegisterService(serviceName string, service interface{}) {

	if j.services == nil {
		j.services = make(ServiceMap)
	}

	if j.judgeServiceIsExist(serviceName) {
		fmt.Errorf("service name already register : %s", serviceName)
		return
	}

	var svr Service
	svr.ServerName = FormatName(serviceName)
	svr.ServiceTyp = reflect.TypeOf(service)
	svr.ServiceVal = reflect.ValueOf(service)
	resp := j.reflectCallback(svr)

	if !resp {
		fmt.Errorf("register service  %s : method is null", svr.ServerName)
		return
	}
	j.services[svr.ServerName] = &svr

}

//func(j *JsonRpcService)

/**
 *reflect service callback function
 */

func (j *JsonRpcService) reflectCallback(service Service) (bool) {

	count := 0
	for i := 0; i < service.ServiceTyp.NumMethod(); i++ {
		var c Callback
		method := service.ServiceTyp.Method(i)
		mTyp := method.Type
		name := FormatName(method.Name)
		c.ServiceName = service.ServerName
		c.Method = method
		c.MethodTyp = mTyp
		c.MethodName = service.ServerName + `.` + name
		if j.judgeMethodIsExist(c.MethodName) {
			fmt.Errorf("service of method already register : %s", c.MethodName)
			continue
		}
		count++

		if j.callbacks == nil {
			j.callbacks = make(CallbackMap)
		}
		j.callbacks[c.MethodName] = &c
	}
	if count > 0 {
		return true
	}
	return false

}

/**
 *handle json rpc request
 */

func (j *JsonRpcService) ServerHandleRequest(json JsonRpcIf) {
	defer json.Destroy()
	req, err := json.ReadJsonRpcRequestHeaders()

	if err != nil || len(req) < 1 {
		json.WriteJsonRpcResponse(json.CreateExceptionResponse(req[0].Id, -32700, err))
		return
	}

	callback,err := j.CheckRpcRequestHeaders(req[0])

	if err != nil {
		json.WriteJsonRpcResponse(json.CreateExceptionResponse(req[0].Id, -32600, err))
		return
	}

	args := j.ParseRpcRequestArgument(req[0].Params)

	resp, err := j.CallMethod(callback,args)

	if err != nil {
		json.WriteJsonRpcResponse(json.CreateExceptionResponse(req[0].Id, -32601, err))
		return
	}

	json.WriteJsonRpcResponse(json.CreateSuccessResponse(req[0].Id,resp))
}

/**
 *check request headers
 */
func (j *JsonRpcService) CheckRpcRequestHeaders(req JsonRpcRequest) (*Callback, error) {

	status, service, method := j.parseServiceAndMethod(req.Method)
	var c *Callback
	var ok bool
	if !status {
		return c, errors.New("service name and method exception")
	}

	if _, ok = j.services[service]; !ok {
		return c, errors.New("service is not exist")
	}
	key := service + `.` + method
	if c, ok = j.callbacks[key]; !ok {
		return c, errors.New("method is not exist")
	}

	return c, nil

}

/**
 *parse rpc request argument
 */

func (j *JsonRpcService) ParseRpcRequestArgument(params []interface{}) ([]reflect.Value) {

	size := len(params)
	if size < 1 {
		return nil
	}

	argument := make([]reflect.Value, size)

	for key, val := range params {
		argument[key] = reflect.ValueOf(val)
	}
	return argument

}

/**
 *call method
 */
func (j *JsonRpcService) CallMethod(method *Callback, args []reflect.Value) (interface{}, error) {

	returnVal := method.Method.Func.Call(args)

	if len(returnVal) == 0 {
		return nil, errors.New(fmt.Sprintf("call callback failed : %s", method.MethodName))

	}

	return returnVal[0].Interface(), nil
}

/**
 *delimit service and method
 */
func (j *JsonRpcService) parseServiceAndMethod(name string) (bool, string, string) {
	list := strings.Split(name, SERVICEANDMETHODSEPARATOR)

	if len(list) < 2 {
		return false, name, list[0]
	}

	server := FormatName(list[0])

	if server == "" {

		return false, "", ""

	}
	method := FormatName(list[1])

	if method == "" {
		return false, "", ""
	}

	if len(list) == 2 {
		return true, list[0], list[1]
	}

	return false, name, list[0]

}

/**
 *
 */
func (j *JsonRpcService) judgeServiceIsExist(serviceName string) (bool) {
	if _, ok := j.services[serviceName]; ok {
		return true
	}
	return false
}

func (j *JsonRpcService) judgeMethodIsExist(callbackKey string) (bool) {

	if _, ok := j.callbacks[callbackKey]; ok {
		return true
	}
	return false
}

/**
 * format string name handle
 */
func FormatName(name string) (string) {

	return strings.ToLower(strings.Trim(name, " "))

}
