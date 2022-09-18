# go-stac-api
STAC api written in go with fiber and mongodb   

To start mongodb on localhost:27017 with docker-compose   
```$ make database```    

To start the api on localhost:6001 with docker-compose    
```$ make api```   

To start the api on localhost:6001 outside of docker   
```$ go run main.go```   
   
To use mongo express for viewing the db on localhost:8081 with docker-compose  
```$ make express```   

Testing    
```$ make test```
   
Swagger/ OpenAPI   
http://localhost:6001/swagger/index.html#/  

Update docs    
```$ swag init```   

Notes:   
- this is not a fully functional stac api at this time    
- a public postman collection is available here:  
https://documenter.getpostman.com/view/12888943/UzBjtU96

![Alt text](data/swagger.png?raw=true "Postman Docs")