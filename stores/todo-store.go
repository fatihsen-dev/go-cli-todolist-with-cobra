package stores

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type TodoT struct {
	Id     int
	Text   string
	Status string
	Date   string
}

var Todos []TodoT

func FetchTodos() {
	data, err := ioutil.ReadFile("./database/todos.json")
	if err != nil {
		fmt.Println("Database dosyası bulunamadı")
		return
	}

	err = json.Unmarshal(data, &Todos)
	if err != nil {
		fmt.Println("todos veri dönüşümü başarısız")
		return
	}
}

func SyncJson() {
	data, marshalErr := json.Marshal(Todos)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	ioutil.WriteFile("./database/todos.json", data, 0644)
}
