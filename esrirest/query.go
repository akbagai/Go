package esrirest

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"



)


//in ESRI REST, you want the features[] - this is where results are stored (other info is metadata).
  type query struct{
  Features []features
  }

//Features has another [] of attributes. You can either hardoce all the values like "OBJECTID string" or if you want more generic, you can map it. then below you can iterate.
type features struct{
  Attributes map[string]interface{}
}

var data query

func Query(s string) []features {



        response, err := http.Get(s)
         if err != nil {
             fmt.Printf("%s", err)

         } else {
             defer response.Body.Close()
             c, _ := ioutil.ReadAll(response.Body)
             json.Unmarshal(c,&data)





          }
  return data.Features
}
