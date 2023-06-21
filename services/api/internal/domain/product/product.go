package product

import (
	
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
)



type Product struct{
	
	Id int
	Name string
	Price float32
	Quantity uint32
	User user.User
	CreatedAt string
	UpdateAt string

}
