// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package avatars

import (
	"context"
	"fmt"
	"image"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"github.com/zeebo/errs"

	"ultimatedivision/cards"
	"ultimatedivision/pkg/imageprocessing"
	"ultimatedivision/pkg/rand"
)

// ErrAvatar indicated that there was an error in service.
var ErrAvatar = errs.Class("avatar service error")

// Service is handling avatars related logic.
//
// architecture: Service
type Service struct {
	avatars DB
	config  Config
}

// NewService is a constructor for avatars service.
func NewService(avatars DB, config Config) *Service {
	return &Service{
		config:  config,
		avatars: avatars,
	}
}

// Create adds avatar in DB.
func (service *Service) Create(ctx context.Context, avatar Avatar) error {
	return ErrAvatar.Wrap(service.avatars.Create(ctx, avatar))
}

// Generate generates a common avatar from different layers of photos.
func (service *Service) Generate(ctx context.Context, card cards.Card, nameFile string) (Avatar, error) {
	var (
		layer                 image.Image
		layers                []image.Image
		originalAvatarsLayers []image.Image
		count                 int
		err                   error
	)

	avatar := Avatar{
		CardID:      card.ID,
		PictureType: PictureTypeFirst,
	}

	// FaceColor
	if count, err = imageprocessing.LayerComponentsCount(service.config.PathToAvararsComponents, service.config.FaceColorFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.FaceColor, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// FaceType
	pathToFaceColor := filepath.Join(service.config.PathToAvararsComponents, fmt.Sprintf(service.config.FaceColorFolder, avatar.FaceColor))
	if count, err = imageprocessing.LayerComponentsCount(pathToFaceColor, service.config.FaceTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.FaceType, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	pathToFaceType := filepath.Join(pathToFaceColor, fmt.Sprintf(service.config.FaceTypeFolder, avatar.FaceType))
	if layer, err = imageprocessing.CreateLayer(pathToFaceType, fmt.Sprintf(service.config.FaceTypeFile, avatar.FaceType)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// NoseType
	pathToNoseType := filepath.Join(pathToFaceType, service.config.NoseFolder)
	if count, err = imageprocessing.LayerComponentsCount(pathToNoseType, service.config.NoseTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Nose, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	pathToNoseType = filepath.Join(pathToNoseType, fmt.Sprintf(service.config.NoseTypeFolder, avatar.Nose))
	if layer, err = imageprocessing.CreateLayer(pathToNoseType, fmt.Sprintf(service.config.NoseFile, avatar.Nose)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// LipsType
	pathToLipsType := filepath.Join(pathToNoseType, service.config.LipsFolder)
	if count, err = imageprocessing.LayerComponentsCount(pathToLipsType, service.config.LipsFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Lips, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToLipsType, fmt.Sprintf(service.config.LipsFile, avatar.Lips)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// EyeBrowsType
	pathToEyeBrowsType := filepath.Join(pathToFaceType, service.config.EyeBrowsFolder)
	if count, err = imageprocessing.LayerComponentsCount(pathToEyeBrowsType, service.config.EyeBrowsTypeFolder); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.EyeBrowsType, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// EyeBrowsColor
	pathToBrowsColor := filepath.Join(pathToEyeBrowsType, fmt.Sprintf(service.config.EyeBrowsTypeFolder, avatar.EyeBrowsType))
	if count, err = imageprocessing.LayerComponentsCount(pathToBrowsColor, service.config.EyeBrowsColorFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.EyeBrowsColor, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToBrowsColor, fmt.Sprintf(service.config.EyeBrowsColorFile, avatar.EyeBrowsColor)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// Tattoo
	if card.IsTattoo {
		pathToTattoo := filepath.Join(service.config.PathToAvararsComponents, service.config.TattooFolder, fmt.Sprintf(service.config.TattooTypeFolder, avatar.FaceType))
		if count, err = imageprocessing.LayerComponentsCount(pathToTattoo, service.config.TattooFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.Tattoo, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToTattoo, fmt.Sprintf(service.config.TattooFile, avatar.Tattoo)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// Hairstyles
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.Hairstyle) {
		// HairstylesColor
		pathToHairstylesColor := filepath.Join(pathToFaceType, service.config.HairstyleFolder)
		if count, err = imageprocessing.LayerComponentsCount(pathToHairstylesColor, service.config.HairstyleColorFolder); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.HairstyleColor, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		// HairstylesType
		pathToHairstylesType := filepath.Join(pathToHairstylesColor, fmt.Sprintf(service.config.HairstyleColorFolder, avatar.HairstyleColor))
		if count, err = imageprocessing.LayerComponentsCount(pathToHairstylesType, service.config.HairstyleTypeFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.HairstyleType, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToHairstylesType, fmt.Sprintf(service.config.HairstyleTypeFile, avatar.HairstyleType)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// T-shirtType
	pathToTshirtType := filepath.Join(pathToFaceType, service.config.TshirtFolder)
	if count, err = imageprocessing.LayerComponentsCount(pathToTshirtType, service.config.TshirtFile); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	if avatar.Tshirt, err = rand.RandomInRange(count); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	if layer, err = imageprocessing.CreateLayer(pathToTshirtType, fmt.Sprintf(service.config.TshirtFile, avatar.Tshirt)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	layers = append(layers, layer)

	// BeardType
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.Beard) {
		pathToBeardType := filepath.Join(pathToNoseType, service.config.BeardFolder)
		if count, err = imageprocessing.LayerComponentsCount(pathToBeardType, service.config.BeardFile); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		if avatar.Beard, err = rand.RandomInRange(count); err != nil {
			return avatar, ErrAvatar.Wrap(err)
		}

		if layer, err = imageprocessing.CreateLayer(pathToBeardType, fmt.Sprintf(service.config.BeardFile, avatar.Beard)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	// EyeLaserType
	if rand.IsIncludeRange(service.config.PercentageFacialFeatures.EyeLaser) {
		pathToEyeLaserFolder := filepath.Join(pathToFaceType, service.config.EyeLaserFolder)
		avatar.EyeLaserType = avatar.EyeBrowsType

		pathToEyeLaserType := filepath.Join(pathToEyeLaserFolder, fmt.Sprintf(service.config.EyeLaserTypeFolder, avatar.EyeLaserType))
		if layer, err = imageprocessing.CreateLayer(pathToEyeLaserType, fmt.Sprintf(service.config.EyeLaserTypeFile, avatar.EyeLaserType)); err != nil {
			return avatar, ErrNoAvatarFile.Wrap(err)
		}
		layers = append(layers, layer)
	}

	originalAvatar := imageprocessing.Layering(layers, 0, 0)

	// Background
	pathToBackground := filepath.Join(service.config.PathToAvararsComponents, service.config.BackgroundFolder)
	if layer, err = imageprocessing.CreateLayer(pathToBackground, string(card.Quality)+"."+string(imageprocessing.TypeFilePNG)); err != nil {
		return avatar, ErrNoAvatarFile.Wrap(err)
	}
	originalAvatarsLayers = append(originalAvatarsLayers, layer)

	reducedOriginalAvatar := resize.Resize(uint(service.config.Sizes.OriginalAvatar.Width), uint(service.config.Sizes.OriginalAvatar.Height), originalAvatar, resize.Lanczos3)
	originalAvatarsLayers = append(originalAvatarsLayers, reducedOriginalAvatar)

	originalImage := imageprocessing.Layering(originalAvatarsLayers, service.config.IndentOriginalAvatar.Left, service.config.IndentOriginalAvatar.Above)

	var fontColors string
	switch card.Quality {
	case cards.QualityWood:
		fontColors = service.config.Inscriptions.FontColors.Wood
	case cards.QualitySilver:
		fontColors = service.config.Inscriptions.FontColors.Silver
	case cards.QualityGold:
		fontColors = service.config.Inscriptions.FontColors.Gold
	case cards.QualityDiamond:
		fontColors = service.config.Inscriptions.FontColors.Diamond
	default:
		return avatar, ErrAvatar.New("quality not exist")
	}

	// PlayerName
	inscriptionPlayerName := imageprocessing.Inscription{
		Img:         originalImage,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.PlayerName.FontSize,
		FontColor:   fontColors,
		Text:        card.PlayerName,
		X:           service.config.Inscriptions.PlayerName.X,
		Y:           service.config.Inscriptions.PlayerName.Y,
		TextAlign:   service.config.Inscriptions.PlayerName.TextAlign,
	}

	originalImageWithLabelPlayerName, err := imageprocessing.ApplyInscription(inscriptionPlayerName)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Tactics
	inscriptionTac := imageprocessing.Inscription{
		Img:         originalImageWithLabelPlayerName,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Tactics),
		X:           service.config.Inscriptions.GameCharacteristics.Tac.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Tac.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelTac, err := imageprocessing.ApplyInscription(inscriptionTac)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Physique
	inscriptionPhy := imageprocessing.Inscription{
		Img:         originalImageWithLabelTac,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Physique),
		X:           service.config.Inscriptions.GameCharacteristics.Phy.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Phy.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelPhy, err := imageprocessing.ApplyInscription(inscriptionPhy)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Technique
	inscriptionTec := imageprocessing.Inscription{
		Img:         originalImageWithLabelPhy,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Technique),
		X:           service.config.Inscriptions.GameCharacteristics.Tec.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Tec.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelTec, err := imageprocessing.ApplyInscription(inscriptionTec)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Offence
	inscriptionOff := imageprocessing.Inscription{
		Img:         originalImageWithLabelTec,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Offence),
		X:           service.config.Inscriptions.GameCharacteristics.Off.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Off.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelOff, err := imageprocessing.ApplyInscription(inscriptionOff)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Defence
	inscriptionDef := imageprocessing.Inscription{
		Img:         originalImageWithLabelOff,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Defence),
		X:           service.config.Inscriptions.GameCharacteristics.Def.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Def.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelDef, err := imageprocessing.ApplyInscription(inscriptionDef)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	// Goalkeeping
	inscriptionGk := imageprocessing.Inscription{
		Img:         originalImageWithLabelDef,
		Width:       service.config.Sizes.Background.Width,
		Height:      service.config.Sizes.Background.Height,
		PathToFonts: service.config.Inscriptions.PathToFonts,
		FontSize:    service.config.Inscriptions.GameCharacteristics.FontSize,
		FontColor:   fontColors,
		Text:        strconv.Itoa(card.Goalkeeping),
		X:           service.config.Inscriptions.GameCharacteristics.Gk.X,
		Y:           service.config.Inscriptions.GameCharacteristics.Gk.Y,
		TextAlign:   service.config.Inscriptions.GameCharacteristics.TextAlign,
	}

	originalImageWithLabelGk, err := imageprocessing.ApplyInscription(inscriptionGk)
	if err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	avatar.OriginalURL = fmt.Sprintf(service.config.PathToOutputAvatarsRemote, nameFile)
	if err = imageprocessing.SaveImage(service.config.PathToOutputAvatarsLocal, filepath.Join(service.config.PathToOutputAvatarsLocal, nameFile+"."+string(imageprocessing.TypeFilePNG)), originalImageWithLabelGk); err != nil {
		return avatar, ErrAvatar.Wrap(err)
	}

	return avatar, nil
}

// Get returns avatar from DB.
func (service *Service) Get(ctx context.Context, cardID uuid.UUID) (Avatar, error) {
	avatar, err := service.avatars.Get(ctx, cardID)
	return avatar, ErrAvatar.Wrap(err)
}

// GetImage returns avatar image.
func (service *Service) GetImage(ctx context.Context, cardID uuid.UUID) ([]byte, error) {
	image, err := ioutil.ReadFile(filepath.Join(service.config.PathToOutputAvatarsLocal, cardID.String()+"."+string(imageprocessing.TypeFilePNG)))
	return image, ErrAvatar.Wrap(err)
}
