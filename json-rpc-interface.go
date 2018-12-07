package json_rpc

import "reflect"

const SERVICEANDMETHODSEPARATOR = `_`

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
	MethodName string
	Method     reflect.Method
	MethodTyp  reflect.Type
	MethodVal  reflect.Value
	ArgTypes   []reflect.Type
}

type ServiceMap map[string]*Service //service map

type MethodMap map[string]*Callback // service method map

type CallbackMap map[string]*Callback // service implement method map

/**
 *  error interface
 */
type Error interface {
	 Error()(string)
	 ErrorId()(int)
}

/**
 *json format handle struct
 */
type JsonRpcIf interface {
	 ReadRequestHead()(JsonRpcRequest)
	 CreateResponse(id interface{}, reply interface{}) interface{}
	 CreateErrorResponse(id interface{}, err Error) interface{}
	 CreateErrorResponseWithInfo(id interface{}, err Error, info interface{}) interface{}
	 CreateNotification(id, namespace string, event interface{}) interface{}
	 Write(msg interface{}) error
	 Close()
}


