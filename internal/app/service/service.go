package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorcon/rcon"
	"log"
	"net/http"
	"os"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

type JSONPlayer struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayname"`
}

type Player struct {
	UUID        string `json:"uuid"`
	DisplayName string `json:"displayname"`
}

type Form struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Wwd      string `json:"wwd"`
	Rules    string `json:"rules"`
}

func (s *Service) Players() ([]JSONPlayer, error) {
	req, err := http.NewRequest("GET", "http://137.74.7.233:4567/v1/players", nil)
	if err != nil {
		return nil, err
	}

	key, _ := os.LookupEnv("SERVER_KEY")
	req.Header.Set("key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var players []Player
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&players)
	if err != nil {
		return nil, err
	}

	var jsonPlayers []JSONPlayer
	for _, player := range players {
		jsonPlayers = append(jsonPlayers, JSONPlayer{
			UUID:        player.UUID,
			DisplayName: player.DisplayName,
		})
	}

	fmt.Println(jsonPlayers)

	return jsonPlayers, nil
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
