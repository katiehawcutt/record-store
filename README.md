# Record Store

##### November 2021

I started learning Go a few weeks ago and this is my first proper project using Go. I wanted to use it to get to grips with some basic CRUD operations using a Postgres DB and pgx (a Go toolkit for PostgreSQL). It runs in the terminal without any web interface or API's.

## Main Learning Points:
- how to structure a Go project
- how [pgx](https://github.com/jackc/pgx) works / what it offers
- tried out a validation library called [ozzo](https://github.com/go-ozzo/ozzo-validation)- how to use `os.Args` to access raw command-line arguments
- how to access .env variables in the project
- getting used to working with a strongly typed langauge
- recapped on SQL (I didn't want to use an ORM)

I'm sure there are lots of improvements that could be made but it's a start!

## Built with:
- Go

## Getting Started
- Firstly you will need to create and connect to a Postgres DB and add your env variables to the `.env` file
- Make sure you are in the root of the project
- Create the table - `psql nameOfYourDb -f ./db/scripts/create_table.sql`
- Insert data to the table - `psql nameOfYourDb -f ./db/scripts/insert_data.sql`
- Run the program - `go run cmd/main.go` (you should get a message saying `Welcome to Katie's Record Store!`)
- Now you know everything is working you can run the program followed by the operation you want to perform and the input data (if neccessary)...

    - get all the albums - `go run cmd/main.go getAlbums`
    - get a specific album - `go run cmd/main.go getAlbum 3` (gets the album with an id of 3)
    - add an album - `go run cmd/main.go addAlbum 21 Adele 4` (creates an album with the title '21', artist 'Adele', rating '4')
    - update an album - `go run cmd/main.go updateAlbum 11 21 Adele 5` (updates the album with an id of 11 to title '21', artist 'Adele', rating '5')
    - delete an album - `go run cmd/main.go deleteAlbum 1` (deletes the album with the id of 1)


- The validation I implemented means that:
    -  the rating has to be between 1 and 5
    - the title and artist have to be at least 1 character
