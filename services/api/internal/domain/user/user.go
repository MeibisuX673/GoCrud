package user



type User struct{

	ID int 
	FirstName string 
	LastName string 
	Phone string 
	Email string 
	Password string 
	CreatedAt string 
	UpdateAt string 

}

func New(

	firstName, lastName, phone, email, password string,


	) *User{

	return &User{

		FirstName: firstName,
		LastName: lastName,
		Phone: phone,
		Email: email,
		Password: password,

	}

}