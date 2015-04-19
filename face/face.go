package face

import (
	"github.com/LapisBlue/lapitar/mc"
	"github.com/LapisBlue/lapitar/util"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
)

const MinimumSize = 8

var DefaultScale = &imaging.NearestNeighbor

func Render(sk mc.Skin, size int, overlay bool, filter *imaging.ResampleFilter) image.Image {
	face := sk.GetFace(mc.Head, mc.Front)
	if overlay {
		helm := sk.OverlayFace(mc.Head, mc.Front)
		if !util.IsSolidColor(helm) {
			temp := imaging.Clone(face)
			draw.Draw(temp, temp.Bounds(), helm, helm.Bounds().Min, draw.Over)
			face = temp
		}
	}

	if size <= MinimumSize {
		return face
	}

	return imaging.Resize(face, size, size, *filter)
}
