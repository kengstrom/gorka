package gorka

import (
	"log"
	"fmt"
	"encoding/json"
	"time"
	"net/http"
	"html/template"
	"github.com/google/uuid"
)

type Session struct {
	id     string
	UID    string
	Name string
}

type EventMessage struct {
	Service   bool
	Timestamp int64
	Code      string
	Origin    string
	dest      string
	Payload   string
}

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan string

	// New client connections
	newClients chan chan string

	// Closed client connections
	closingClients chan chan string

	// Client connections registry
	clients map[chan string] *Session
}

func NewServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan string, 1),
		newClients:     make(chan chan string),
		closingClients: make(chan chan string),
		clients:        make(map[chan string]*Session),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()
	http.Handle("/events/", broker)
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request ) {
		mainHandler(w, r)
	}))
	http.Handle("/command/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CommandHandler(w, r, broker)
	}))

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", nil))
	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	newId,_ := uuid.NewUUID()
	newUid,_ := uuid.NewUUID()
	// Make sure that the writer supports flushing.
	//
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan string)

	// Signal the broker that we have a new connection
	//broker.newClients <- messageChan
	broker.clients[messageChan] = &Session{
		id:  newId.String(),
		UID: newUid.String(),
	}

	// Send UUID back to client

	go func() {
		msg, _ := json.Marshal(EventMessage{
			Service:   true,
			Timestamp: time.Now().Unix(),
			Code:      "uuid-return",
			Origin:    "auth",
			dest:      newId.String(),
			Payload:   newId.String()})
		messageChan <- string(msg)
		log.Printf("Sent new client its uuid %s", newId.String())
	}()
	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		broker.closingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	notify := rw.(http.CloseNotifier).CloseNotify()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {

		// Write to the ResponseWriter
		// Server Sent Events compatible
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)

		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}

}

func findSession(b *Broker, uuid string) (chan string, *Session) {
	for channel, session := range b.clients {
		if session.id == uuid {
			return channel, session
		}
	}
	return nil, nil
}

func CommandHandler(w http.ResponseWriter, r *http.Request, broker *Broker) {
	v := r.URL.Query()
	command := v.Get("command")
	userid := v.Get("uuid")
	data := v.Get("data")
	//fmt.Println(payload)
	fmt.Println(userid)
	switch command {
	case "?":
		msg, _ := json.Marshal(EventMessage{
			true,
			time.Now().Unix(),
			"message",
			"nil",
			"",
			"These are the available commands:....",
		})
		targetchan, _ := findSession(broker, userid)
		targetchan <- string(msg)
	case "login":
		msg, _ := json.Marshal(EventMessage{
			true,
			time.Now().Unix(),
			"login",
			"nil",
			"",
			"...",
		})
		targetchan, session := findSession(broker, userid)
		session.Name = data
		targetchan <- string(msg)
	case "message":

		_, session := findSession(broker, userid)

		eventString := fmt.Sprintf("%s said %s",session.Name, data)
		msg, _ := json.Marshal(EventMessage{
			true,
			time.Now().Unix(),
			"message",
			"nil",
			"",
			eventString,
		})

		broker.Notifier <- string(msg)
	}
}

func (broker *Broker) listen() {
	for {
		select {
		/*		case s := <-broker.newClients:

					// A new client has connected.
					// Register their message channel
					//broker.clients[s] = true
					log.Printf("Client added. %d registered clients", len(broker.clients))
		*/
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}

func (broker* Broker) BroadcastMessage( message string) {

	msg, _ := json.Marshal(EventMessage{
		true,
		time.Now().Unix(),
		"message",
		"nil",
		"",
		message,
	})

	broker.Notifier <- string(msg)
}
func (broker* Broker) SendMessage( userid string, message string) {
	msg, _ := json.Marshal(EventMessage{
		true,
		time.Now().Unix(),
		"message",
		"nil",
		"",
		message,
	})
	targetchan, _ := findSession(broker, userid)
	targetchan <- string(msg)


}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainTemplate := template.New("sse.html")
	mainTemplate,_ = mainTemplate.ParseFiles("html/templates/sse.html")
	err := mainTemplate.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}
