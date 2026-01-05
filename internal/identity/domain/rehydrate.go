package domain

func RehydrateUser(
	id int,
	userID UserID,
	name string,
	email string,
	passwordHash string,
	role string,
	occupation string,
	avatar string,
	isActive bool,
	isVerified bool,
) *User {
	e, _ := NewEmail(email)

	return &User{
		id:         userID,
		name:       name,
		email:      e,
		password:   Password{value: passwordHash},
		role:       Role(role),
		occupation: occupation,
		avatar:     avatar,
		isActive:   isActive,
		isVerified: isVerified,
	}
}
