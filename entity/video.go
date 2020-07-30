package entity

//Person is ...
type Person struct {
	FirstName string `JSON:"firstname" binding:"required"`
	LastName  string `JSON:"lastname" binding:"required"`
	Age       int8   `JSON:"age" binding:"gte=1,lte=100"`
	Email     string `JSON:"email" binding:"required,email"`
}

// Video is ...
type Video struct {
	Title       string `JSON:"title" binding:"max=20" validate:"is-okey"`
	Description string `JSON:"description" binding:"max=20"`
	URL         string `JSON:"url" binding:"required,url"`
	Author      Person `JSON:"author" binding:"required"`
}
