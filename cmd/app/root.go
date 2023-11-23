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
	Use:   "http-backend",
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
		userController := controller.NewUserController(userUsecase)

		r := router.NewRouter(userController)

		port := ":5005"
		log.Printf("Server listening on port %s...\n", port)
		log.Fatal(http.ListenAndServe(port, r))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
