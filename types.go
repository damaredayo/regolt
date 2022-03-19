package regolt

type ClearType string

const ClearTypeIcon ClearType = "Icon"
const ClearTypeDescription ClearType = "Description"
const ClearTypeBanner ClearType = "Banner"
const ClearTypeColour ClearType = "Colour"

func (c ClearType) IsIcon() bool {
	return c == "Icon"
}

func (c ClearType) IsDescription() bool {
	return c == "Description"
}

func (c ClearType) IsBanner() bool {
	return c == "Banner"
}

func (c ClearType) IsColour() bool {
	return c == "Colour"
}

type RelationshipStatus string

const RelationshipBlocked RelationshipStatus = "Blocked"
const RelationshipBlockedOther RelationshipStatus = "BlockedOther"
const RelationshipFriend RelationshipStatus = "Friend"
const RelationshipIncoming RelationshipStatus = "Incoming"
const RelationshipNone RelationshipStatus = "None"
const RelationshipOutgoing RelationshipStatus = "Outgoing"
const RelationshipUser RelationshipStatus = "User"

func (c ClearType) IsBlocked() bool {
	return c == "Blocked"
}

func (c ClearType) IsBlockedOther() bool {
	return c == "BlockedOther"
}

func (c ClearType) IsFriend() bool {
	return c == "Friend"
}

func (c ClearType) IsIncoming() bool {
	return c == "Incoming"
}

func (c ClearType) IsNone() bool {
	return c == "None"
}

func (c ClearType) IsOutgoing() bool {
	return c == "Outgoing"
}

func (c ClearType) IsUser() bool {
	return c == "User"
}
