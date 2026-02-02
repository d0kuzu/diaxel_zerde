package repository

type User struct {
	ID       string
	Role     string
	Email    string
	Password string
}

type UserRepo struct {
	users map[string]User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: map[string]User{
			"test@mail.com": {
				ID:       "user-1",
				Email:    "test@mail.com",
				Password: "$2a$10$QZ3...", // заглушка
			},
		},
	}
}

func (r *UserRepo) FindByEmail(email string) (User, bool) {
	u, ok := r.users[email]
	return u, ok
}

func (r *UserRepo) FindByID(userID string) (User, bool) {
	u, ok := r.users["test@mail.com"] //TODO: пока глушилка
	return u, ok
}
