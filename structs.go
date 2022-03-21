package regolt

// User
type User struct {
	ID       string `json:"_id"`
	Username string `json:"username"`

	Avatar struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"avatar"`

	Relations []struct {
		Status string `json:"status"`
		ID     string `json:"_id"`
	} `json:"relations"`

	Badges int `json:"badges"`

	Status struct {
		Text     string `json:"text"`
		Presence string `json:"presence"`
	} `json:"status"`

	Relationship string `json:"relationship"`
	Online       bool   `json:"online"`
	Flags        int    `json:"flags"`

	Bot struct {
		Owner string `json:"owner"`
	} `json:"bot,omitempty"`
}

type Categories struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Channels []string `json:"channels"`
}

type SystemMessages struct {
	UserJoined string `json:"user_joined"`
	UserLeft   string `json:"user_left"`
	UserKicked string `json:"user_kicked"`
	UserBanned string `json:"user_banned"`
}

// Server
type Server struct {
	ID          string   `json:"_id"`
	Owner       string   `json:"owner"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Channels    []string `json:"channels"`

	Categories []*Categories `json:"categories"`

	SystemMessages *SystemMessages `json:"system_messages"`

	Roles []Role `json:"roles"`

	DefaultPermissions []int `json:"default_permissions"`

	Icon struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"icon"`

	Banner struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"banner"`

	Nsfw         bool `json:"nsfw"`
	Flags        int  `json:"flags"`
	Analytics    bool `json:"analytics"`
	Discoverable bool `json:"discoverable"`
}

type EditServer struct {
	Name           string          `json:"name,omitempty"`
	Description    string          `json:"description,omitempty"`
	Icon           string          `json:"icon,omitempty"`
	Banner         string          `json:"banner,omitempty"`
	Categories     []*Categories   `json:"categories,omitempty"`
	SystemMessages *SystemMessages `json:"system_messages,omitempty"`
	Nsfw           bool            `json:"nsfw,omitempty"`
	Remove         string          `json:"remove,omitempty"`
}

type EditMember struct {
	Nickname string   `json:"nickname,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Remove   string   `json:"remove,omitempty"`
}

// Channel
type Channel struct {
	ID          string `json:"_id"`
	Server      string `json:"server"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Icon struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"icon"`

	DefaultPermissions int `json:"default_permissions"`

	RolePermissions map[string]int `json:"role_permissions"`

	Nsfw        bool   `json:"nsfw"`
	ChannelType string `json:"channel_type"`
}

// Group
type Group struct {
	ID            string   `json:"_id"`
	ChannelType   string   `json:"channel_type"`
	Recipients    []string `json:"recipients"`
	Name          string   `json:"name"`
	Owner         string   `json:"owner"`
	Description   string   `json:"description"`
	LastMessageID string   `json:"last_message_id"`

	Icon struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"icon"`

	Permissions int  `json:"permissions"`
	Nsfw        bool `json:"nsfw"`
}

// DM
type DirectMessage []struct {
	ID            string   `json:"_id"`
	ChannelType   string   `json:"channel_type"`
	Active        bool     `json:"active"`
	Recipients    []string `json:"recipients"`
	LastMessageID string   `json:"last_message_id"`
}

type Role struct {
	Name        string `json:"name"`
	Permissions []int  `json:"permissions"`
	Colour      string `json:"colour"`
	Hoist       bool   `json:"hoist"`
	Rank        int    `json:"rank"`
	Remove      string `json:"remove,omitempty"`
}

type Member struct {
	ID struct {
		Server string `json:"server"`
		User   string `json:"user"`
	} `json:"_id"`

	Nickname string `json:"nickname"`

	Avatar struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"avatar"`

	Roles []string `json:"roles"`
}

type UserProfile struct {
	Content string `json:"content"`

	Background struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"background"`
}

type MutualFriendsAndServers struct {
	Users   []string `json:"users"`
	Servers []string `json:"servers"`
}

