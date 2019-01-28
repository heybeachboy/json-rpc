package json_rpc

import (
	"encoding/json"
	"io"
)

/**
* json rpc request head
*/

type JsonRpcRequest struct {
	Id      interface{} `json:"id"`
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

/**
 *json rpc error information
 */
type JsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

/**
 *json rpc successful of response
 */
type JsonRpcSuccessResponse struct {
	Id      interface{}
	JsonRpc string
	Result  interface{}
}

/**
 *json rpc error of response
 */
type JsonRpcExceptionResponse struct {
	Id      interface{}
	JsonRpc string
	Error   JsonError
}



/**
 * parse json string
 */

type JsonRpc struct {
	JsonDecode func(v interface{}) error
	JsonEncode func(v interface{}) error
	Rw         io.ReadWriteCloser
}

/**
 *init json-rpc service
 */

func NewJsonRpc(io *httpReaderWriterAndCloser) (*JsonRpc) {
	return &JsonRpc{json.NewDecoder(io).Decode, json.NewEncoder(io).Encode, io}

}

/**
 *
 */

func (j *JsonRpc) ReadJsonRpcRequestHeaders() ([]JsonRpcRequest, error) {
	var RequestBytes json.RawMessage
	err := j.JsonDecode(&RequestBytes)

	if err != nil {
		return nil, err
	}

	return  j.parseJsonPrcRequestHead(RequestBytes)

}

/**
 *
 */
func (j *JsonRpc) parseJsonPrcRequestHead(data json.RawMessage)([]JsonRpcRequest,error) {
	 var request []JsonRpcRequest
	 err := json.Unmarshal(data,&request[0])

	 if err != nil {
	 	return nil, err
	 }

	 return request, nil

}

/**
 *replay
 */
func (j *JsonRpc) WriteJsonRpcResponse(resp interface{})(error) {

	return j.JsonEncode(resp)
}

/**
 *create error response
 */
func (j *JsonRpc) CreateExceptionResponse(reqId interface{}, code int)(JsonRpcExceptionResponse) {
	var resp JsonRpcExceptionResponse
	resp.Id = reqId
	resp.JsonRpc = JsonRpcVersion
	resp.Error.Message = GetErrorMessage(code)
	return resp
}

/**
 *create default error response
 */

func (j *JsonRpc)CreateDefaultExceptionResponse(reqId interface{}, code int, message string) (JsonRpcExceptionResponse) {
	var resp JsonRpcExceptionResponse
	resp.Id = reqId
	resp.JsonRpc = JsonRpcVersion
	resp.Error.Code = code
	resp.Error.Message = message
	return resp

}

func (j *JsonRpc) Destroy() {
	j.Rw.Close()
}
