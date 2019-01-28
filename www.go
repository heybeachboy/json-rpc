package json_rpc

import (
	"encoding/json"
	"fmt"
	"github.com/name5566/leaf/log"
	"io"
	"net/http"

)

const (
	USERNAME        = `admin`
	PASSWORD        = `123`
	MAXCONTENTLIMIT = 1024
	CONTENTTYPE     = `application/json`
)

type Replay struct {
	 ProtocolVersion string
	 Name string
	 Password string
	 Age int
}

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.URL.Host)
	fmt.Println("User:", r.URL.User)
	fmt.Println("Query:", r.URL.RawQuery)
	fmt.Println("Scheme:", r.URL.Scheme)
	fmt.Println("Content-Type:",r.Header.Get("Content-Type"))

	user, password, ok := r.BasicAuth()

	if !ok {
		w.WriteHeader(http.StatusForbidden)
		return
	} else {
		if user != USERNAME || password != PASSWORD {
			w.WriteHeader(http.StatusForbidden)
			return
		}

	}

	fmt.Println("user:", user)
	fmt.Println("password:", password)
	body := io.LimitReader(r.Body, MAXCONTENTLIMIT)

	decodeJson := json.NewDecoder(body)
	var data Replay
	decodeJson.Decode(&data)

	encodeJson := json.NewEncoder(w)

	w.Header().Set("Content-Type", CONTENTTYPE)
	encodeJson.Encode(data)
	//fmt.Fprintf(w, content)


}

func main() {
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":7080", nil); err != nil {
		log.Fatal("http server exception : %s", err.Error())
	}

}
