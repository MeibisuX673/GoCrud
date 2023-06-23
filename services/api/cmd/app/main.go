package main

import (

	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/MeibisuX673/GoCrud/pkg/store/mysql"
	deliveryHttp "github.com/MeibisuX673/GoCrud/services/api/internal/delivery/http"
	repository "github.com/MeibisuX673/GoCrud/services/api/internal/repository/mysql"
	useCaseProduct "github.com/MeibisuX673/GoCrud/services/api/internal/useCase/product"
	useCaseUser "github.com/MeibisuX673/GoCrud/services/api/internal/useCase/user"
)

func main(){

	db, err := mysql.ConnectDb(mysql.ConfigMysql{})

	defer db.Close()

	if err != nil{
		panic(err)
	}

	var (

		repo  = repository.New(db)
		ucProduct = useCaseProduct.New(repo)
		ucUser   = useCaseUser.New(repo)
		listenerHttp = deliveryHttp.New(ucUser, ucProduct)

	)

	go func() {
		fmt.Println("service started successfully on http port: 8080")
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()
	
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}

