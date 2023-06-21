package mysql

import (
	
	"database/sql"
	"fmt"
	"strings"
	"time"
	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/user"

)


func (r *Repository) CreateUser(user *user.User) (*user.User, error){

	result, err := r.db.Exec("INSERT INTO User (first_name, last_name, phone, email, password, created_at, update_at) values(?,?,?,?,?,?,?)", 

		user.FirstName,
		user.LastName,
		user.Phone,
		user.Email,
		user.Password,
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

	newUser, err := r.GetByUserId(int(id))

	if err != nil{
		return nil, err
	}

	return newUser, nil
	
}

func (r *Repository) GetCollectionUser(queryParams *queryParametrs.QueryParams) ([]*user.User, error){
	
	var users []*user.User

	pre := queryParams.Page * queryParams.Limit - queryParams.Limit
	rows, err := r.db.Query("SELECT * FROM User LIMIT ?, ?", pre, queryParams.Limit)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	for rows.Next(){

		var user user.User

		if err := rows.Scan(

			&user.ID, 
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdateAt,
			
			); err != nil{
			return nil, err
		}

		users = append(users, &user)

	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return users, nil
	
}

func (r *Repository) GetByUserId(id int) (*user.User, error){

	var user user.User

	row := r.db.QueryRow("SELECT * FROM User WHERE id=?", id)

	if err := row.Scan(

		&user.ID, 
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdateAt,

		); err != nil{
		if err == sql.ErrNoRows{
			return nil, err
		}
		return nil, err
	}
	
	return &user, nil
} 

func (r *Repository) DeleteUser(id int) error{

	_, err := r.db.Exec("DELETE FROM User WHERE id=?", id)

	if err != nil{
		return err
	}

	return nil

}

func (r *Repository) UpdateUser(id int, updateUser user.User) (*user.User, error){

	_, err := r.db.Exec(
		"UPDATE User set first_name = ?, last_name = ?, update_at = ?, phone = ?, email = ?", 
		updateUser.FirstName, 
		updateUser.LastName,
		time.Now().Format(time.RFC3339),
		updateUser.Phone,
		updateUser.Email,
	)

	if err != nil{
		return nil, err
	}

	user, err := r.GetByUserId(id)

	if err != nil{
		return nil, err
	}

	return user, nil
	
}

func (r *Repository) GetUserBy(attribute map[string]string) (*user.User, error){

	var user user.User

	query := "SELECT * FROM User WHERE "
	for key, value := range attribute{

		query = query + key + "=" + "\"" + value + "\"" + " AND "

	}

	query = strings.TrimRight(query, "AND ")
	fmt.Println(query)

	row := r.db.QueryRow(query)

	if err := row.Scan(

		&user.ID, 
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdateAt,

	); err != nil{
		return nil, err
	}

	return &user, nil

}

