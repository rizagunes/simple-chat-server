package route

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Router struct {
	rules map[string]Handler
}

type Handler func(*Client, interface{})

func (r *Router)  Handle(msgName string, handler Handler){
	r.rules[msgName] = handler
}

func (r *Router)  FindHandler(msgName string) (Handler, bool){
	 handler, found := r.rules[msgName]
	 return handler, found
}

func NewRouter()  *Router{
	return &Router{rules: make(map[string]Handler),}

}
func New() *echo.Echo {
	e := echo.New()


	router := NewRouter()
	router.Handle("add_channel", addChannel)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")

	e.GET("/ws", ServeWS)
	return e
}

func ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil);
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Logger().Error(err)
		return
	}
	e := &Router{}

	client := NewClient(ws, e.FindHandler())
	go client.Write()
	client.Read()

}
