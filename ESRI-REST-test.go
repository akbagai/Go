package main

import (
  "fmt"
  "esrirest"

)

func main(){

d :=esrirest.Query("http://coagisweb.cabq.gov/arcgis/rest/services/public/PublicArt/MapServer/0/query?where=1=1&outFields=*&f=json")
fmt.Println(d[0].Attributes["TITLE"])
      /* Iterate through the features and print TITLE
      for i:=0; i<len(d); i++{
           fmt.Println(d[i].Attributes["TITLE"])
      */
}
}