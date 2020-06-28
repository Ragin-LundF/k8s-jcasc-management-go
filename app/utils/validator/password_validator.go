package validator

func ValidateConfirmPasswords(password string, confirmPassword string) (isValid bool, errMessage string) {
	// check first, if both passwords are equal
	if password != confirmPassword {
		return false, "Passwords did not match!"
	}

	// check if password has a acceptable length (it is not enough, but better than nothing)
	if len(password) < 5 {
		return false, "Password length must be minimum 5 characters! Better will be more than 8!"
	}
	return true, ""
}
