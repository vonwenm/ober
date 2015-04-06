package main

import(
	"flag"
	"net/http"
	"fmt"
  "strconv"
)

func main(){

  //Flags, invoked with $ ober --help
  verbose := flag.Bool( "verbose", false, "turns on verbose mode" )
	fileName := flag.String( "file", "", "used to specify the file you want to serve (required)" )
  port := flag.Int( "port", 3232, "used to specify the port you want to serve on (defaults to 3232)" )

	flag.Parse()

  if *fileName == ""{
    panic( "You should atleast specify a file" )
  }

	if *verbose{
		fmt.Printf( "Serving file: %s", *fileName )
	}

	http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ){
		http.ServeFile( res, req, *fileName )
	})

	panic(http.ListenAndServe(":" + strconv.Itoa( *port ), nil))
}
