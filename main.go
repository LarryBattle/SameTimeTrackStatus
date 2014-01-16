// Tracks users in Lotus Notes SameTime Instant Messager
// Created by: Larry Battle
// Version: 0.0.1
// @todo Turn this into a CLI tool with options: -users=id1,id2,id3 -output=file -verbose=bool -api_url=string
// @todo add check for Lotus Notes and setting.
package main

import (
	"os"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)
var (
	sametime_getstatus_URL = `http://localhost:59449/stwebapi/getstatus?userId=`
	outputFile = "status.txt"
	userId = "?"
)

func checkError(e error ){
	if e != nil {
		panic( e )
	}
}
func getSameTimeStatusOfUser( userId string ) []byte {
	res, err := http.Get( sametime_getstatus_URL + userId )
	checkError(err)
	defer res.Body.Close()
	json_response, err := ioutil.ReadAll(res.Body)
	checkError(err)
	return json_response
}
func addNewLine( x []byte) []byte{
	return append( x, []byte("\n")...)
}
func appendStringToFile( filename string, data []byte ){
	f, err := os.OpenFile( filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660);
	checkError(err)
	// @todo Find out if `f.WriteString( dataStr + "\n" )` works.
	_, err = f.Write( addNewLine( data ) )
	checkError(err)
}
func logSameTimeStatus( userId string ){
	appendStringToFile( outputFile, getSameTimeStatusOfUser( userId ) )
}
func startCounter( t time.Duration, fn func() ){
	i := 0
	c := time.Tick( t )
	for _ = range c {
		i++
		log.Printf( "%d: Status at %s\n", i, time.Now() )
		fn()
	}
}
func main() {
	log.Println( "Logging " + userId );
	startCounter( 5 * time.Minute, func(){
		logSameTimeStatus( userId )
	})
}
