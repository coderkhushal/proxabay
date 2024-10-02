package service

import (
	"encoding/json"
	"fmt"
	"os"
)

type Cache struct {
	Origin  string
	Port    string
	Headers []byte

	Body []byte
}

func CreateNewCache(origin string, port string, headers []byte, body []byte) error {
	c := Cache{
		Origin:  origin,
		Port:    port,
		Headers: headers,
		Body:    body,
	}

	_, err := os.Stat("proxycache.json")
	if os.IsNotExist(err) {
		os.Create("proxycache.json")
		emptycachejson, _ := json.Marshal(make([]Cache, 0))
		os.WriteFile("proxycache.json", emptycachejson, 0644)
		fmt.Println("created json file for cache")

	}
	file, err := os.ReadFile("proxycache.json")
	if err != nil {
		fmt.Println(err)
	}
	var filecontent []Cache
	json.Unmarshal(file, &filecontent)

	updatedfilecontent, err := json.Marshal(append(filecontent, c))

	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("proxycache.json", updatedfilecontent, 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GetCacheForProxy(origin string, port string) (Cache, error) {
	_, err := os.Stat("proxycache.json")
	if err != nil {
		fmt.Println("some Error occured while Accessing cache file ")
		return Cache{}, err
	}
	file, err := os.ReadFile("proxycache.json")

	if err != nil {
		fmt.Println("some Error occured while Reading cache file ")
		return Cache{}, err
	}
	var filecontent []Cache
	err = json.Unmarshal(file, &filecontent)

	if err != nil {
		fmt.Println("some Error occured while reading cache file ")
		os.Remove("proxycache.json")
		return Cache{}, err
	}

	for _, value := range filecontent {
		if (value.Origin == origin) && (value.Port == port) {
			return value, nil
		}
	}
	return Cache{}, nil
}

func ClearCache() {

	_, err := os.Stat("proxycache.json")
	if err != nil {
		fmt.Println("some Error occured while Accessing cache file ")
		return
	}

	emptycachejson, _ := json.Marshal(make([]Cache, 0))
	err = os.WriteFile("proxycache.json", emptycachejson, 0644)
	if err != nil {
		fmt.Println("Error while clearing cache")

	}
	fmt.Println("Cache cleared")

}
