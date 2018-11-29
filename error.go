package json_rpc

import "errors"

/**
 define error code
 */
const (
	JsonRpcVersion     = "2.0" // json rpc version verbose
	ParseErrorCode     = -32700
	InvalidRequestCode = -32600
	MethodNotFoundCode = -32601
	InvalidParamsCode  = -32602
	InterbalErrorCode  = -32603
	ServerErrorCode    = -32000
)

/**
 *define error message
 */
var (
	ParseErrorMessage     = errors.New("parse error")
	InvalidRequestMessage = errors.New("invalid request")
	MethodNotFoundMessage = errors.New("method not found")
	InvalidParamsMesssage = errors.New("invalid params")
	InternalErrorMessage  = errors.New("internal error")
	ServerErrorMesssage   = errors.New("server error")
)



