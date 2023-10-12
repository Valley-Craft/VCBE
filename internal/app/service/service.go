package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorcon/rcon"
	"github.com/gtuk/discordwebhook"
	"log"
	"net/http"
	"os"
	"strconv"
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

type Donate struct {
	PubId       string `json:"pubId"`
	ClientName  string `json:"clientName"`
	Message     string `json:"message"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Source      string `json:"source"`
	Goal        string `json:"goal"`
	IsPublished bool   `json:"isPublished"`
	CreatedAt   string `json:"createdAt"`
}

func sendWebHook(nickname, name, age, wwd, rules, status string) {
	var username = "Анкета"
	data := Form{
		Nickname: nickname,
		Name:     name,
		Age:      age,
		Wwd:      wwd,
		Rules:    rules,
	}
	content := fmt.Sprintf("Статус: %s\nНік: %s\nІм'я: %s\nРік: %s\nЩо буду робити: %s\nПравила: %s",
		status, data.Nickname, data.Name, data.Age, data.Wwd, data.Rules)
	var url, _ = os.LookupEnv("WEB_HOOK_URL")

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) Players() ([]JSONPlayer, error) {
	req, err := http.NewRequest("GET", "http://178.63.27.177:4567/v1/players", nil)
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

		conn, err := rcon.Dial("178.63.27.177:25575", rconPass)
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

		sendWebHook(person.Nickname, person.Name, person.Age, person.Wwd, person.Rules, "Прийнятий")

		return true
	} else {
		sendWebHook(person.Nickname, person.Name, person.Age, person.Wwd, person.Rules, "Не прийнятий")
	}

	return false
}

func (s *Service) Donate(body string) bool {
	var donate Donate
	err := json.Unmarshal([]byte(body), &donate)

	if err != nil {
		fmt.Println("error:", err)
	}

	amount, err := strconv.Atoi(donate.Amount)

	if amount >= 50 {

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

		_, err = conn.Execute(fmt.Sprintf("luckperms user %s parent add donate", donate.ClientName))
		//_, err = conn.Execute(fmt.Sprintf("say %s", donate.ClientName))

		if err != nil {
			panic(err)
		}

		return true
	}
	return true
}
