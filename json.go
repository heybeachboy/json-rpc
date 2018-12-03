package json_rpc

/**
 * parse json string
 */

/**
* json rpc request head
*/

type JsonRpcRequest struct {
	Id      interface{} `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JsonRpcSuccessResponse struct {
	Id      interface{}
	JsonRpc string
	Result  interface{}
}

type JsonRpcExceptionResponse struct {
	Id      interface{}
	JsonRpc string
	Error   JsonError
}

type JsonRpcRequestDecodeIf interface {
}

type JsonRpcResponseEncodeIf interface {

}
