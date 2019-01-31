package json_rpc

import (
	"fmt"
	"io"
	"mime"
	"net/http"
)

/**
 *@auth Mr.Zhou leon
 *@mail zhouletian1234@live.com
 *@date 2018-12-15
 */
const (
	MAX_REQUEST_LIMIT = 1024
	CONTENT_TYPE      = `application/json`
)

type httpReaderWriterAndCloser struct {
	 io.Reader
	 io.Writer
	 io.Closer
}

/**
 *http
 */
func (j *JsonRpcService) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	code, msg := j.CheckRequest(r)

	if code != 0 {
		http.Error(w, msg, code)
		return
	}

	content := io.LimitReader(r.Body,MAX_REQUEST_LIMIT)

    w.Header().Set("Content-Type",CONTENT_TYPE)
	jsonParse := NewJsonRpc(&httpReaderWriterAndCloser{content,w,r.Body})

	j.ServerHandleRequest(jsonParse)


}



/**
 *check request
 */
func (j *JsonRpcService) CheckRequest(r *http.Request) (int, string) {

	if r.Method != http.MethodPost {
		return http.StatusMethodNotAllowed, "method not allowed"

	}

	if r.ContentLength > MAX_REQUEST_LIMIT {
		return http.StatusRequestEntityTooLarge, fmt.Sprintf("content length too large (%d>%d)", r.ContentLength, MAX_REQUEST_LIMIT)
	}

	mt, _, err := mime.ParseMediaType(r.Header.Get("content-type"))
	if r.Method != http.MethodOptions && (err != nil || mt != CONTENT_TYPE) {
		err := fmt.Errorf("invalid content type, only %s is supported", CONTENT_TYPE)
		return http.StatusUnsupportedMediaType, err.Error()
	}

	return 0, ""

}
