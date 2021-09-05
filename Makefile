start-service:
	minikube start 
	kubectl config use-context minikube
	kubectl config set-context --current --namespace=default
	make start-db-service
	sleep 60
	kubectl port-forward svc/postgres 3000:5432

stop-service:
	make stop-db-service
	minikube stop

start-server:
	go run main.go

stop-db-service:
	kubectl delete -n default deployment postgres   
	kubectl delete svc postgres

start-db-service:
	kubectl create -f config/local/postgres_deployment.yaml
	kubectl create -f config/local/postgres_service.yaml

