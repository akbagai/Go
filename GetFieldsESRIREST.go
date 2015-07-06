package main

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
  
//Features has another [] of attributes. You can either hardcode all the values like "OBJECTID string" or if you want more generic, you can map it. then below you can iterate.
   
type features struct{
  Attributes map[string]interface{}
}

  func main(){

    response, err := http.Get("http://coagisweb.cabq.gov/arcgis/rest/services/public/PublicArt/MapServer/0/query?where=1=1&outFields=*&f=json")
       if err != nil {
           fmt.Printf("%s", err)

       } else {
           defer response.Body.Close()
           c, _ := ioutil.ReadAll(response.Body)

           var data query
           json.Unmarshal(c,&data)

           //  fmt.Println(data.Features[0].Attributes["OBJECTID"])  -->returns value for named key.
           for key, _ := range data.Features[0].Attributes{
                  //fmt.Println(data.Features[0].Attributes[key]) -->returns the values for each key
                  fmt.Println(key)  //print the list of named keys
           }


        }
  }
