package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-message-kafka/internal/infra/akafka"
	"go-message-kafka/internal/infra/repositories"
	"go-message-kafka/internal/usecases"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repositories.NewProductRepositoryMysql(db)
	createProductUsecase := usecases.NewCreateProductUsecase(repository)

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
