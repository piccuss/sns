package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Message struct {
	Text, Desp string
}

const (
	pushUrl = "https://sc.ftqq.com/SCU150195T2c5855e7da9b91fef1774a40de01c2cd5fffb1ee24c03.send"
)

// Send message to remote server
func (m Message) Push() {
	form := make(url.Values)
	form.Set("text", m.Text)
	form.Set("desp", m.Desp)
	resp, err := http.PostForm(pushUrl, form)
	if err != nil {
		log.Println("message push failed", err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("message push resp: %s", body)
}

func NewMessage(stocks []Stock) Message {
	var text, desp string
	for i := 0; i < len(stocks); i++ {
		text += fmt.Sprintf("%s:%.2f%% ", stocks[i].Name, stocks[i].ChgR)
		desp += fmt.Sprintf("* %s(%s)\t%.3f\t%.2f%%(%.3f)\n\r", stocks[i].Name, stocks[i].Code, stocks[i].Price, stocks[i].ChgR, stocks[i].Chg)
	}
	return Message{Text: text, Desp: desp}
}
