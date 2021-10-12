// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// ErrNoAvatar indicated that avatar does not exist.
var ErrNoAvatar = errs.Class("avatar does not exist")

// ErrNoAvatarFile indicated that avatar's file does not exist.
var ErrNoAvatarFile = errs.Class("avatar's file does not exist")

// DB is exposing access to avatars db.
//
// architecture: DB
type DB interface {
	// Create adds avatar in the data base.
	Create(ctx context.Context, avatar Avatar) error
	// Get returns avatar by id from the data base.
	Get(ctx context.Context, id uuid.UUID) (Avatar, error)
}

// Avatar entity describes the values that make up the avatar.
type Avatar struct {
	CardID         uuid.UUID   `json:"cardId"`
	PictureType    PictureType `json:"pictureType"`
	FaceColor      int         `json:"faceColor"`
	FaceType       int         `json:"faceType"`
	EyeBrowsType   int         `json:"eyeBrowsType"`
	EyeBrowsColor  int         `json:"eyeBrowsColor"`
	EyeLaserType   int         `json:"eyeLaserType"`
	HairstyleColor int         `json:"hairstyleColor"`
	HairstyleType  int         `json:"hairstyleType"`
	Nose           int         `json:"nose"`
	Tshirt         int         `json:"tshirt"`
	Beard          int         `json:"beard"`
	Lips           int         `json:"lips"`
	Tattoo         int         `json:"tattoo"`
	OriginalURL    string      `json:"originalUrl"`
	PreviewURL     string      `json:"previewUrl"`
}

// PictureType defines the list of possible type of picture.
type PictureType int

const (
	// PictureTypeFirst indicates the type of photo is the first.
	PictureTypeFirst PictureType = 1
)

// Config defines values needed by generate avatars.
type Config struct {
	PathToAvararsComponents   string `json:"pathToAvararsComponents"`
	PathToOutputAvatarsLocal  string `json:"pathToOutputAvatarsLocal"`
	PathToOutputAvatarsRemote string `json:"pathToOutputAvatarsRemote"`

	FaceColorFolder string `json:"faceColorFolder"`

	TattooFolder     string `json:"tattooFolder"`
	TattooTypeFolder string `json:"tattooTypeFolder"`
	TattooFile       string `json:"tattooFile"`

	FaceTypeFolder string `json:"faceTypeFolder"`
	FaceTypeFile   string `json:"faceTypeFile"`

	EyeBrowsFolder     string `json:"eyeBrowsFolder"`
	EyeBrowsTypeFolder string `json:"eyeBrowsTypeFolder"`
	EyeBrowsColorFile  string `json:"eyeBrowsColorFile"`

	EyeLaserFolder     string `json:"eyeLaserFolder"`
	EyeLaserTypeFolder string `json:"eyeLaserTypeFolder"`
	EyeLaserTypeFile   string `json:"eyeLaserTypeFile"`

	HairstyleFolder      string `json:"hairstyleFolder"`
	HairstyleColorFolder string `json:"hairstyleColorFolder"`
	HairstyleTypeFile    string `json:"hairstyleTypeFile"`

	NoseFolder     string `json:"noseFolder"`
	NoseTypeFolder string `json:"noseTypeFolder"`
	NoseFile       string `json:"noseFile"`

	BeardFolder string `json:"beardFolder"`
	BeardFile   string `json:"beardFile"`

	LipsFolder string `json:"lipsFolder"`
	LipsFile   string `json:"lipsFile"`

	TshirtFolder string `json:"tshirtFolder"`
	TshirtFile   string `json:"tshirtFile"`

	PercentageFacialFeatures struct {
		EyeLaser  int `json:"eyeLaser"`
		Hairstyle int `json:"hairstyle"`
		Beard     int `json:"beard"`
	} `json:"percentageFacialFeatures"`

	SizePreviewImage struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"sizePreviewImage"`
}

// TypeImage defines the list of possible type of avatar image.
type TypeImage string

const (
	// TypeImagePNG indicates that the type image avatar is png.
	TypeImagePNG TypeImage = "png"
)

// FormatImage defines the list of possible format of avatar image.
type FormatImage string

const (
	// FormatImageOriginal indicates that the format image avatar is original.
	FormatImageOriginal = "original"
	// FormatImagePreview indicates that the format image avatar is preview.
	FormatImagePreview = "preview"
)
