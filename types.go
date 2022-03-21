package regolt

type ClearType string

const ClearTypeIcon ClearType = "Icon"
const ClearTypeDescription ClearType = "Description"
const ClearTypeBanner ClearType = "Banner"
const ClearTypeColour ClearType = "Colour"

func (c ClearType) IsIcon() bool {
	return c == ClearTypeIcon
}

func (c ClearType) IsDescription() bool {
	return c == ClearTypeDescription
}

func (c ClearType) IsBanner() bool {
	return c == ClearTypeBanner
}

func (c ClearType) IsColour() bool {
	return c == ClearTypeColour
}

type RelationshipStatus string

const RelationshipBlocked RelationshipStatus = "Blocked"
const RelationshipBlockedOther RelationshipStatus = "BlockedOther"
const RelationshipFriend RelationshipStatus = "Friend"
const RelationshipIncoming RelationshipStatus = "Incoming"
const RelationshipNone RelationshipStatus = "None"
const RelationshipOutgoing RelationshipStatus = "Outgoing"
const RelationshipUser RelationshipStatus = "User"

func (r RelationshipStatus) IsBlocked() bool {
	return r == RelationshipBlocked
}

func (r RelationshipStatus) IsBlockedOther() bool {
	return r == RelationshipBlockedOther
}

func (r RelationshipStatus) IsFriend() bool {
	return r == RelationshipFriend
}

func (r RelationshipStatus) IsIncoming() bool {
	return r == RelationshipIncoming
}

func (r RelationshipStatus) IsNone() bool {
	return r == RelationshipNone
}

func (r RelationshipStatus) IsOutgoing() bool {
	return r == RelationshipOutgoing
}

func (r RelationshipStatus) IsUser() bool {
	return r == RelationshipUser
}

type ChannelType string

const ChannelTypeText ChannelType = "Text"
const ChannelTypeVoice ChannelType = "Voice"

func (c ChannelType) IsText() bool {
	return c == ChannelTypeText
}

func (c ChannelType) IsVoice() bool {
	return c == ChannelTypeVoice
}

type RoleClearType string

const RoleClearTypeColour RoleClearType = "Colour"

func (r RoleClearType) IsColour() bool {
	return r == RoleClearTypeColour
}
