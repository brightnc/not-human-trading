package repository

import (
	"time"

	"github.com/brightnc/not-human-trading/internal/core/domain"
)

var (
	clientTimeout time.Duration = 5000 * time.Millisecond
)

type User struct {
	url string
}

type user struct {
	id             string
	signature      string
	packageType    int
	isActive       bool
	expirationDate string
	expirationTime string
	createdDate    string
	createdTime    string
}

func NewUser() *User {
	return &User{}
}

func (r *User) RetrieveUserByID(userID string) (domain.User, error) {
	// client := httpclient.NewClient(httpclient.WithHTTPTimeout(clientTimeout))
	// res, err := client.Get(r.url, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// // Heimdall returns the standard *http.Response object
	// body, err := ioutil.ReadAll(res.Body)
	return domain.User{}, nil

}
