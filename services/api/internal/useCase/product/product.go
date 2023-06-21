package product

import (

	"encoding/json"
	"fmt"

	"github.com/MeibisuX673/GoCrud/pkg/queryParametrs"
	jsonProduct "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http/dto/product"
	"github.com/MeibisuX673/GoCrud/services/api/internal/domain/product"

)


func (uc *UseCase) Create(product *product.Product) (*product.Product, error){

	newProduct, err := uc.productRepository.CreateProduct(product)

	if err != nil{
		return nil, err
	}

	return newProduct, nil
}

func (uc *UseCase) GetCollection(queryParams *queryParametrs.QueryParams) ([]*product.Product, error){

	products, err := uc.productRepository.GetCollectionProduct(queryParams)

	if err != nil{
		return nil, err
	}

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