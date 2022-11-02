package db

type Store struct {
	db *sql.DB
}

func NewDB() *Store {...}

func (s *Store) Insert(item interface{}) error{
	...
}

func (s *Store) Get(id int) error{
	...
}