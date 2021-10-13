// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
)

// AvatarCards describes avatar card entity.
type AvatarCards struct {
	cards.Card
	OriginalURL string `json:"originalUrl"`
}

// Config defines values needed by generate avatar cards.
type Config struct {
	CardConfig           cards.Config              `json:"cardConfig"`
	PercentageQualities  cards.PercentageQualities `json:"percentageQualities"`
	AvatarConfig         avatars.Config            `json:"avatarConfig"`
	PathToOutputJSONFile string                    `json:"pathToOutputJSONFile"`
	NameOutputJSONFile   string                    `json:"nameOutputJSONFile"`
}
