package app

import (
	"fmt"
	"golang-project-template/internal/delivery/http/controller"
	"golang-project-template/internal/delivery/http/router"
	"golang-project-template/internal/repository"
	"golang-project-template/internal/usecases"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backend",
	Short: "Run the backend server",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Println("Hello, ", name)
		fmt.Println("before connection...")
		db, err := repository.OpenDatabaseConnection()
		if err != nil {
			log.Println(err)
		}
		defer db.Close()

		userRepository := repository.NewUserRepository(db)
		userUsecase := usecases.NewUserUsecase(userRepository)

		// Create an instance of UserController
		userController := controller.NewUserController(userUsecase)

		// Create a new Gorilla Mux router
		r := router.NewRouter(userController)

		// Start the HTTP server
		port := ":5005"
		log.Printf("Server listening on port %s...\n", port)
		log.Fatal(http.ListenAndServe(port, r))
	},
}

/*
	func startHTTPServer(servive foobar) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, this is an HTTP server!")
		})

		log.Fatal(http.ListenAndServe(":5005", nil))
	}
*/
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
