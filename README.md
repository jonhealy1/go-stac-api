# go-stac-api
STAC api written in go with fiber and mongodb

To start mongodb on localhost:27017   
```$ make database```    

To start the api on localhost:6001   
```$ make api```   
   
To use mongo express for viewing the db on localhost:8081   
```$ make express```   

Testing    
```$ go test -v ./...```
   
Swagger/ OpenAPI   
http://localhost:6001/swagger/index.html#/  

Update docs    
```$ swag init```   

Notes:   
- this is not a fully functional stac api at this time    
- checkout the postman collection to test and see what works
https://documenter.getpostman.com/view/12888943/UzBjtU96

![Alt text](data/swagger.png?raw=true "Postman Docs")