type Message struct {
	ID      string `json:"_id"`
	Nonce   string `json:"nonce"`
	Channel string `json:"channel"`
	Author  string `json:"author"`

	Content interface{} `json:"content"`

	Attachments []struct {
		ID       string `json:"_id"`
		Tag      string `json:"tag"`
		Size     int    `json:"size"`
		Filename string `json:"filename"`

		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`

		ContentType string `json:"content_type"`
	} `json:"attachments"`

	Edited struct {
		Date string `json:"$date"`
	} `json:"edited"`

	Embeds []*Embed `json:"embeds"`

	Mentions []string `json:"mentions"`
	Replies  []string `json:"replies"`

	Masquerade struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"masquerade"`
}

type Embed struct {
	Type string `json:"type"`
	URL  string `json:"url"`

	Special struct {
		Type string `json:"type"`
	} `json:"special"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Image struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		Size   string `json:"size"`
	} `json:"image"`

	Video struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"video"`

	SiteName string `json:"site_name"`
	IconURL  string `json:"icon_url"`
	Colour   string `json:"colour"`
}

type RestContent struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type RoleCreated struct {
	Id          string `json:"id"`
	Permissions []int  `json:"permissions"`
}

type MessageSend struct {
	Content     string   `json:"content"`
	Attachments []string `json:"attachments,omitempty"`

	Embeds []*struct {
		Type        string `json:"type"`
		IconURL     string `json:"icon_url"`
		URL         string `json:"url"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Media       string `json:"media"`
		Colour      string `json:"colour"`
	} `json:"embeds,omitempty"`

	Replies []*struct {
		ID      string `json:"id"`
		Mention bool   `json:"mention"`
	} `json:"replies,omitempty"`

	Masquerade *struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"masquerade,omitempty"`
}

type EditUser struct {
	Status struct {
		Text     string `json:"text,omitempty"`
		Presence string `json:"presence,omitempty"`
	} `json:"status,omitempty"`

	Profile struct {
		Content    string `json:"content,omitempty"`
		Background string `json:"background,omitempty"`
	} `json:"profile,omitempty"`

	Avatar string `json:"avatar,omitempty"`
	Remove string `json:"remove,omitempty"`
}

type EditChannel struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Nsfw        bool   `json:"nsfw,omitempty"`
	Remove      string `json:"remove,omitempty"`
}

type Ban struct {
	ID struct {
		Server string `json:"server"`
		User   string `json:"user"`
	} `json:"_id"`

	Reason string `json:"reason"`
}

// Events / Handlers

type Event struct {
	Type string `json:"type"`
}

type Pong struct {
	Type string `json:"type"`
	Data int    `json:"data"`
}
type Ready struct {
	Type     string     `json:"type"`
	Users    []*User    `json:"users"`
	Servers  []*Server  `json:"servers"`
	Channels []*Channel `json:"channels"`
}

type readyEventHandler struct {
	Handler func(*Session, *Ready)
	Type    string
}

type MessageCreate struct {
	Type    string `json:"type"`
	Message Message
}

type messageCreateEventHandler struct {
	Handler func(*Session, *MessageCreate)
	Type    string
}

type MessageUpdate struct {
	Type    string  `json:"type"`
	Id      string  `json:"id"`
	Message Message `json:"data"`
}

type messageUpdateEventHandler struct {
	Handler func(*Session, *MessageUpdate)
	Type    string
}

type MessageDelete struct {
	Type    string `json:"type"`
	Id      string `json:"id"`
	Channel string `json:"channel"`
}

type messageDeleteEventHandler struct {
	Handler func(*Session, *MessageDelete)
	Type    string
}

type ChannelCreate struct {
	Type    string  `json:"type"`
	Channel Channel `json:"channel"`
}

type channelCreateEventHandler struct {
	Handler func(*Session, *ChannelCreate)
	Type    string
}

type ChannelUpdate struct {
	Type    string    `json:"type"`
	Id      string    `json:"id"`
	Channel Channel   `json:"data"`
	Clear   ClearType `json:"clear,omitempty"`
}

type channelUpdateEventHandler struct {
	Handler func(*Session, *ChannelUpdate)
	Type    string
}

type ChannelDelete struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type channelDeleteEventHandler struct {
	Handler func(*Session, *ChannelDelete)
	Type    string
}

type ChannelGroupJoin struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type channelGroupJoinEventHandler struct {
	Handler func(*Session, *ChannelGroupJoin)
	Type    string
}

type ChannelGroupLeave struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type channelGroupLeaveEventHandler struct {
	Handler func(*Session, *ChannelGroupLeave)
	Type    string
}

type ChannelStartTyping struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type channelStartTypingEventHandler struct {
	Handler func(*Session, *ChannelStartTyping)
	Type    string
}

type ChannelStopTyping struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type channelStopTypingEventHandler struct {
	Handler func(*Session, *ChannelStopTyping)
	Type    string
}

type ChannelAck struct {
	Type      string `json:"type"`
	Id        string `json:"id"`
	User      string `json:"user"`
	MessageId string `json:"message_id"`
}

type channelAckEventHandler struct {
	Handler func(*Session, *ChannelAck)
	Type    string
}

type ServerUpdate struct {
	Type   string    `json:"type"`
	Id     string    `json:"id"`
	Server Server    `json:"data"`
	Clear  ClearType `json:"clear,omitempty"`
}

type serverUpdateEventHandler struct {
	Handler func(*Session, *ServerUpdate)
	Type    string
}

type ServerDelete struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type serverDeleteEventHandler struct {
	Handler func(*Session, *ServerDelete)
	Type    string
}

type ServerMemberUpdate struct {
	Type string `json:"type"`
	Ids  struct {
		Server string `json:"server"`
		User   string `json:"user"`
	} `json:"id"`

	Member Member    `json:"data"`
	Clear  ClearType `json:"clear,omitempty"`
}

type serverMemberUpdateEventHandler struct {
	Handler func(*Session, *ServerMemberUpdate)
	Type    string
}

type ServerMemberJoin struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type serverMemberJoinEventHandler struct {
	Handler func(*Session, *ServerMemberJoin)
	Type    string
}

type ServerMemberLeave struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	User string `json:"user"`
}

type serverMemberLeaveEventHandler struct {
	Handler func(*Session, *ServerMemberLeave)
	Type    string
}

type ServerRoleUpdate struct {
	Type   string    `json:"type"`
	Id     string    `json:"id"`
	RoleId string    `json:"role_id"`
	Role   Role      `json:"data"`
	Clear  ClearType `json:"clear,omitempty"`
}

type serverRoleUpdateEventHandler struct {
	Handler func(*Session, *ServerRoleUpdate)
	Type    string
}

type ServerRoleDelete struct {
	Type   string `json:"type"`
	Id     string `json:"id"`
	RoleId string `json:"role_id"`
}

type serverRoleDeleteEventHandler struct {
	Handler func(*Session, *ServerRoleDelete)
	Type    string
}

type UserUpdate struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type userUpdateEventHandler struct {
	Handler func(*Session, *UserUpdate)
	Type    string
}

type UserRelationship struct {
	Type   string             `json:"type"`
	Id     string             `json:"id"`
	User   string             `json:"user"`
	Status RelationshipStatus `json:"status"`
}

type userRelationshipEventHandler struct {
	Handler func(*Session, *UserRelationship)
	Type    string
}
