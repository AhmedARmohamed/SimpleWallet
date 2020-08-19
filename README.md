SimpleWallet is simple wallet service.


Features include: 

    1: User Registration and login.
    2: Wallet Creation.
    3: Deposit.
    4: Withdrawal.
    5: Balance Enquiry.
    
Installation.
Should have Go and Postgres installed

Clone the repository.

git clone https://github.com/AhmedARmohamed/SimpleWallet.git

Getting the Dependencies used in th project.

These are the package dependecies we will need.

    badoux/checkmail - for validating user emails.
    
    dgrijalva/jwt-go - to sign and verify jwt tokens.
    
    gorilla/mux - it is a router and dispatcher, for matching URLs to their handlers.
    
    crypto - to hash and verify user passwords.
    
    github.com/lib/pq - postgres database drive

To install these dependencies, open the terminal and type go get github.com/{package-name}
Example-    go get github.com/lib/pq .

However for crypto installation, type go get golang.org/x/crypto .

Getting Started.

Environment variables.
Create a postgres database to store data. 

The .env file keeps the database credentials and secrets like the secret for signing jwt tokens.
Replace its values with the actual values.

    SECRET=anything-secret-and-should-be-hard-to-guess
    DB_HOST=127.0.0.1
    DB_USER=username
    DB_PASSWORD=user-password
    DB_NAME=database-name
    DB_PORT=5432 # default port number

Api/Endpoint Specifications

Endpoints that are used in this project include:

    
    Request	Endpoints	Functionality
    
    POST	/register	        User Signup ( firstname, lastname, email, password)
    
    POST	/login	            User Login ( email, password)
    
    POST	/api/wallet	        Add wallet ( amount)
    
    POST	/api/deposit/id	    deposit money to the wallet ( amount )
    
    GET	/api/balance/id	        View Balance
    
    POST	/api/wallet/id	    Withdraw money from the  Wallet ( amount )
    

Running application
Change directory into SimpleWallet then

$ go run main.go
API endpoint can be accessed. Via http://localhost:8020/
