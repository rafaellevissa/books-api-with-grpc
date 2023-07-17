# Comandos para rodar:


## Abra dois terminais, rode esses comandos no 1° terminal:

> $ docker-compose up -d => sobe o arquivo docker-compose.yaml

> $ docker exec -it "kafka_container_name<apikafka_kafka_1>" bash => roda o kafka

> $ kafka-topics --list --bootstrap-server=localhost:9092 => lista os topicos

> $ kafka-topics --create --topic=<topic_name> --bootstrap-server=localhost:9092 --partitions=3 => cria topicos

> $ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=<topic_name> => faz com que o consumidor consuma do topico passado como parametro


## Comandos adicionais:

> $ docker-compose ps => lista os containers

> $ kafka-topics => possibilita ver todas opções referente ao uso dos topicos 

> $ kafka-topics --bootstrap-server=localhost:9092 --topic=topic_name --describe => descreve detalhado o topico passado no parametro

> $ kafka-console-consumer => possibilita ver todas opções referente a um consumidor

> $ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=topic_name --from-beginning => faz com que o consumidor consuma o topico passado por parametro por inteiro desde o inicio mas sem ordem devido ao parametro --from-beginning

> $ kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe => comando que descreve os grupos de consumidores

> $ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=topic_name --from-beginning --group=x => faz com que o consumidor consuma o topico passado por parametro por inteiro desde o inicio mas sem ordem e integra esse consumidor a um grupo especifico passasdo por parametro


## No 2° terminal rode esses:

> $ docker exec -it "kafka_container_name<apikafka_kafka_1>" bash => roda o kafka

> $ kafka-console-producer --bootstrap-server=localhost:9092 --topic=<topic_name> => faz com que um produtor esteja responsavel por enviar mensagens para o topico passado como parametro


## Ao final de tudo rode esse comando para derrubar os containers do docker:

> $ docker-compose down --remove-orphans