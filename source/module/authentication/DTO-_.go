package authentication

type SignUpDTO struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=30"`
	LastName  string `json:"lastName" binding:"required,min=2,max=30"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8"`
	Phone     string `json:"phone" binding:"required"`
	DOB       string `json:"dob" binding:"required"`
}

type LogInDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


