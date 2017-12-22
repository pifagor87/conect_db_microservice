# A simple connect Postgresql DB implemented using pgx.
For DB works, we use a lightweight high-performance DB driver pgx.
To start using "conect_db_microservice" you need:
* Install and configure the latest version of Golang
* Be sure to modify the data access data in the connect_db.json to more complex.
* Move the connect_db.json file to a secure location on the server. Provide the necessary access rights.
In file main.go, specify the correct path to the file data.json. Replace "./connect_db.json" with the desired path to the file.
* To use "conect_db_microservice", add in import - "github.com/pifagor87/conect_db_microservice".
When constructing your own microservice, use them in the following way, for example:

  const AccessDbPatch string = "./connect_db.json"

  db := ConectPosqgresqlDb(AccessDbPatch)
  
  rows, err := db.Query("SELECT name FROM users LIMIT 1")
  "SELECT name FROM users LIMIT 1" - it is your sql code

## Dependencies
* github.com/jackc/pgx