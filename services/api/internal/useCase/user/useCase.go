package user

import "github.com/MeibisuX673/GoCrud/services/api/internal/useCase/InterfaceForReposytory/repositoryInterface"

type UseCase struct{

	userRepository repositoryInterface.UserRepository

}

func New(repository repositoryInterface.UserRepository) *UseCase{

	useCase := &UseCase{
		userRepository: repository,
	}

	return useCase

}