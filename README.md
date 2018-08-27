## currencyapp ##

Currencyapp is designed using MVC pattern. 

Model are all stored in `internal` package.

View are all stored in `files/` package.

Controller are all stored in `package/controller` package

All incoming requests are handled in handler package. Package handler is designed to have many handler such as REST, RPC, and MQ messages. But for now only REST is used. Then it would be processed by controller. 
