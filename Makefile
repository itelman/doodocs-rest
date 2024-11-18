build:
	docker build -t doodocs .
run-img:
	docker run --name=archive -p 8081:8081 --rm -d doodocs 
run:
	go run .
stop:
	docker stop doodocs