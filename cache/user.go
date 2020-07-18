package cache

import (
	"kiss2u/model"
)

type Users map[string]*model.User

func (s *Storage) PutUser(user *model.User) error {
	s.Users[user.Name] = user
	return nil
}

func (s *Storage) GetAllUsers() (model.Users, error) {
	var users model.Users
	for _, user := range s.Users {
		users = append(users, user)
	}
	return users, nil
}

func (s *Storage) FindUserPkg(key string) (*model.User, error) {
	if user := s.Users[key]; user != nil {
		return user, nil
	}

	return model.NewUser(), nil
}
