package model

type Check struct {
	Time     string `json:"time"`
	Path     string `json:"path"`
	Duration int64  `json:"duration"`
	Timeout  int    `json:"timeout"`
}
