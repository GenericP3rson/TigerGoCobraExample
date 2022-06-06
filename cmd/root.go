package cmd

import (
   "fmt"
   "github.com/spf13/cobra"
   "github.com/GenericP3rson/TigerGo"
)

var rootCmd = &cobra.Command{
   Use:   "TigerGoCobraExample",
   Short: "Interact with MovieRec via the CLI",
   Long: `This CLI library is a sample for interacting with TigerGraph via a CLI. It uses TigerGo and Cobra and uses the Movie Recommendation Starter Kit`,
}

// Executes the root command.
func Execute() error {
   return rootCmd.Execute()
}

var conn = TigerGo.TigerGraphConnection{
   Token:     "",
   Host:      "https://movie-rec-cli.i.tgcloud.io",
   GraphName: "MyGraph",
   Username:  "tigergraph",
   Password:  "tigergraph",
}

func init() {
   conn.Token, _ = conn.GetToken() // Add Token
   
   // Add all commands to the root
   rootCmd.AddCommand(echo)
   rootCmd.AddCommand(addUser)
   rootCmd.AddCommand(addMovie)
   rootCmd.AddCommand(recMovies)
   rootCmd.AddCommand(rate)
}

var echo = &cobra.Command{ // Test if the connection is okay using echo!
   Use:   "echo",
   Short: "Test TigerGraph Connection",
   Long:  `Testing TigerGraph connection by calling the echo endpoint`,
   Run: func(cmd *cobra.Command, args []string) {
      res, err := conn.Echo() // Run echo command
      if err == nil {
         fmt.Print(res)
         fmt.Println("Success!")
      } else {
         fmt.Println(err.Error())
         fmt.Println("Uh oh! There's an error!")
      }
   },
}

var addUser = &cobra.Command{
   Use:   "addUser",
   Short: "Create a new user",
   Long:  `Upsert a new person vertex; parameters is userid`,
   Run: func(cmd *cobra.Command, args []string) {
      userid := args[0] // Grab the first argument
      res, err := conn.UpsertVertex("person", userid, map[string]interface{}{"id": userid}) // Upsert vertex
      if err == nil {
         fmt.Print(res)
         fmt.Println("Success!")
      } else {
         fmt.Println(err.Error())
         fmt.Println("Uh oh! There's an error!")
      }
   },
}

var addMovie = &cobra.Command{
   Use:   "addMovie",
   Short: "Adds a new movie",
   Long:  `Upsert a new movie vertex; parameters are movieid, movie title, and genres`,
   Run: func(cmd *cobra.Command, args []string) {
      movieid := args[0]
      moviename := args[1]
      genres := args[2]
      res, err := conn.UpsertVertex("movie", movieid, map[string]interface{}{"title": moviename, "genres": genres})
      if err == nil {
         fmt.Print(res)
         fmt.Println("Success!")
      } else {
         fmt.Println(err.Error())
         fmt.Println("Uh oh! There's an error!")
      }
   },
}

var rate = &cobra.Command{
   Use:   "rate",
   Short: "Create a new rating",
   Long:  `Upsert a rating edge; parameters are userid, movieid, and rating`,
   Run: func(cmd *cobra.Command, args []string) {
      userid := args[0]
      movieid := args[1]
      rating := args[2]
      res, err := conn.UpsertEdge("person", userid, "rate", "movie", movieid, map[string]interface{}{"rating": rating})
      if err == nil {
         fmt.Print(res)
         fmt.Println("Success!")
      } else {
         fmt.Println(err.Error())
         fmt.Println("Uh oh! There's an error!")
      }
   },
}

var recMovies = &cobra.Command{
   Use:   "recMovies",
   Short: "Recommend a movie",
   Long:  `Run the RecommendMovie query; parameters is userid`,
   Run: func(cmd *cobra.Command, args []string) {
      userid := args[0]
      res, err := conn.RunInstalledQuery("RecommendMovies", map[string]interface{}{"p": map[string]string{"id": userid, "type": "person"}, "k1": 100, "k2": 10})
      if err == nil {
         fmt.Print(res)
         fmt.Println("Success!")
      } else {
         fmt.Println(err.Error())
         fmt.Println("Uh oh! There's an error!")
      }
   },
}