package actions

import "goapp/models"

func (as *ActionSuite) Test_UsersRegisterGet() {
	res := as.HTML("/users/register").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "<h3>Register</h3>")
}

func (as *ActionSuite) Test_UsersRegisterPost() {
	u := &models.User{Email: "ko@gmail.com", Username: "Lili"}
	res := as.HTML("/users/register").Post(u)
	as.Equal(302, res.Code)
	as.Equal("/", res.Location())

	err := as.DB.First(u)
	as.NoError(err)
	as.NotZero(u.ID)
	as.NotZero(u.CreatedAt)
	as.Equal("ko@gmail.com", u.Email)
}
