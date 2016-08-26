package main

import "io"
import "golang.org/x/oauth2"
import "golang.org/x/oauth2/clientcredentials"
import "log"
import "os"
import "flag"



var serverRoot string
var appKey string
var appSecret string
var tokenUrl string
var usersUrl string
var coursesUrl string

func init() {
	flag.StringVar(&serverRoot, "serverRoot", "", "The base URL of the Bb Learn server. e.g. https://mybb.inst.edu.au")
	flag.StringVar(&appKey, "appKey", "", "The Application Key")
	flag.StringVar(&appSecret, "appSecret", "", "The Application Secret")

	flag.Parse()

	if serverRoot == "" || appKey == "" || appSecret == "" {
		flag.Usage()
		os.Exit(1)
	}

	tokenUrl = serverRoot + "/learn/api/public/v1/oauth2/token"
	usersUrl = serverRoot + "/learn/api/public/v1/users"
	coursesUrl = serverRoot + "/learn/api/public/v1/courses"

}

func main() {

	/*
	 * setup the clientcredentials Configuration data
	 */
	conf := &clientcredentials.Config{
	    ClientID:     appKey,
	    ClientSecret: appSecret,
	    Scopes:       []string{},
	    TokenURL: tokenUrl,
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

