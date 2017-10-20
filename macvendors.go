// Thanks to the fine folks at https://macvendors.com/

package main

import(
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"os"
	"io"
	"bytes"
	"flag"
)

const MACVENDORS_API_URL = "http://api.macvendors.com"

func main() {
	// Flag config
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s MAC\n", filepath.Base(os.Args[0]) )
		flag.PrintDefaults()
	}
	
	flag.Parse()
	
	// Main logic
	
	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Please input a MAC adress\n")
		flag.Usage()
		os.Exit(2)
	}
	
	main_url, err := url.Parse(MACVENDORS_API_URL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error while parsing main url: %v\n", err)
		os.Exit(1)
    }
    
    mac_url, err := url.Parse("/" + flag.Arg(0))
    if err != nil {
        fmt.Fprintf(os.Stderr, "error while parsing MAC for url: %v\n", err)
		os.Exit(1)
    }
	
	resp, err := http.Get( main_url.ResolveReference(mac_url).String() )
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while requesting: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	
	buff := bytes.NewBuffer(nil)
	
	_, err = io.Copy(buff, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error while doenloading: %v\n", err)
		os.Exit(1)
	}
	
	fmt.Println(buff.String())
}
