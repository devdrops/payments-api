package account

import (
	"payments-api/internal/database"
)

type AccountRepository struct {
	db *database.Adapter
}

func NewRepository(adapter *database.Adapter) *AccountRepository {
	return &AccountRepository{
		db: adapter,
	}
}

func (a *AccountRepository) Create(doc string) error {
	st, err := a.db.Conn.Prepare("INSERT INTO accounts(document) VALUES( $1)")
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err := st.Exec(doc); err != nil {
		return err
	}

	return nil
}

func (a *AccountRepository) Get(id int) (Account, error) {
	var acc Account
	err := a.db.Conn.QueryRow("SELECT id, document FROM accounts WHERE id = $1", id).
		Scan(&acc.Id, &acc.Document)
	if err != nil {
		return acc, err
	}

	return acc, nil
}
