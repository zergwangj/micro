package models

import "database/sql"

type User struct {
	Id                   int64
	Username             string
	Password             string
}

func ListUsers(db *sql.DB, page int64, pageSize int64) ([]*User, error) {
	rows, err := db.Query("SELECT id, username, password FROM users LIMIT ?,?", (page - 1) * pageSize, pageSize)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0, pageSize)
	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(db *sql.DB, id int64) (*User, error) {
	row := db.QueryRow("SELECT id, username, password FROM users WHERE id=?", id)
	user := new(User)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(db *sql.DB, username string) (*User, error) {
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username)
	user := new(User)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user *User) (*User, error) {
	stmt, err := db.Prepare("INSERT INTO users (username, password) VALUES (?,?)")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	ret, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = id
	return user, nil
}

func DeleteUser(db *sql.DB, id int64) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return err
	}
	ret, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *sql.DB, user *User) (*User, error) {
	stmt, err := db.Prepare("UPDATE users SET username=?, password=? WHERE id=?")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	ret, err := stmt.Exec(user.Username, user.Password, user.Id)
	if err != nil {
		return nil, err
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return nil, err
	}
	return user, nil
}