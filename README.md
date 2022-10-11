# ASSIGNMENT 2 (Build Rest API In GO)- Scalable Web Service with Golang - DTS Kominfo
By : Vika Putri Ariyanti

# Init Project
go mod init Assignment2

# Install Packages
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
go get -u github.com/gin-gonic/gin 

# Fitur
## Create Order
Path: [/orders ](http://localhost:8080/orders)

Method: Post 

Request Body :

{ "customerName":"Vika Putri",
    
    "items":[
    
        {
        
            "itemCode":"M00001",
            
            "description":"Nasi Goreng",
            
            "quantity":1
            
        }
        
    ]
    
}

## Get Orders
Path: [/orders ](http://localhost:8080/orders)

Method: GET

## Update Order
Path: [/orders ](http://localhost:8080/orders/:orderId)

Method:  PUT 
Request Body:

{

    "customerName":"Vika Putri",
    
    "items":[
    
        {
        
            "itemCode":"M00002",
            
            "description":"Nasi Goreng",
            
            "quantity":1
            
        }
        
    ]
    
}

## Delete Order
Path: [/orders ](http://localhost:8080/orders/:orderId)

Method: DELETE
