package json_rpc

import "reflect"



/**
 *save service
 */
type Service struct {
	 ServerName string
	 ServiceTyp reflect.Type
	 ServiceVal reflect.Value
}

/**
 *save service callback
 */
type Callback struct {
	ServiceName string
	MethodName  string
	Method      reflect.Method
	MethodTyp   reflect.Type
	MethodVal   reflect.Value
	ArgTypes    []reflect.Type
}

type ServiceMap map[string]*Service //service map

type MethodMap map[string]*Callback // service method map

type CallbackMap map[string]*Callback // service implement method map

/**
 *
 */
type JsonRpcIf interface {
	 ReadJsonRpcRequestHeaders()([]JsonRpcRequest,error)
	 WriteJsonRpcResponse(interface{})(error)
	 CreateSuccessResponse(reqId interface{}, data interface{}) (JsonRpcSuccessResponse)
	 CreateExceptionResponse(reqId interface{}, code int,err error)(JsonRpcExceptionResponse)
	 CreateDefaultExceptionResponse(reqId interface{}, code int, message string) (JsonRpcExceptionResponse)
	 Destroy()
}

/**
 *  error interface
 */
type Error interface {
	 Error()(string)
	 ErrorId()(int)
}




