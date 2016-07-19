package main

import "fmt"
import "io"
import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "log"
import "os"



func main() {
	fmt.Printf("Starting Application...\n")

	/*
	 * setup the clientcredentials Configuration data
	 */
	conf := &clientcredentials.Config{
	    ClientID:     "---- Application Key goes Here ----",
	    ClientSecret: "---- Application Secret goes Here ----",
	    Scopes:       []string{},
	    TokenURL: "---- Your Bb server Root --- /learn/api/public/v1/oauth2/token",
	}

	/*
	 * Get a client based on the configuration
	 */
	client := conf.Client(oauth2.NoContext)

	/*
	 * Make a request to the API
	 */
	resp, err := client.Get("http://localhost:9876/learn/api/public/v1/users")
	if err != nil {
	    log.Fatal(err)
	}

	io.Copy(os.Stdout, resp.Body)

}

