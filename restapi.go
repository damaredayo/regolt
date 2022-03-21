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

func makeChannelPermissionsURL(in, in2 string) string {
	return BASE_URL + "/channels" + in + "/permissions/" + in2
}

func makeServerURL(in string) string {
	return BASE_URL + "/servers" + in
}

func makeRolesURL(in, in2 string) string {
	return BASE_URL + "/servers" + in + "/roles/" + in2
}

func makeServerMemberURL(in, in2 string) string {
	return BASE_URL + "/servers" + in + "/members/" + in2
}

func makeServerMemberBanURL(in, in2 string) string {
	return BASE_URL + "/servers" + in + "/bans/" + in2
}

func makeInviteURL(in string) string {
	return BASE_URL + "/invites/" + in
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
	return fmt.Errorf("Unimplemented (Session Token)")
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

func (s *Session) FetchChannel(channelId string) (channel *Channel, err error) {
	resp, err := s.request(http.MethodGet, makeChannelsURL(channelId), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, channel)
	return
}

func (s *Session) EditChannel(channelId string, edit *EditChannel) (channel *Channel, err error) {
	resp, err := s.request(http.MethodPatch, makeChannelsURL(channelId), "application/json", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, channel)
	return
}

func (s *Session) CloseChannel(channelId string, edit *EditChannel) (err error) {
	_, err = s.request(http.MethodDelete, makeChannelsURL(channelId), "application/json", nil)
	return
}

// Invites

// Create invite from channel ID
func (s *Session) CreateInvite(channelId string) (invite string, err error) {
	i := struct {
		Code string `json:"code"`
	}{}

	resp, err := s.request(http.MethodPost, makeChannelsURL(channelId)+"/invites", "application/json", nil)

	err = json.Unmarshal(resp, &i)
	if err != nil {
		return
	}

	return i.Code, err
}

// Permissions

func (s *Session) SetRolePermission(channelId string, roleId string, permissions int) (err error) {
	i := struct {
		Permissions int `json:"permissions"`
	}{permissions}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodPut, makeChannelPermissionsURL(channelId, roleId), "application/json", body)
	return
}

func (s *Session) SetDefaultPermission(id string, permissions int) (err error) {
	return s.SetRolePermission(id, "default", permissions)
}

// Permissions (TODO)

// Messages
func (s *Session) ChannelMessageSend(channelId string, msg *MessageSend) (resp []byte, err error) {
	var body []byte
	if msg != nil {
		body, err = json.Marshal(msg)
		if err != nil {
			return
		}
		resp, err = s.request(http.MethodPost, makeSendMessageURL(channelId), "application/json", body)
	}
	return
}

func (s *Session) ChannelMessageSendString(channelId string, msg string) (resp []byte, err error) {
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
		resp, err = s.request(http.MethodPost, makeSendMessageURL(channelId), "application/json", body)
	}
	return
}

func (s *Session) FetchMessages(channelId string, limit int, before string, after string, sort string, nearby string, includeUsers bool) (
	messages []*Message, users []*User, members []*Member, err error) {

	var limitPtr *int

	if sort == "" {
		return
	}

	if limit > 100 {
		limit = 100
	}

	limitPtr = &limit

	if limit < 1 {
		limitPtr = nil
	}

	i := struct {
		Limit        *int   `json:"limit,omitempty"`
		Before       string `json:"before,omitempty"`
		After        string `json:"after,omitempty"`
		Sort         string `json:"sort"`
		Nearby       string `json:"nearby,omitempty"`
		IncludeUsers *bool  `json:"include_users,omitempty"`
	}{
		limitPtr, before, after, sort, nearby, &includeUsers,
	}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	resp, err := s.request(http.MethodGet, makeChannelsURL(channelId)+"/messages", "application/json", body)

	j := struct {
		Messages []*Message
		Users    []*User
		Members  []*Member
	}{}

	err = json.Unmarshal(resp, &j)

	if err != nil {
		return
	}

	return j.Messages, j.Users, j.Members, err
}

func (s *Session) FetchMessage(channelId, messageId string) (message *Message, err error) {

	resp, err := s.request(http.MethodGet, makeChannelsURL(channelId)+"/messages"+messageId, "application/json", nil)

	err = json.Unmarshal(resp, message)
	return
}

