package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gen2brain/beeep"
)


const PLATE_NUMBER = "NS686RM"

func main() {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://portal.parkingns.rs/portal/auth/checkPPK?platePr=%s", PLATE_NUMBER), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:121.0) Gecko/20100101 Firefox/121.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		err := beeep.Alert("Parking", "Ne radi dobro", "")
		if err != nil {
			panic(err)
		}
		return
	}
	respContent, _ := io.ReadAll(resp.Body)

	var body = []interface{}{}
	err = json.Unmarshal(respContent, &body)
	if err != nil {
		panic(err)
	}

	if len(body) > 0 {
		err := beeep.Alert("Parking", "Jeo si ga panju", "")
		if err != nil {
			panic(err)
		}
	}
}
