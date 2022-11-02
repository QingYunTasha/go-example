//postgres.go
package db

type Store interface {
   Insert(item interface{}) error
   Get(id int) error
}

type MyStore struct {
   db *sql.DB
}

func InitDB() Store { ... } //func to initialise DB
func (s *Store) Insert(item interface{}) error { ... } //insert item
func (s *Store) Get(id int) error { ... } //get item by id
//user.go
package user

type UserService struct {
   store db.Store
}

func NewUserService(s db.Store) *UserService {
   return &UserService{
      store: s,
   }
}
func (u *UserService) CreateUser() { ... }
func (u *UserService) RetrieveUser(id int) User { ... }