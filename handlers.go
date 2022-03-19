package regolt

type eventHandler struct {
	name string
	h    interface{}
}

func (s *Session) AddHandler(handler interface{}) {
	h := interfaceToHandler(handler)

	s.handlers[h.name] = append(s.handlers[h.name], h)
}

func interfaceToHandler(h interface{}) eventHandler {
	switch v := h.(type) {
	case func(*Session, *Ready):
		return eventHandler{TypeReady, readyEventHandler{v, TypeReady}}
	case func(*Session, *MessageCreate):
		return eventHandler{TypeMessageCreate, messageCreateEventHandler{v, TypeMessageCreate}}
	case func(*Session, *MessageUpdate):
		return eventHandler{TypeMessageUpdate, messageUpdateEventHandler{v, TypeMessageUpdate}}
	case func(*Session, *MessageDelete):
		return eventHandler{TypeMessageDelete, messageDeleteEventHandler{v, TypeMessageDelete}}
	case func(*Session, *ChannelCreate):
		return eventHandler{TypeChannelCreate, channelCreateEventHandler{v, TypeChannelCreate}}
	case func(*Session, *ChannelUpdate):
		return eventHandler{TypeChannelUpdate, channelUpdateEventHandler{v, TypeChannelUpdate}}
	case func(*Session, *ChannelDelete):
		return eventHandler{TypeChannelDelete, channelDeleteEventHandler{v, TypeChannelDelete}}
	case func(*Session, *ChannelGroupJoin):
		return eventHandler{TypeChannelGroupJoin, channelGroupJoinEventHandler{v, TypeChannelGroupJoin}}
	case func(*Session, *ChannelGroupLeave):
		return eventHandler{TypeChannelGroupLeave, channelGroupLeaveEventHandler{v, TypeChannelGroupLeave}}
	case func(*Session, *ChannelStartTyping):
		return eventHandler{TypeChannelStartTyping, channelStartTypingEventHandler{v, TypeChannelStartTyping}}
	case func(*Session, *ChannelStopTyping):
		return eventHandler{TypeChannelStopTyping, channelStopTypingEventHandler{v, TypeChannelStopTyping}}
	case func(*Session, *ChannelAck):
		return eventHandler{TypeChannelAck, channelAckEventHandler{v, TypeChannelAck}}
	case func(*Session, *ServerUpdate):
		return eventHandler{TypeServerUpdate, serverUpdateEventHandler{v, TypeServerUpdate}}
	case func(*Session, *ServerDelete):
		return eventHandler{TypeServerDelete, serverDeleteEventHandler{v, TypeServerDelete}}
	case func(*Session, *ServerMemberUpdate):
		return eventHandler{TypeServerMemberUpdate, serverMemberUpdateEventHandler{v, TypeServerMemberUpdate}}
	case func(*Session, *ServerMemberJoin):
		return eventHandler{TypeServerMemberJoin, serverMemberJoinEventHandler{v, TypeServerMemberJoin}}
	case func(*Session, *ServerMemberLeave):
		return eventHandler{TypeServerMemberLeave, serverMemberLeaveEventHandler{v, TypeServerMemberLeave}}
	case func(*Session, *ServerRoleUpdate):
		return eventHandler{TypeServerRoleUpdate, serverRoleUpdateEventHandler{v, TypeServerRoleUpdate}}
	case func(*Session, *ServerRoleDelete):
		return eventHandler{TypeServerRoleDelete, serverRoleDeleteEventHandler{v, TypeServerRoleDelete}}
	case func(*Session, *UserUpdate):
		return eventHandler{TypeUserUpdate, userUpdateEventHandler{v, TypeUserUpdate}}
	case func(*Session, *UserRelationship):
		return eventHandler{TypeUserRelationship, userRelationshipEventHandler{v, TypeUserRelationship}}
	}

	return eventHandler{}
}
