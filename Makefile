docker-run:
	docker-compose up --build -d

docker-stop:
	docker-compose down

docker-run-single:
	docker-compose -f docker-kafka-single.yml up -d

docker-stop-single:
	docker-compose -f docker-kafka-single.yml down	

docker-run-cluster:
	docker-compose -f docker-kafka-cluster.yml up -d

docker-stop-cluster:
	docker-compose -f docker-kafka-cluster.yml down	