package storage

import (
	"encoding/json"

	"kiss2u/model"

	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	userPrefix = "user."
)

func getUserKeyName(name string) []byte {
	return []byte(userPrefix + name)
}

func (s *Storage) PutUser(user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return s.kv.Put(getUserKeyName(user.Name), data, nil)
}

func (s *Storage) GetAllUsers() (model.Users, error) {
	iter := s.kv.NewIterator(util.BytesPrefix(getUserKeyName("")), nil)
	var users model.Users
	for iter.Next() {
		var user model.User
		json.Unmarshal(iter.Value(), &user)
		users = append(users, &user)
	}
	iter.Release()
	return users, iter.Error()
}

func (s *Storage) FindUserPkg(key string) (*model.User, error) {
	data, err := s.kv.Get(getUserKeyName(key), nil)
	user := model.NewUser()
	err = json.Unmarshal(data, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
