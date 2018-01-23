package conect_db_microservice

import (
  "os"
  "fmt"
  "errors"
  "io/ioutil"
  "github.com/jackc/pgx"
  "github.com/json-iterator/go"
)

/* DB connect error. */
const DbConnectError int = 34

/* json iterator variable. */
var json = jsoniter.ConfigCompatibleWithStandardLibrary

/* Accsess values. */
type AccessDbJson struct {
  Username   string `json:"username"`
  Password   string `json:"password"`
  Host       string `json:"host"`
  Database   string `json:"database"`
  Port       uint16 `json:"port"`
}

/* Load DB connect access. */
func LoadAccessUser(AccessDbPatch string) (username, password, host, database string, port uint16, err error) {
  file, err1 := ioutil.ReadFile(AccessDbPatch)
  if err1 != nil {
    fmt.Printf("File error: %v\n", err1)
    os.Exit(31)
  }
  data := AccessDbJson{}
  err2 := json.Unmarshal(file, &data)
  if err2 != nil {
    fmt.Println("error:", err2)
    os.Exit(32)
  }
  if data.Username == "" {
    return username, password, host, database, port, errors.New("No username!")
  }
  if data.Password == "" {
    return username, password, host, database, port, errors.New("No password!")
  }
  if data.Host == "" {
    return username, password, host, database, port, errors.New("No host!")
  }
  if data.Database == "" {
    return username, password, host, database, port, errors.New("No database!")
  }
  if data.Database == "" {
    return username, password, host, database, port, errors.New("No port!")
  }
  return data.Username, data.Password, data.Host, data.Database, data.Port, err
}

/* Add Posqgresql DB configs. */
func ExtractConfig(AccessDbPatch string) pgx.ConnConfig {
  username, pass, host, database, port, err := LoadAccessUser(AccessDbPatch);
  if err != nil {
    fmt.Println("error:", err)
    os.Exit(33)
  }
  var config pgx.ConnConfig
  config.Host = host
  config.User = username
  config.Password = pass
  config.Database = database
  config.Port = port
  return config
}

/* Add Posqgresql DB connect. */
func ConectPosqgresqlDb(AccessDbPatch string) (conn *pgx.Conn) {
  conn, err := pgx.Connect(ExtractConfig(AccessDbPatch))
  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to establish connection: %v\n", err)
    os.Exit(DbConnectError)
  }
  return conn
}