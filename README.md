# Go
My Go projects

### C# GUI
Build this:
```go
package main
import (
  "fmt"
  "esrirest"
)
func main(){
d :=esrirest.Query("http://coagisweb.cabq.gov/arcgis/rest/services/public/PublicArt/MapServer/0/query?where=1=1&outFields=*&f=json")
for i:=0; i<len(d); i++{
  fmt.Println(d[i].Attributes["TITLE"])
}}
```
Then place the path to the EXE in the CsharpGUI file as FileName.

