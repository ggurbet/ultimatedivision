// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatarcards

import (
	"ultimatedivision/cards"
	"ultimatedivision/cards/avatars"
)

// CardWithLinkToAvatar describes card entity with link to avatar.
type CardWithLinkToAvatar struct {
	cards.Card
	OriginalURL string `json:"originalUrl"`
}

// Config defines values needed to generate card with avatar.
type Config struct {
	CardConfig           cards.Config              `json:"cardConfig"`
	PercentageQualities  cards.PercentageQualities `json:"percentageQualities"`
	AvatarConfig         avatars.Config            `json:"avatarConfig"`
	PathToOutputJSONFile string                    `json:"pathToOutputJsonFile"`
	NameOutputJSONFile   string                    `json:"nameOutputJsonFile"`
	PathToNamesDataset   string                    `json:"pathToNamesDataset"`
}
