package user

type UserStore interface{
	Insert(item interface{}) error
	Get(id int) error
}

type UserService struct {
	store UserStore
}

func NewUserService(s UserStore) *UserService {
	return &UserService{
        store: s,
    }
}

func (u *UserService) CreateUser(){
	...
}

func (u *UserService) RetrieveUser(id int){
    ...
}
