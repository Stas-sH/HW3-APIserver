package server

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Db() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func GetAdminsFromDB() ([]Administrator, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	admins := make([]Administrator, 0)

	rows, err := db.Query("select * from admins")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		a := NewAdmin()

		err := rows.Scan(&a.Id, &a.AdminName, &a.Password)
		if err != nil {
			return nil, err
		}
		admins = append(admins, a)
	}
	return admins, nil
}

func GetUsersFromDB() ([]User, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	users := make([]User, 0)

	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := NewUser()

		err := rows.Scan(&u.Id, &u.UserName, &u.Mail, &u.Phone, &u.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func AddUserToDB(u User) error {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec("insert into users(username, mail, phone, password) values($1, $2, $3, $4)", u.UserName, u.Mail, u.Phone, u.Password)
	return err
}

func GetUserByIdFromDB(id string) (User, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return User{}, err
	}

	var u User
	err = db.QueryRow("select * from users where id = $1", id).Scan(&u.Id, &u.UserName, &u.Mail, &u.Phone, &u.Password)
	if err != nil {

		return User{}, err
	}
	return u, nil
}

func DeleteUserById(id string) error {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec("delete from users where id=$1", id)
	if err != nil {
		return err
	}

	return nil

}

func UpdateUserById(id string, user User) error {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET username=$1, mail=$2, phone=$3, password=$4 WHERE id=$5", user.UserName, user.Mail, user.Phone, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}
