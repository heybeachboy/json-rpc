package json_rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"errors"
)

/**
* json rpc request head
*/

type JsonRpcRequest struct {
	Id      interface{}   `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  interface{} `json:"params"`
}

/**
 *json rpc error information
 */
type JsonError struct {
	Code    int         `json:"Code"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
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

	return j.parseJsonPrcRequestHead(RequestBytes)

}

/**
 *
 */
func (j *JsonRpc) parseJsonPrcRequestHead(data json.RawMessage) ([]JsonRpcRequest, error) {
	var request JsonRpcRequest
	err := json.Unmarshal(data, &request)

	if err != nil {
		return nil, err
	}

	return []JsonRpcRequest{request}, nil

}

func (j *JsonRpc) ParseRequestArguments(argTypes []reflect.Type, params interface{}) ([]reflect.Value,error) {
	if args, ok := params.(json.RawMessage); !ok {
		return nil, errors.New("Invalid params supplied")
	} else {
		return j.parsePositionalArguments(args, argTypes)
	}
}

func (j *JsonRpc)parsePositionalArguments(rawArgs json.RawMessage, types []reflect.Type) ([]reflect.Value, error) {
	// Read beginning of the args array.


	dec := json.NewDecoder(bytes.NewReader(rawArgs))

	if tok, _ := dec.Token(); tok != json.Delim('[') {
		return nil, errors.New("non-array args")
	}
	// Read args.

	args := make([]reflect.Value, 0)
	for i := 0; dec.More(); i++ {
		if i >= len(types) {
			return nil, errors.New(fmt.Sprintf("too many arguments, want at most %d", len(types)))
		}
		argval := reflect.New(types[i])
		if err := dec.Decode(argval.Interface()); err != nil {
			return nil, errors.New(fmt.Sprintf("invalid argument %d: %v", i, err))
		}
		if argval.IsNil() && types[i].Kind() != reflect.Ptr {
			return nil, errors.New(fmt.Sprintf("missing value for required argument %d", i))
		}
		args = append(args, argval.Elem())
	}

	// Read end of args array.
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	// Set any missing args to nil.
	for i := len(args); i < len(types); i++ {
		if types[i].Kind() != reflect.Ptr {
			return nil, errors.New(fmt.Sprintf("missing value for required argument %d", i))
		}
		args = append(args, reflect.Zero(types[i]))
	}

	return args, nil
}

/**
 *replay
 */
func (j *JsonRpc) WriteJsonRpcResponse(resp interface{}) (error) {

	return j.JsonEncode(resp)
}

func (j *JsonRpc) CreateSuccessResponse(reqId interface{}, data interface{}) (JsonRpcSuccessResponse) {
	var resp JsonRpcSuccessResponse
	resp.JsonRpc = JsonRpcVersion
	resp.Id = reqId
	resp.Result = data
	return resp
}

/**
 *create error response
 */
func (j *JsonRpc) CreateExceptionResponse(reqId interface{}, code int, err error) (JsonRpcExceptionResponse) {
	var resp JsonRpcExceptionResponse
	resp.Id = reqId
	resp.JsonRpc = JsonRpcVersion
	resp.Error.Message = GetErrorMessage(code)
	resp.Error.Code = code
	resp.Error.Data = err.Error()
	return resp
}

/**
 *create default error response
 */

func (j *JsonRpc) CreateDefaultExceptionResponse(reqId interface{}, code int, message string) (JsonRpcExceptionResponse) {
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
