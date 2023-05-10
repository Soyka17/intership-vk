package longpoll

import (
	"VK_BOT_RAW/internal/usecase/messages"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type LongPoll struct {
	token string
	v     string
	group string
	Obj   map[string]interface{}
	wait  string
}

func NewLongPoll(token string, v string, g string, mh *messages.MessageHandler, wait string) *LongPoll {
	obj := make(map[string]interface{})
	obj["message_new"] = mh.NewMessage

	return &LongPoll{token: token, v: v, group: g, Obj: obj, wait: wait}
}

func (l *LongPoll) GetLongPollServer() (GetLongPollServerResponse, error) {
	var res GetLongPollServerResponse

	reqUrl := "https://api.vk.com/method/groups.getLongPollServer"

	data := url.Values{
		"access_token": {l.token},
		"v":            {l.v},
		"group_id":     {l.group},
	}

	resp, err := http.PostForm(reqUrl, data)

	if err != nil {
		log.Println("|WARN| Can't make http lonpoll server request : ", err)
		return res, err
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&res)

	log.Print("|INFO| Longpoll server:", res)

	return res, nil
}

func (l *LongPoll) LongPollRequest(lps GetLongPollServerResponse, ts string) (*LongPollResponse, error) {

	reqUrl := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%s", lps.Response.Server, lps.Response.Key, ts, l.wait)

	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Println("|WARN| Can't make longpoll request:", err)
		return nil, err
	}

	var res LongPollResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("|WARN| Can't decode longpoll body:", err)
		return nil, err
	}

	json.Unmarshal(body, &res)

	if len(res.Updates) > 0 && res.Failed == 0 {
		for _, upd := range res.Updates {
			log.Print("|INFO| New event:", upd.Type)
			if val, ok := l.Obj[upd.Type]; ok {
				val.(func(interface{}))(upd.Object)
			}
		}
	}

	return &res, nil
}

func (l *LongPoll) Run() error {
	var err error
	var resp *LongPollResponse
	server, err := l.GetLongPollServer()
	if err != nil {
		return err
	}
	ts := server.Response.Ts

	for {
		resp, err = l.LongPollRequest(server, ts)
		for err != nil {
			log.Println("|WARN| Longpoll request error:", err)
			for err != nil {
				log.Println("|INFO| Trying get new longpoll server")
				server, err = l.GetLongPollServer()
			}
			log.Println("|INFO| New server received")
			resp, err = l.LongPollRequest(server, ts)
		}
		switch resp.Failed {
		case 2:
			log.Print("|INFO| The key has expired")
			server, err = l.GetLongPollServer()
			if err != nil {
				return err
			}
			log.Print("|INFO| New server received")
		case 3:
			log.Print("|INFO| Information lost")
			server, err = l.GetLongPollServer()
			if err != nil {
				return err
			}
			log.Print("|INFO| New server received")
		default:
			ts = resp.Ts
		}
	}
}
