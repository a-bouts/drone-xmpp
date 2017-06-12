package main

import (
	"crypto/tls"
        "fmt"
	"github.com/mattn/go-xmpp"
	"log"
	"os"
	"strings"
)

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func main() {

	message := fmt.Sprintf("[%s] <%s> (%s)『%s』by %s", os.Getenv("DRONE_BUILD_STATUS"), os.Getenv("DRONE_BUILD_LINK"), os.Getenv("DRONE_COMMIT_BRANCH"), os.Getenv("DRONE_COMMIT_MESSAGE"), os.Getenv("DRONE_COMMIT_AUTHOR"));

	log.Print("Sending message " + message);


	var jid = os.Getenv("XMPP_JID")
	var password = os.Getenv("XMPP_PASSWORD")
	var to = os.Getenv("XMPP_TO")


	var host string
	host, ok := os.LookupEnv("XMPP_HOST")
	if !ok {
		host = strings.Split(jid, "@")[1]
	}

	var talk *xmpp.Client
	var err error
	options := xmpp.Options {
		Host:          host,
		User:          jid,
		Password:      password,
		NoTLS:         true,
		StartTLS:      true,
        	TLSConfig: &tls.Config{
            		ServerName: serverName(host),
            		InsecureSkipVerify: false,
        	},
		Debug:         false,
		Session:       false,
	}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	_, err = talk.Send(xmpp.Chat{Remote: to, Type: "chat", Text: message})
	if err != nil {
		log.Fatal(err)
	}
}
