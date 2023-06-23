package product

import (
	
	"encoding/json"
	"fmt"
	"net/http"
	jsonProduct "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/services/queryService"

)


func (uc *UseCase) Create(product *product.Product) (*product.Product, error){

	newProduct, err := uc.productRepository.CreateProduct(product)

	if err != nil{
		return nil, err
	}

	return newProduct, nil
}

func (uc *UseCase) GetCollection(r *http.Request) ([]*product.Product, error){

	queryParams := queryService.GetQueries(r)

	products, _ := uc.productRepository.GetCollectionProduct(queryParams)


	return products, nil

}

func (uc *UseCase) GetById(id int) (*product.Product, error){

	product, err := uc.productRepository.GetByProductId(id)

	if err != nil{
		return nil, err
	}

	return product, nil

}

func(uc *UseCase) Delete(id int) error{

	if err := uc.productRepository.DeleteProduct(id); err != nil{
		return err
	}

	return nil
	
}

func (uc *UseCase) Update(id int, updateproduct  jsonProduct.UpdateJsonProduct) (*product.Product, error){

	var updateProductMap map[string]string

	
	updateJsonProduct, _ := json.Marshal(updateproduct)

	json.Unmarshal(updateJsonProduct, &updateProductMap)

	fmt.Println(updateProductMap)

	Udatedproduct, err := uc.productRepository.UpdateProduct(id, updateProductMap)

	if err != nil{
		return nil, err
	}

	return Udatedproduct, nil

}