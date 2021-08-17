## Setting up the application
    1. docker-compose up #start db and mountebank container
    2. source .env       #import env vars
    3. go run main.go    #run go api

## Accessing api
    application will start at port 8080
        1.To getBook from database
            * localhost:8080/book/:id
        2.To getBank from usio
            * localhost:8080/usio/bank
                Headers
                    1. Authorization {token value}
        
    
