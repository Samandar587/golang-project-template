package app

import (
	"fmt"
	"golang-project-template/internal/models"
	"golang-project-template/internal/repository"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

/*
	var rootCmd = &cobra.Command{
		Use: "grpc-server",
		Run: func(cmd *cobra.Command, args []string) {
			// Application entrypoint...
			name, _ := cmd.Flags().GetString("name")
			fmt.Println("hello, ", name)
			fmt.Println("Salom, Dunyo!")

			// remove it when you run http/grpc server
			c := make(chan string)
			<-c
		},
	}
*/

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
		fmt.Println("connection is successfull")

		var user models.User
		row := db.QueryRow("select * from users where id = $1", 17)
		err = row.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("After QueryRow", user.Email, user.Id, user.Name)

		ch := make(chan string)
		<-ch
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
