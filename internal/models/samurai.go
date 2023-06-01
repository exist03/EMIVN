package models

import "fmt"

type Samurai struct {
	Username string
	Nickname string
	Owner    string
}

func (s Samurai) String() string {
	return fmt.Sprintf("\nТэг: %s\nИмя: %s\nДайме: %s\n", s.Username, s.Nickname, s.Owner)
}
