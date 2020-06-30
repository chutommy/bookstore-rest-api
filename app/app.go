package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"tommychu/workdir/027_api-example-v2/app/routing"
	services "tommychu/workdir/027_api-example-v2/app/services/db"
	"tommychu/workdir/027_api-example-v2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// App is an application struct.
type App struct {
	Log    io.Writer
	Srv    *http.Server
	Router *gin.Engine
	DB     *gorm.DB
}

// New returns a new App.
func New() *App {
	return &App{}
}

// Initialize applies the config of the App.
func (a *App) Initialize(cfg *config.Config) {

	// log
	a.Log = cfg.Log.Output

	// database
	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.DBName,
		cfg.DB.Password,
	)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(fmt.Errorf("could Initialize a db connection: %v", err))
	}
	a.DB = services.DBMigrate(db)
	db.LogMode(false) // idsable annoying database logs

	// router
	a.Router = routing.GetRouter(cfg, a.DB)

	// server
	a.Srv = &http.Server{
		Addr:              cfg.Srv.Addr,
		Handler:           a.Router, // router
		ReadTimeout:       cfg.Srv.ReadTimeout,
		ReadHeaderTimeout: cfg.Srv.ReadHeaderTimeout,
		WriteTimeout:      cfg.Srv.WriteTimeout,
		IdleTimeout:       cfg.Srv.IdleTimeout,
		MaxHeaderBytes:    cfg.Srv.MaxHeaderBytes,
	}
}

// Close takes care of the whole application closure.
func (a *App) Close() []error {
	return []error{
		a.DB.Close(),
	}
}

// Run starts the application server.
func (a *App) Run() {
	fmt.Fprintf(a.Log, "Listening and serving HTTP on %s\n", a.Srv.Addr)
	log.Fatal(a.Srv.ListenAndServe())
}