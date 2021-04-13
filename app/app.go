package app

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB *gorm.DB
}

func (a *App) initializeRoutes(){
	a.Router.HandleFunc("/", a.Root).Methods("GET")
	a.Router.HandleFunc("/server", a.createBackendServer).Methods("POST")
	a.Router.HandleFunc("/server/{serverID}", a.getBackendServer).Methods("GET")
	a.Router.HandleFunc("/server", a.getBackendServer).Methods("GET")
	a.Router.HandleFunc("/server/{serverID}", a.updateBackendServer).Methods(http.MethodPut)
	//a.Router.HandleFunc("/server/{serverID}", a.deleteBackendServer).Methods(http.MethodDelete)
}

func (a *App) InitializeApp(user, password, dbname string){

	connectionString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?parseTime=True", user, password, dbname)

	var err error
	a.DB, err = gorm.Open("mysql", connectionString)
	if err != nil{
		log.Fatal(err)
	}

	a.DB.Exec("USE " + dbname)
	a.DB.AutoMigrate(&Server{}, &ServerEndpoint{})

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) Run(addr string){
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}