package stream

import (
	"fmt"
	"time"
)

type Item struct {
	UserID       string
	UserLogin    string
	UserName     string
	GameName     string
	Title        string
	ViewerCount  int
	StartedAt    time.Time
	ThumbnailURL string
	Preview      string
}

func (s Item) FilterValue() string {
	return fmt.Sprintf("%s %s %s", s.UserLogin, s.GameName, s.Title)
}
