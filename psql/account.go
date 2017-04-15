package psql

type AccountDB interface {
	FindById(id int) (*Account, error)
	FindByEmail(email string) (*Account, error)
}

type Account struct {
	Id          int
	Role        string
	Surmane     string
	Name        string
	Patronymic  string
	DateOfBirth string
	Email       string
	Password    string
	Photo       string
	Phone       string
}

func (db *DB) FindById(id int) (*Account, error) {
	var account Account
	row := db.QueryRow("SELECT * FROM account WHERE id=$1", id)
	err := row.Scan(&account.Id, &account.Role, &account.Surmane, &account.Name, &account.Patronymic,
		&account.DateOfBirth, &account.Email, &account.Password, &account.Photo, &account.Phone)
	return &account, err
}

func (db *DB) FindByEmail(email string) (*Account, error) {
	var account Account
	row := db.QueryRow("SELECT * FROM account WHERE email=$1", email)
	err := row.Scan(&account.Id, &account.Role, &account.Surmane, &account.Name, &account.Patronymic,
		&account.DateOfBirth, &account.Email, &account.Password, &account.Photo, &account.Phone)
	return &account, err
}