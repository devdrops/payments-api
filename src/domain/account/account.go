package account

import (
	"fmt"
	"regexp"
	"time"
)

type Account struct {
	Id        int       `json:"id"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"-"`
}

func (acc *Account) Valid() bool {
	// A Document will be valid if it's length is between 11 and 14 characters.
	if len(acc.Document) < 11 || len(acc.Document) > 14 {
		return false
	}
	// A Document will be valid if they have only digits and/or letters.
	if ok, err := regexp.Match("[^0-9A-Za-z]", []byte(acc.Document)); ok || err != nil {
		return false
	}

	return true
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account [ Id: %d, Document: %s, Created At: %v ]", acc.Id, acc.Document, acc.CreatedAt)
}
