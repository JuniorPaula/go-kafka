package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-message-kafka/internal/infra/akafka"
	"go-message-kafka/internal/infra/repositories"
	"go-message-kafka/internal/infra/web"
	"go-message-kafka/internal/usecases"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repositories.NewProductRepositoryMysql(db)
	createProductUsecase := usecases.NewCreateProductUsecase(repository)
	listProductsUsecase := usecases.NewListProductsUsecase(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductsUsecase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHanlder)
	r.Get("/products", productHandlers.ListProductsHanlder)

	fmt.Println("start server on port 8000")
	http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecases.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)
		if err != nil {
			fmt.Println("error to get messages from kafka")
			continue
		}
		_, err = createProductUsecase.Execute(dto)
		if err != nil {
			fmt.Println(err)
		}
	}
}
