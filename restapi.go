package regolt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL = "http://api.revolt.chat"

var ErrNoToken = fmt.Errorf("no token found")

func makeSendMessageURL(id string) string {
	return BASE_URL + "/channels/" + id + "/messages"
}

func (s *Session) request(method, url, contentType string, json []byte) (resp []byte, err error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(json))
	if err != nil {
		return
	}

	if s.Token == "" {
		err = ErrNoToken
		return
	}

	req.Header.Set("x-bot-token", s.Token)

	req.Header.Set("User-Agent", "Bot (https://github.com/damaredayo/regolt)")

	r, err := s.http.Do(req)
	if err != nil {
		return
	}

	resp, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	switch r.StatusCode {
	case http.StatusOK:
	case http.StatusNoContent:
	case http.StatusBadGateway:
		// TODO: retry logic (recursion?)
		err = fmt.Errorf("bad gateway")
		return
	}

	return
}

func (s *Session) ChannelMessageSend(id string, msg *MessageSend) (resp []byte, err error) {
	var body []byte
	if msg != nil {
		body, err = json.Marshal(msg)
		if err != nil {
			return
		}
		resp, err = s.request(http.MethodPost, makeSendMessageURL(id), "application/json", body)
	}
	return
}

func (s *Session) ChannelMessageSendString(id string, msg string) (resp []byte, err error) {
	var body []byte

	if msg != "" {

		msgStruct := &MessageSend{
			Content: msg,
		}

		body, err = json.Marshal(msgStruct)
		if err != nil {
			return
		}
		fmt.Println(string(body))
		resp, err = s.request(http.MethodPost, makeSendMessageURL(id), "application/json", body)
	}

	return
}
