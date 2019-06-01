package api

import (
	"Grocery-Shopping-Order-Module/src/app/handler"
	"Grocery-Shopping-User-Module/src/app/model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *Config) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.PortNo,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database %s",err.Error())
	}
	a.DB = DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.MobileNumber{})
	db.AutoMigrate(&model.Address{})
	return db
}
func (a *App) setRouters(){
	fmt.Println("initializing request")
	a.Get("/", a.init)
	a.Get("/order/{id}", a.GetOrder)
	a.Post("/order", a.CreateOrder)
	a.Get("/orderss", a.GetOrders)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) CreateOrder(w http.ResponseWriter, r *http.Request) {
	handler.CreateOrder(a.DB, w, r)
}

func (a *App) GetOrders(w http.ResponseWriter, r *http.Request) {
	handler.GetOrders(a.DB, w, r)
}

func (a *App) GetOrder(w http.ResponseWriter, r *http.Request) {
	handler.GetOrder(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App)init(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello This is User Module")
}