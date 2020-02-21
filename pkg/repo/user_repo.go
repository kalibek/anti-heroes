package repo

import (
	"github.com/kalibek/anti-heroes/pkg/model"
)

func (r *Repo) SaveUser(user model.User) error {
	_, err := r.Exec("insert into users (email, name) values ($1, $2)", user.Email, user.Name)
	return err
}

func (r *Repo) UpdateUser(user model.User) error {
	_, err := r.Exec("update users set email = $1, name = $2 where id = $3",
		user.Email,
		user.Name,
		user.UserId)
	return err
}

func (r *Repo) GetUser(id int64) (*model.User, error) {
	user := new(model.User)
	row := r.QueryRow("select id, email, name from users where id = $1", id)
	err := row.Scan(&user.UserId, &user.Email, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
