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

func makeUsersURL(in string) string {
	return BASE_URL + "/users/" + in
}

func makeChannelsURL(in string) string {
	return BASE_URL + "/channels" + in
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

// Users

// User Information

func (s *Session) FetchUser(id string) (user *User, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL(id), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, user)
	return
}

func (s *Session) EditUser(edit *EditUser) (err error) {
	var body []byte
	if edit != nil {
		body, err = json.Marshal(edit)
		if err != nil {
			return
		}
		_, err = s.request(http.MethodPatch, makeUsersURL("@me"), "application/json", body)
	}
	return
}

func (s *Session) FetchSelf() (user *User, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL("@me"), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, user)
	return
}

// Unimplemented
func (s *Session) ChangeUsername() (err error) {
	return fmt.Errorf("Unimplemented")
}

func (s *Session) FetchUserProfile(id string) (profile *UserProfile, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL("@me")+"/profile", "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, profile)
	return
}

func (s *Session) FetchDefaultAvatar(id string) (img []byte, err error) {
	return s.request(http.MethodGet, makeUsersURL(id)+"/default_avatar", "application/json", nil)
}

func (s *Session) FetchMutualFriendsAndServers(id string) (mutuals *MutualFriendsAndServers, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL(id)+"/mutual", "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, mutuals)
	return
}

// Alias for FetchMutualFriendsAndServers
func (s *Session) FetchMutuals(id string) (mutuals *MutualFriendsAndServers, err error) {
	return s.FetchMutualFriendsAndServers(id)
}

// DMs

func (s *Session) FetchDMs() (dms []*DirectMessage, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL("dms"), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &dms)
	return
}

func (s *Session) OpenDM(id string) (dm *DirectMessage, err error) {
	resp, err := s.request(http.MethodGet, makeUsersURL(id)+"/dm", "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, dm)
	return
}

// Channels

func (s *Session) FetchChannel(id string) (channel *Channel, err error) {
	resp, err := s.request(http.MethodGet, makeChannelsURL(id), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, channel)
	return
}

func (s *Session) EditChannel(id string, edit *EditChannel) (channel *Channel, err error) {
	resp, err := s.request(http.MethodPatch, makeChannelsURL(id), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, channel)
	return
}

func (s *Session) CloseChannel(id string, edit *EditChannel) (err error) {
	_, err = s.request(http.MethodDelete, makeChannelsURL(id), "application/json", nil)
	return
}

// Invites

func (s *Session) CreateInvite(id string) (invite string, err error) {
	i := struct {
		Code string `json:"code"`
	}{}

	resp, err := s.request(http.MethodPost, makeChannelsURL(id)+"/invites", "application/json", nil)

	err = json.Unmarshal(resp, &i)
	if err != nil {
		return
	}

	return i.Code, err
}

// Permissions (TODO)

// Messages
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
