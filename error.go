package json_rpc

/**
 define error code
 */
const (
	JsonRpcVersion     = "2.0" // json rpc version verbose
	ParseErrorCode     = -32700
	InvalidRequestCode = -32600
	MethodNotFoundCode = -32601
	InvalidParamsCode  = -32602
	InternalErrorCode  = -32603
	ServerErrorCode    = -32000
)


/**
 * get json rpc error message
 */
func GetErrorMessage(errorCode int) (string) {
	switch errorCode {
	case ParseErrorCode:
		return "Parse error"
	case InvalidRequestCode:
		return "Invalid request"
	case MethodNotFoundCode:
		return "Method not found"
	case InvalidParamsCode:
		return "Invalid params"
	case InternalErrorCode:
		return "Internal error"
	case ServerErrorCode:
		return "Service"
	default:
		return "Unknown error"
	}

}




