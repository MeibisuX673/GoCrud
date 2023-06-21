package product

import "github.com/MeibisuX673/GoCrud/services/api/internal/useCase/InterfaceForReposytory/repositoryInterface"

type UseCase struct{

	productRepository repositoryInterface.ProductRepository

}

func New(repository repositoryInterface.ProductRepository) *UseCase{

	useCase := &UseCase{
		productRepository: repository,
	}

	return useCase

}