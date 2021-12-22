// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package imageprocessing

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/fogleman/gg"
	"github.com/zeebo/errs"
)

// TypeFile defines the list of possible type of files.
type TypeFile string

const (
	// TypeFilePNG indicates that the type file is png.
	TypeFilePNG TypeFile = "png"
	// TypeFileJSON indicates that the type file is json.
	TypeFileJSON TypeFile = "json"
)

// LayerComponentsCount searches count files in the specified path and by name of file.
func LayerComponentsCount(pathToLayerComponents, nameFile string) (int, error) {
	files, err := ioutil.ReadDir(pathToLayerComponents)
	if err != nil {
		return 0, fmt.Errorf(pathToLayerComponents + " - folder does not exist")
	}

	var count int
	for _, file := range files {
		isMatched, err := regexp.Match(fmt.Sprintf(nameFile, `\d`), []byte(file.Name()))
		if err != nil {
			return 0, err
		}
		if isMatched {
			count++
		}
	}

	return count, nil
}

// CreateLayer searches and decodes image to layer.
func CreateLayer(path, name string) (image.Image, error) {
	image, err := os.Open(filepath.Join(path, name))
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errs.Combine(err, image.Close())
	}()

	layer, err := png.Decode(image)
	if err != nil {
		return nil, err
	}
	return layer, err
}

// Layering overlays image layers on the base image.
func Layering(layers []image.Image, width, height int) *image.RGBA {
	var generalImage *image.RGBA
	for k, layer := range layers {
		if k == 0 {
			baseLayer := layer.Bounds()
			generalImage = image.NewRGBA(baseLayer)
			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
			continue
		}

		if layer != nil {
			offset := image.Pt(width, height)
			draw.Draw(generalImage, layer.Bounds().Add(offset), layer, image.Point{}, draw.Over)
		}
	}
	return generalImage
}

// SaveImage saves image by path.
func SaveImage(path, fullPath string, baseImage image.Image) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	resultImage, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	if err = png.Encode(resultImage, baseImage); err != nil {
		return err
	}
	defer func() {
		err = errs.Combine(err, resultImage.Close())
	}()

	return nil
}

// Inscription entity describes values required to apply inscription to the image.
type Inscription struct {
	Img         image.Image
	Width       int
	Height      int
	PathToFonts string
	FontSize    float64
	FontColor   string
	Text        string
	X           float64
	Y           float64
	TextAlign   float64
}

// ApplyInscription overlays the inscription on the image.
func ApplyInscription(inscription Inscription) (image.Image, error) {
	dc := gg.NewContext(inscription.Width, inscription.Height)
	if err := dc.LoadFontFace(inscription.PathToFonts, inscription.FontSize); err != nil {
		return nil, err
	}

	dc.SetHexColor(inscription.FontColor)
	dc.DrawImage(inscription.Img, 0, 0)
	dc.DrawStringAnchored(inscription.Text, inscription.X, inscription.Y, inscription.TextAlign, 0.5)
	dc.Clip()
	return dc.Image(), nil
}
