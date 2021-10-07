package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {

	getdata()
	println("-----------------------------------------------------------")
	store()
}

func getdata() {

	cidread := "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"

	//Get Local Directory
	path, err2 := os.Getwd()
	if err2 != nil {
		fmt.Printf("error: %s", err2)
	}
	// Where your local node is running on localhost:5001

	sh := shell.NewShell("localhost:5001")
	cid := sh.Get(cidread, path)
	if cid != nil {
		fmt.Fprint(os.Stderr, "error: ", cid)
	}

	content, err := ioutil.ReadFile(cidread)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	fmt.Println(text)

}

func store() {

	//Get Local Directory
	path, err2 := os.Getwd()
	if err2 != nil {
		fmt.Printf("error: %s", err2)
	}

	//Instance of shell
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("This is What I Wrote by Aron J"))
	if err != nil {
		log.Fatal(err)
	}

	cidcont := sh.Get(cid, path)
	if cidcont != nil {
		fmt.Fprint(os.Stderr, "error: ", cidcont)
	}

	content, err := ioutil.ReadFile(cid)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	fmt.Println(text)
}