func (s *Session) EditMessage(channelId, messageId, content string, embeds []*Embed) (err error) {
	var body []byte

	i := struct {
		Content string   `json:"content"`
		Embeds  []*Embed `json:"embeds,omitempty"`
	}{
		content, embeds,
	}

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodGet, makeChannelsURL(channelId)+"/messages"+messageId, "application/json", body)

	return
}

func (s *Session) DeleteMessage(channelId, messageId string) (err error) {
	_, err = s.request(http.MethodDelete, makeChannelsURL(channelId)+"/messages"+messageId, "application/json", nil)
	return
}

func (s *Session) PollMessageChanges(channelId string, messageIds []string) (changed []*Message, deleted []string, err error) {
	var body []byte

	i := struct {
		Ids []string `json:"ids"`
	}{messageIds}

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	resp, err := s.request(http.MethodPost, makeChannelsURL(channelId)+"/messages/stale", "application/json", body)
	if err != nil {
		return
	}

	j := struct {
		Changed []*Message `json:"changed"`
		Deleted []string   `json:"deleted"`
	}{}

	err = json.Unmarshal(resp, &j)

	return j.Changed, j.Deleted, err
}

func (s *Session) SearchMessages(channelId string, query string, limit int, before string, after string, sort string, nearby string, includeUsers *bool) (
	messages []*Message, users []*User, members []*Member, err error) {

	var limitPtr *int

	if sort == "" {
		return
	}

	if limit > 100 {
		limit = 100
	}

	limitPtr = &limit

	if limit < 1 {
		limitPtr = nil
	}

	i := struct {
		Query        string `json:"query"`
		Limit        *int   `json:"limit,omitempty"`
		Before       string `json:"before,omitempty"`
		After        string `json:"after,omitempty"`
		Sort         string `json:"sort,omitempty"`
		Nearby       string `json:"nearby,omitempty"`
		IncludeUsers *bool  `json:"include_users,omitempty"`
	}{
		query, limitPtr, before, after, sort, nearby, includeUsers,
	}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	resp, err := s.request(http.MethodPost, makeChannelsURL(channelId)+"/messages/search", "application/json", body)

	j := struct {
		Messages []*Message
		Users    []*User
		Members  []*Member
	}{}

	err = json.Unmarshal(resp, &j)

	if err != nil {
		return
	}

	return j.Messages, j.Users, j.Members, err

}

// Unimplemented (Session Token)
func (s *Session) MessageAck(channelId, messageId string) (err error) {
	return fmt.Errorf("Unimplemented")
}

// Groups

// Unimplemented (Session Token)
func (s *Session) CreateGroup(name, description string, users []string, nsfw bool) (group *Group, err error) {
	err = fmt.Errorf("Unimplemented")
	return
}

func (s *Session) FetchGroupMembers(channelId, messageId string) (members []*Member, err error) {

	resp, err := s.request(http.MethodGet, makeChannelsURL(channelId)+"/members", "application/json", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &members)

	return
}

// Unimplemented (Session Token) + documentation error
func (s *Session) AddGroupMember(channelId, userId string) (err error) {
	return fmt.Errorf("Unimplemented")
}

// Unimplemented (Session Token) + documentation error
func (s *Session) RemoveGroupMember(channelId, userId string) (err error) {
	return fmt.Errorf("Unimplemented")
}

// Voice

func (s *Session) JoinVoiceChannel(channelId string) (token string, err error) {
	resp, err := s.request(http.MethodGet, makeChannelsURL(channelId)+"/join_call", "application/json", nil)
	if err != nil {
		return
	}

	i := struct {
		Token string `json:"token"`
	}{}

	err = json.Unmarshal(resp, &i)
	return
}

// Servers

func (s *Session) FetchServer(serverId string) (server *Server, err error) {
	resp, err := s.request(http.MethodGet, makeServerURL(serverId), "application/json", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &server)

	return
}

func (s *Session) EditServer(serverId string, edit *EditServer) (err error) {
	var body []byte

	body, err = json.Marshal(edit)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodPatch, makeServerURL(serverId), "application/json", body)

	return
}

func (s *Session) DeleteServer(serverId string) (err error) {
	_, err = s.request(http.MethodDelete, makeServerURL(serverId), "application/json", nil)
	return
}

// Unimplemented (Session Token)
func (s *Session) CreateServer(name, description string, nsfw bool) (server *Server, err error) {
	err = fmt.Errorf("Unimplemented")
	return
}

