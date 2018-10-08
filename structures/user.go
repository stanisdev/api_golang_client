package structures

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}