package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	//"encoding/json"
)

func callAPI(header map[string]string, URL string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("Get", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("api-key", header["api-key"])
	req.Header.Set("clientID", header["clientID"])

	res, err := client.Do(req)

	return res, err

}

func main_old() {
	fmt.Println("Hello World")
	client, station := "196", "1108"
	sensorName := "Stage"

	header := map[string]string{
		"api-key":  "FBdDvhm9SPNxvUc3",
		"clientID": client,
	}

	start := "2020-01-01 01:00:00"
	end := "2020-01-02 01:00:00"

	fmt.Printf("Client #: %s\nStation ID: %s\nSensor Name: %s\nStart Date: %s\nEnd Date: %s\n", client, station, sensorName, start, end)

	url := fmt.Sprintf("http://www.hydrometcloud.com:8080/Data/rest/api/stations?clientId=%s", client)
	//url := fmt.Sprintf("http://www.hydrometcloud.com:8080/Data/rest/api/sensordata?stationId=%s&sensorName=%s&startTime=%s&endTime=%s", station, sensorName, start, end)

	res, err := callAPI(header, url)

	if err != nil {
		log.Fatalf("Failed to make API call with error:\n{%s}", err)
	}

	fmt.Println(res)

}

func main() {
	urlString := "http://35.88.227.145:8080/pipes?pipeType=TRANSFORMER&adminProperties=Yes"

	type PipeObj struct {
			Pipes []struct {
				AdminProperties []struct {
						Name              string `json:"name"`
						SelectListName    string `json:"selectListName"`
						ValidationRules   string `json:"validationRules"`
						Ordering          string `json:"ordering"`
						DefaultValue      string `json:"defaultValue"`
						DataType          string `json:"dataType"`
						HelpInfo          string `json:"helpInfo"`
						DeactivationRules string `json:"deactivationRules"`
						Title             string `json:"title"`
						Value             string `json:"value"`
					} `json:"adminProperties"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Id          string `json:"id"`
				PipeType    string `json:"pipeType"`
			} `json:"pipes"`
	}

	res, err := http.Get(urlString)
	if err != nil {
		log.Fatalf("Failed to execute request with error:\n%s", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	var p1 PipeObj

	err = json.Unmarshal(body, &p1)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range p1.Pipes {
		fmt.Println(p.Name)
		for _, q := range p.AdminProperties {
			fmt.Printf("\t%s\n", q)
		}
	}

	//fmt.Println(p1.Pipes[0].AdminProperties[0].Name)
}
