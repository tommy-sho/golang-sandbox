package mysql

import (
	"database/sql"
	"strconv"
	"time"

	"golang.org/x/xerrors"

	_ "github.com/go-sql-driver/mysql"
	db "github.com/tommy-sho/golang-sandbox/test/error-testing/database"
	"github.com/tommy-sho/golang-sandbox/test/error-testing/domain"
)

var (
	NotExistErr = xerrors.New("record not exists")
)

type database struct {
	DB *sql.DB
}

func NewDatabase(DB *sql.DB) db.Database {
	return &database{DB}
}

func (d database) Create(user *domain.User) error {
	stmt, err := d.DB.Prepare("INSERT INTO users(id, firstname, lastname, age, created_at, updated_at) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	now := time.Now().String()
	_, err = stmt.Exec(user.ID, user.FirstName, user.LastName, strconv.Itoa(user.Age), now, now)
	if err != nil {
		return err
	}

	return nil
}
func (d database) GetByID(ID string) (*domain.User, error) {
	var user domain.User
	row, err := d.DB.Query("SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d database) UpdateUser(ID string, user *domain.User) error {
	stmt, err := d.DB.Prepare("UPDATE users SET firstname=?, lastname=?, age=? WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Age, user.ID)
	if err != nil {
		return err
	}

	return nil
}
