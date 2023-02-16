package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/rpc"
	"tz/models"
)

func main() {
	c, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	user := models.User{}
	if err := c.Call("Service.Get", "asd@asd", &user); err != nil {
		log.Println(err)
	} else {
		fmt.Println(user)
	}

	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, "http://127.0.0.1:3000/generate-salt", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(body, &user)
	user.Email = "asd@asd"
	user.Password = "12345"
	fmt.Println(user.Salt)

	if err := c.Call("Service.Create", user, &user); err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}
