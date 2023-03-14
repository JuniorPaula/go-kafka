
# Microservice Golang and Apache Kafka

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![Apache Kafka](https://img.shields.io/badge/Apache%20Kafka-000?style=for-the-badge&logo=apachekafka)



A aplicação consiste em um **Microserviço** de cadastros de produtos e posterior persistência no banco de dados **mysql**. Sob duas formas de inserção, uma através do endpoint `[POST]/products`, e a outra publicando no topic `products` to **kafka**. Além de `[GET]/products` para consultar os produtos salvos.   

Esse é um estudo de casos utilizando **Golang** e testando toda a sua capacidade e performace em lidar com situações que exigem uma alta demanda de processamento e múltiplos ponto de conexão.

#### As tecnologias utilizadas foram:
- golang
- docker
- mysql
- apache kafka