package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var url string
	url = `https://api.mozambiquehe.re/bridge?version=5&platform=PC&uid=1008001556434&auth=xGapGoNDcuXhN3ggJAwt&enableClubsBeta`
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	type data struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Key    string `json:"key"`
		Global string `json:"global"`
	}

	type v struct {
		Global struct {
			Name string `json:"name"`
			UID  int    `json:"uid"`
		}
		Legends struct {
			Selected struct {
				Legendname string `json:"LegendName"`
				data       []data `json:"data"`
			}
		}
	}

	/*type v struct {
		Name string `json:"name"`
	}*/

	var value v
	//var dat map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&value)

	fmt.Println(value)
}
