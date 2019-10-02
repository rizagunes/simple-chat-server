package main

import (
	"fmt"
	"net/http"
	"simple-socket-server/route"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mitchellh/mapstructure"

)



type Message struct {
	Name string `json:"name"`
	Data interface{} `json:"data"`
}

type Channel struct {
	Id string `json:"id"`
	Name string `json:"name"`
}


var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {

		var inMessage Message
		var outMessage Message
		if err := ws.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("%#v\n", inMessage)

		switch inMessage.Name {
		case "add_channel":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := ws.WriteJSON(outMessage); err != nil{
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			go subscribeChannel(ws)
		}


	}
	return nil
}

func subscribeChannel(ws *websocket.Conn) {

	for{
		time.Sleep(time.Second * 1 )
		message := Message{"add_channel",
			Channel{"1", "Software Support"}}
		ws.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}

func addChannel(data interface{}) (error) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)

	if err != nil {
		return err
	}
	channel.Id = "1"
	fmt.Println("added channel")
	return nil

}

func main() {

	e := route.New()

	e.Logger.Fatal(e.Start(":1323"))
}

