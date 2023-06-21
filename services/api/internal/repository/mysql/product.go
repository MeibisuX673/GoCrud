package mysql

import (

	"fmt"
	"strconv"
	"strings"
	"time"
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"
	"github.com/MeibisuX673/GoCrud/services/api/internal/repository/mysql/question"
	
)



func (r *Repository) CreateProduct(product *product.Product) (*product.Product, error){
	
	result, err := r.db.Exec("INSERT INTO Product (name, price, quantity, user_id, created_at, update_at) values(?, ?, ?, ?, ?, ?)", 
		product.Name,
		product.Price,
		product.Quantity,
		product.User.ID,
		time.Now().Format(time.RFC3339),
		time.Now().Format(time.RFC3339),
		
	)
	
	if err != nil{
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil{
		return nil, err
	}

	newProduct, err := r.GetByProductId(int(id))

	if err != nil{
		return nil, err
	}

	return newProduct, nil

}

func (r *Repository) GetCollectionProduct(queryParams *queryParametrs.QueryParams) ([]*product.Product, error){
	
	pre := queryParams.Page * queryParams.Limit - queryParams.Limit
	rows, err := r.db.Query("SELECT p.id, p.name, p.price, p.quantity, p.user_id, p.created_at, p.update_at FROM Product AS p JOIN User AS u ON p.user_id=u.id LIMIT ?, ?", pre, queryParams.Limit)

	if err != nil{
		return nil, err
	}

	var products []*product.Product

	for rows.Next(){

		var productQuestion question.ProductQuestion

		if err := rows.Scan(

			&productQuestion.Id,
			&productQuestion.Name,
			&productQuestion.Price,
			&productQuestion.Quantity,
			&productQuestion.UserId,
			&productQuestion.CreatedAt,
			&productQuestion.UpdateAt,	

		); err != nil{
			return nil, err
		}

		user := r.getUserProduct(productQuestion.UserId)

		product := &product.Product{

			Id: productQuestion.Id,
			Name: productQuestion.Name,
			Price: productQuestion.Price,
			Quantity: productQuestion.Quantity,
			CreatedAt: productQuestion.CreatedAt,
			UpdateAt: productQuestion.UpdateAt,
			User: user,

		}

		products = append(products, product)

	}

	return products, nil

}

func (r *Repository) GetByProductId(id int) (*product.Product, error){
	
	rowProduct := r.db.QueryRow("SELECT p.id, p.name, p.price, p.quantity, p.user_id, p.created_at, p.update_at FROM Product AS p JOIN User AS u ON p.user_id=u.id WHERE p.id=?", id)

	var productQuestion question.ProductQuestion

	if err := rowProduct.Scan(
		&productQuestion.Id,
		&productQuestion.Name,
		&productQuestion.Price,
		&productQuestion.Quantity,
		&productQuestion.UserId,
		&productQuestion.CreatedAt,
		&productQuestion.UpdateAt, 
	); err != nil{
		fmt.Println(err.Error())
	}

	user := r.getUserProduct(productQuestion.UserId)

	product := product.Product{
		Id: productQuestion.Id,
		Name: productQuestion.Name,
		Price: productQuestion.Price,
		Quantity: productQuestion.Quantity,
		CreatedAt: productQuestion.CreatedAt,
		UpdateAt: productQuestion.UpdateAt,
		User: user,
	}

	return &product, nil

}

func (r *Repository) DeleteProduct(id int) error{
	
	_, err := r.db.Exec("DELETE FROM Product WHERE id=?", id)

	if err != nil{
		return err
	}

	return nil
	
}

func (r *Repository) UpdateProduct(id int, arguments map[string]string) (*product.Product, error){
	
	query := "UPDATE Product set "

	for key, value := range arguments{

		if len(value) == 0{
			continue
		}

		query = query + key + "=" + "\"" + value + "\"" + " "

	}

	query = strings.TrimRight(query, " ")

	query = query + " WHERE id=" + strconv.Itoa(id)

	_, err := r.db.Exec(query)

	if err != nil{

		return nil, err

	}


	updatedProduct, err := r.GetByProductId(id)

	if err != nil{

		return nil, err
	}

	return updatedProduct, nil

}

func (r *Repository) getUserProduct(userId int) user.User{

	var user user.User

	rowUser := r.db.QueryRow("SELECT *  FROM User WHERE id=?", userId)

	if errUser := rowUser.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdateAt,
	); errUser != nil{
		fmt.Println(errUser.Error())
	}

	return user

}

