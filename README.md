## currencyapp ##

Currencyapp is designed using MVC pattern. 

Model are all stored in `internal` package.

View are all stored in `files/` package.

Controller are all stored in `package/controller` package

All incoming requests are handled in handler package. Package handler is designed to have many handler such as REST, RPC, and MQ messages. But for now only REST is used. Then it would be processed by controller. 

## IF DOCKER IS NOT RUNNING PLEASE TEST THE APP WITH INSTRUCTION FOLLOWS ##

- DB installation
    please create db setting as follow, just run in in terminal
    `sudo -u postgres createuser test_user`
    `sudo -u postgres createdb test_currency` 
    `sudo -u postgres psql`
    `alter user test_user with encrypted password 'testpass';`
    `grant all privileges on database test_currency to test_user;`
    `\q`
    `psql test_user -h 127.0.0.1 -d test_currency`
    `set search_path to public;`
    `
        CREATE TABLE ws_currency(
        id SERIAL,
        from_curr VARCHAR NOT NULL,
        to_curr VARCHAR NOT NULL,
        status INT NOT NULL,
        PRIMARY KEY(from_curr, to_curr)
        );

        CREATE TABLE ws_currency_rates(
        id SERIAL,
        date DATE NOT NULL,
        from_curr VARCHAR NOT NULL,
        to_curr VARCHAR NOT NULL,
        rates decimal NOT NULL,
        PRIMARY KEY(date, from_curr, to_curr)
        );
    `

- dependency 
    after running the above commands, please run this on project working dir `github.com/nabillarahmani/currencyapp`

    `dep ensure -v`

- running the application

    `cd app`
    `go build`
    `sudo ./app`