func (s *Session) CreateServerChannel(serverId, name, description string, channelType ChannelType, nsfw bool) (err error) {
	i := struct {
		Type        ChannelType `json:"type"`
		Name        string      `json:"name"`
		Description string      `json:"description,omitempty"`
		Nsfw        bool        `json:"nsfw,omitempty"`
	}{
		channelType, name, description, nsfw,
	}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodPatch, makeServerURL(serverId)+"/channels", "application/json", body)
	return

}

func (s *Session) FetchInvites(serverId string) (invites []string, err error) {
	resp, err := s.request(http.MethodGet, makeServerURL(serverId)+"/invites", "application/json", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &invites)
	return
}

// Unimplemented (Session Token)
func (s *Session) MarkServerAsRead() (err error) {
	return fmt.Errorf("Unimplemented")
}

// Server Members
func (s *Session) FetchMember(serverId, memberId string) (member *Member, err error) {
	resp, err := s.request(http.MethodGet, makeServerMemberURL(serverId, memberId), "application/json", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &member)
	return
}

func (s *Session) EditMember(serverId, memberId string, edit *EditMember) (err error) {
	var body []byte

	body, err = json.Marshal(edit)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodPatch, makeServerMemberURL(serverId, memberId), "application/json", body)

	return
}

func (s *Session) KickMember(serverId, memberId string) (err error) {
	_, err = s.request(http.MethodDelete, makeServerMemberURL(serverId, memberId), "application/json", nil)
	return
}

func (s *Session) FetchMembers(serverId, memberId string) (members []*Member, users []*User, err error) {
	resp, err := s.request(http.MethodGet, makeServerMemberURL(serverId, memberId), "application/json", nil)
	if err != nil {
		return
	}

	i := struct {
		Members []*Member `json:"members"`
		Users   []*User   `json:"users"`
	}{}

	err = json.Unmarshal(resp, &i)
	return i.Members, i.Users, err
}

// Bans
func (s *Session) BanMember(serverId, memberId, reason string) (err error) {
	_, err = s.request(http.MethodPut, makeServerMemberBanURL(serverId, memberId), "application/json", nil)
	return
}

func (s *Session) UnbanMember(serverId, memberId, reason string) (err error) {
	_, err = s.request(http.MethodDelete, makeServerMemberBanURL(serverId, memberId), "application/json", nil)
	return
}

func (s *Session) FetchBans(serverId string) (users []*User, bans []*Ban, err error) {
	resp, err := s.request(http.MethodDelete, makeServerURL(serverId)+"/bans", "application/json", nil)

	i := struct {
		Users []*User `json:"users"`
		Bans  []*Ban  `json:"bans"`
	}{}

	err = json.Unmarshal(resp, &i)
	return i.Users, i.Bans, err
}

// Roles

func (s *Session) CreateRole(serverId, name string) (role *RoleCreated, err error) {
	if len(name) > 32 {
		return nil, fmt.Errorf("name greater than 32")
	}

	i := struct {
		Name string `json:"name"`
	}{name}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	resp, err := s.request(http.MethodPost, makeServerURL(serverId)+"/roles", "application/json", body)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, role)
	return
}

func (s *Session) EditRole(serverId, roleId, name, colour string, hoist bool, rank int, remove RoleClearType) (err error) {
	i := struct {
		Name   string        `json:"name"`
		Colour string        `json:"colour"`
		Hoist  bool          `json:"hoist"`
		Rank   int           `json:"rank"`
		Remove RoleClearType `json:"remove"`
	}{
		name, colour, hoist, rank, remove,
	}

	var body []byte

	body, err = json.Marshal(i)
	if err != nil {
		return
	}

	_, err = s.request(http.MethodPatch, makeRolesURL(serverId, roleId), "application/json", body)
	return
}

func (s *Session) DeleteRole(serverId, roleId string) (err error) {
	_, err = s.request(http.MethodDelete, makeRolesURL(serverId, roleId), "application/json", nil)
	return
}

// Invites

func (s *Session) FetchInvite(invite string) (server *Server, err error) {
	resp, err := s.request(http.MethodGet, makeInviteURL(invite), "application/json", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, server)
	return
}

// Unimplemented (Session Token)
func (s *Session) JoinInvite(invite string) (server *Server, err error) {
	err = fmt.Errorf("Unimplemented")
	return
}

func (s *Session) DeleteInvite(invite string) (err error) {
	_, err = s.request(http.MethodDelete, makeInviteURL(invite), "application/json", nil)
	return
}
