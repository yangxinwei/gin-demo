CXX=go

TARGET=gin-demo-api

default: build

install:
	$(CXX) install

build:
	$(CXX) build

doc:
	swag init

clean:
	rm -f $(TARGET) *.o

swag:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/gin-swagger/swaggerFiles