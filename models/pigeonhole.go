package models

import "ziweiBlog/config"

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]Post
}
