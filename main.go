package main

// import "bytes"
import "fmt"
import "io"
// import "encoding/json"
import "golang.org/x/oauth2"
// import "golang.org/x/net/context"
import "golang.org/x/oauth2/clientcredentials"
import "log"
// import "net/http"
import "os"

func main() {
	fmt.Printf("Starting...\n")


	// resp, err := http.Get("http://localhost:9876/learn/api/public/v1/courses")

	// if err != nil {

 //    	log.Fatal(err)
 //    } 
 //    if resp.StatusCode == 401 {
 //    	fmt.Printf("Getting new Auth Token...\n")
 //    	client := &http.Client{}

 //    	var grantTypeStr = []byte(`grant_type=client_credentials`)
 //    	req, err := http.NewRequest("POST", "http://localhost:9876/learn/api/public/v1/oauth2/token", bytes.NewBuffer(grantTypeStr));
 //    	if err != nil {
 //    		log.Fatal(err)
 //    	}
 //    	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
 //    	req.SetBasicAuth("bf6212ef-6bac-461d-b36e-90c0678dbca5","p6T1MTUB3sf89swDa8qfwHy8nXnZQ3FL")
 //    	resp, err := client.Do(req)

	// 	token := AuthToken{}
	// 	decoder := json.NewDecoder(resp.Body)
	// 	err = decoder.Decode(&token)

 //        if err != nil {
 //        	log.Fatal(err)
 //        }
 //    	// io.Copy(os.Stdout, resp.Body)
 //    	fmt.Printf("Bearer Token: %s", token.AccessToken)
 //    }
    



	conf := &clientcredentials.Config{
	    ClientID:     "bf6212ef-6bac-461d-b36e-90c0678dbca5",
	    ClientSecret: "p6T1MTUB3sf89swDa8qfwHy8nXnZQ3FL",
	    Scopes:       []string{},
	    TokenURL: "http://localhost:9876/learn/api/public/v1/oauth2/token",
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	
	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	// var code string
	// if _, err := fmt.Scan(&code); err != nil {
	//     log.Fatal(err)
	// }

	fmt.Printf("Getting Token...\n")
	// tok, err := conf.PasswordCredentialsToken(oauth2.NoContext, conf.ClientID, conf.ClientSecret)
	// tok, err := conf.Exchange(oauth2.NoContext, code)
	// if err != nil {
	    // log.Fatal(err)
	// }
	// fmt.Printf("Token: %s \n", )

	client := conf.Client(oauth2.NoContext)
	resp, err := client.Get("http://localhost:9876/learn/api/public/v1/courses")
	if err != nil {
	    log.Fatal(err)
	}
	io.Copy(os.Stdout, resp.Body)
    
}


type AuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}