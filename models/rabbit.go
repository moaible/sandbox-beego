package models

import (
	"strconv"
	"time"
	"errors"
)

var (
	RabbitList map[string]*Rabbit
)

func init() {
	RabbitList = make(map[string]*Rabbit)
	u := Rabbit{"user_11111", "astaxie", "11111", 123}
	RabbitList["user_11111"] = &u
}

type Rabbit struct {
	Id       string
	Username string
	Password string
	Score int64
}

func AddRabbit(u Rabbit) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	RabbitList[u.Id] = &u
	return u.Id
}

func GetRabbit(uid string) (u *Rabbit, err error) {
	if u, ok := RabbitList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("Rabbit not exists")
}

func GetAllRabbit() map[string]*Rabbit {
	return RabbitList
}
