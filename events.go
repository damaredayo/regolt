package regolt

import (
	"encoding/json"
	"fmt"
	"log"
)

// Client -> Server

// Authenticate with Revolt

func (s *Session) authenticate() error {
	auth := struct {
		Type  string `json:"type"`
		Token string `json:"token"`
	}{
		Type:  "Authenticate",
		Token: s.Token,
	}

	s.wsMutex.Lock()
	defer s.wsMutex.Unlock()

	return s.ws.WriteJSON(auth)
}

// Server -> Client events

var (
	TypeAuthenticated      = "Authenticated"
	TypePong               = "Pong"
	TypeReady              = "Ready"
	TypeMessageCreate      = "Message"
	TypeMessageUpdate      = "MessageUpdate"
	TypeMessageDelete      = "MessageDelete"
	TypeChannelCreate      = "ChannelCreate"
	TypeChannelUpdate      = "ChannelUpdate"
	TypeChannelDelete      = "ChannelDelete"
	TypeChannelGroupJoin   = "ChannelGroupJoin"
	TypeChannelGroupLeave  = "ChannelGroupLeave"
	TypeChannelStartTyping = "ChannelStartTyping"
	TypeChannelStopTyping  = "ChannelStopTyping"
	TypeChannelAck         = "ChannelAck"
	TypeServerUpdate       = "ServerUpdate"
	TypeServerDelete       = "ServerDelete"
	TypeServerMemberUpdate = "ServerMemberUpdate"
	TypeServerMemberJoin   = "ServerMemberJoin"
	TypeServerMemberLeave  = "ServerMemberLeave"
	TypeServerRoleUpdate   = "ServerRoleUpdate"
	TypeServerRoleDelete   = "ServerRoleDelete"
	TypeUserUpdate         = "UserUpdate"
	TypeUserRelationship   = "UserRelationship"
)

func (s *Session) handler(message []byte) {
	fmt.Println("debug:", string(message))

	var e Event
	if err := json.Unmarshal(message, &e); err != nil {
		return
	}

	switch e.Type {
	case TypeAuthenticated:
	case TypePong:
	// Ready
	case TypeReady:
		var ready Ready

		if err := json.Unmarshal(message, &ready); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeReady]

		for _, handler := range handlers {
			handler.h.(readyEventHandler).Handler(s, &ready)
		}

	// Message Create
	case TypeMessageCreate:
		var msg Message

		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeMessageCreate]

		for _, handler := range handlers {
			handler.h.(messageCreateEventHandler).Handler(s, &MessageCreate{TypeMessageCreate, msg})
		}

	// Message Update
	case TypeMessageUpdate:
		var msg MessageUpdate

		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeMessageUpdate]

		for _, handler := range handlers {
			handler.h.(messageUpdateEventHandler).Handler(s, &msg)
		}

	// Message Delete
	case TypeMessageDelete:
		var msg MessageDelete

		if err := json.Unmarshal(message, &msg); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeMessageDelete]

		for _, handler := range handlers {
			handler.h.(messageDeleteEventHandler).Handler(s, &msg)
		}

	// Channel Create
	case TypeChannelCreate:
		var channel ChannelCreate

		if err := json.Unmarshal(message, &channel); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelCreate]
		for _, handler := range handlers {
			handler.h.(channelCreateEventHandler).Handler(s, &channel)
		}

	// Channel Update
	case TypeChannelUpdate:
		var channel ChannelUpdate

		if err := json.Unmarshal(message, &channel); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelUpdate]
		for _, handler := range handlers {
			handler.h.(channelUpdateEventHandler).Handler(s, &channel)
		}

	// Channel Delete
	case TypeChannelDelete:
		var channel ChannelDelete

		if err := json.Unmarshal(message, &channel); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelDelete]
		for _, handler := range handlers {
			handler.h.(channelDeleteEventHandler).Handler(s, &channel)
		}

	// Group Join
	case TypeChannelGroupJoin:
		var channel ChannelGroupJoin

		if err := json.Unmarshal(message, &channel); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelGroupJoin]
		for _, handler := range handlers {
			handler.h.(channelGroupJoinEventHandler).Handler(s, &channel)
		}

	// Group Leave
	case TypeChannelGroupLeave:
		var channel ChannelGroupLeave

		if err := json.Unmarshal(message, &channel); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelGroupLeave]
		for _, handler := range handlers {
			handler.h.(channelGroupLeaveEventHandler).Handler(s, &channel)
		}

	// Start Typing
	case TypeChannelStartTyping:
		var typing ChannelStartTyping

		if err := json.Unmarshal(message, &typing); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelStartTyping]
		for _, handler := range handlers {
			handler.h.(channelStartTypingEventHandler).Handler(s, &typing)
		}

	// Stop Typing
	case TypeChannelStopTyping:
		var typing ChannelStopTyping

		if err := json.Unmarshal(message, &typing); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelStopTyping]
		for _, handler := range handlers {
			handler.h.(channelStopTypingEventHandler).Handler(s, &typing)
		}

	// Channel ack
	case TypeChannelAck:
		var ack ChannelAck

		if err := json.Unmarshal(message, &ack); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeChannelAck]
		for _, handler := range handlers {
			handler.h.(channelAckEventHandler).Handler(s, &ack)
		}

	// Server update
	case TypeServerUpdate:
		var server ServerUpdate

		if err := json.Unmarshal(message, &server); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerUpdate]
		for _, handler := range handlers {
			handler.h.(serverUpdateEventHandler).Handler(s, &server)
		}

	// Server delete
	case TypeServerDelete:
		var server ServerDelete

		if err := json.Unmarshal(message, &server); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerDelete]
		for _, handler := range handlers {
			handler.h.(serverDeleteEventHandler).Handler(s, &server)
		}

	// Server member update
	case TypeServerMemberUpdate:
		var member ServerMemberUpdate

		if err := json.Unmarshal(message, &member); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerMemberUpdate]
		for _, handler := range handlers {
			handler.h.(serverMemberUpdateEventHandler).Handler(s, &member)
		}

	// Server member join
	case TypeServerMemberJoin:
		var member ServerMemberJoin

		if err := json.Unmarshal(message, &member); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerMemberJoin]
		for _, handler := range handlers {
			handler.h.(serverMemberJoinEventHandler).Handler(s, &member)
		}

	// Server member leave
	case TypeServerMemberLeave:
		var member ServerMemberLeave

		if err := json.Unmarshal(message, &member); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerMemberLeave]
		for _, handler := range handlers {
			handler.h.(serverMemberLeaveEventHandler).Handler(s, &member)
		}

	// User update
	case TypeUserUpdate:
		var member ServerMemberLeave

		if err := json.Unmarshal(message, &member); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeServerMemberLeave]
		for _, handler := range handlers {
			handler.h.(serverMemberLeaveEventHandler).Handler(s, &member)
		}

	// Relationship update
	case TypeUserRelationship:
		var relationship UserRelationship

		if err := json.Unmarshal(message, &relationship); err != nil {
			log.Println(err)
			return
		}

		handlers := s.handlers[TypeUserRelationship]
		for _, handler := range handlers {
			handler.h.(userRelationshipEventHandler).Handler(s, &relationship)
		}

	default:
	}
}
