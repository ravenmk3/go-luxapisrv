package service

import (
	"github.com/iawia002/lux/extractors"
)

func Extract(url string) ([]*extractors.Data, error) {
	options := extractors.Options{
		Playlist:         false,
		Items:            "",
		ItemStart:        0,
		ItemEnd:          0,
		ThreadNumber:     0,
		Cookie:           "",
		EpisodeTitleOnly: false,
		YoukuCcode:       "",
		YoukuCkey:        "",
		YoukuPassword:    "",
	}
	return extractors.Extract(url, options)
}
