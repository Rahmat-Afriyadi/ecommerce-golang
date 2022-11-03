package helper

import "golang.org/x/exp/slices"

func CheckImageType(f string) bool {
	imageType := []string{
		"image/jpeg",
		"image/png",
	}
	return slices.Contains(imageType, f)
}
