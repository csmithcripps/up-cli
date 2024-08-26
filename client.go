package client


import (
  "fmt"
  "net/url"
  "net/http"
  "log"
)


var URL string = "https://api.up.com.au/api/v1"
var PAGE_SIZE int = 100


type Client struct {
  token string

}


func New(token string) Client {
  c := Client {token}
  return c
}


func (c Client) get(path string, params url.Values) {
  uri := URL + path
  
  req,err := http.NewRequest(http.MethodGet, uri, params)
  if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

  
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    log.Printf("Request Failed: %s", err)
    return err
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)  // Log the request body 
  bodyString := string(body)
  log.Print(bodyString)
}

func 
