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

	"github.com/zeebo/errs"
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
	layer, err := png.Decode(image)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = errs.Combine(err, image.Close())
	}()
	return layer, nil
}

// Layering overlays image layers on the base image.
func Layering(layers []image.Image) *image.RGBA {
	var generalImage *image.RGBA
	for k, layer := range layers {
		if k == 0 {
			baseLayer := layer.Bounds()
			generalImage = image.NewRGBA(baseLayer)
			draw.Draw(generalImage, baseLayer, layer, image.Point{}, draw.Src)
			continue
		}

		if layer != nil {
			draw.Draw(generalImage, layer.Bounds(), layer, image.Point{}, draw.Over)
		}
	}
	return generalImage
}

// SaveImage saves image by path.
func SaveImage(fullPath string, baseImage image.Image) error {
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
