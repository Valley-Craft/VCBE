package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorcon/rcon"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

type Form struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Wwd      string `json:"wwd"`
	Rules    string `json:"rules"`
}

func (s *Service) Players() string {

	req, err := http.NewRequest("GET", "http://137.74.7.233:4567/v1/players", nil)

	if err != nil {
		panic(err)
	}

	key, _ := os.LookupEnv("SERVER_KEY")

	req.Header.Set("key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	return string(body)
}

func (s *Service) Form(body string) bool {

	var person Form
	err := json.Unmarshal([]byte(body), &person)
	if err != nil {
		fmt.Println("error:", err)
	}

	if person.Rules == "123" && len(person.Nickname) <= 16 {
		fmt.Println("User added")

		rconPass, _ := os.LookupEnv("RCON_PASSWORD")

		conn, err := rcon.Dial("137.74.7.233:25575", rconPass)
		if err != nil {
			panic(err)
		}
		defer func(conn *rcon.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)

		_, err = conn.Execute(fmt.Sprintf("easywl add %s", person.Nickname))
		if err != nil {
			panic(err)
		}

		return true
	}

	return false
}
