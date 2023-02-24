package verimor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiUrl = "http://sms.verimor.com.tr/v2/send.json"

type API struct{}

type Request struct {
	Username      string     `json:"username,omitempty"`
	Password      string     `json:"password,omitempty"`
	Source        string     `json:"source_addr,omitempty"`
	SendAt        string     `json:"send_at,omitempty"`
	CustomId      string     `json:"custom_id,omitempty"`
	DataCoding    string     `json:"datacoding,omitempty"`
	Commercial    bool       `json:"is_commercial,omitempty"`
	RecipientType string     `json:"iys_recipient_type,omitempty"`
	Messages      []*Message `json:"messages,omitempty"`
}

type Message struct {
	Msg string `json:"msg,omitempty"`
	No  string `json:"dest,omitempty"`
	Id  string `json:"id,omitempty"`
}

func (api *API) Sms(request Request) bool {
	if len(request.Messages) > 50000 {
		log.Println("max 50000 messages")
		return false
	}
	postdata, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return false
	}
	res, err := http.Post(ApiUrl, "application/json; charset=utf-8", bytes.NewReader(postdata))
	if err != nil {
		log.Println(err)
		return false
	}
	defer res.Body.Close()
	x, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(x))
	return true
}
