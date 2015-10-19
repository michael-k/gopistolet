package main

import (
	"fmt"
	"log"

	"github.com/gopistolet/gopistolet/mta"
)

func mail(state *mta.State) {
	log.Printf("From: %s\n", state.From.Address)
	log.Printf("To: ")
	for i, to := range state.To {
		log.Printf("%s", to.Address)
		if i != len(state.To)-1 {
			log.Printf(",")
		}
	}
	log.Printf("\nCONTENT_START:\n")
	log.Printf("%s\n", string(state.Data))
	log.Printf("CONTENT_END\n\n\n\n")
}

func main() {
	//sigc := make(chan os.Signal, 1)
	//signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	fmt.Println("GoPistolet at your service!")

	c := mta.Config{
		Hostname: "localhost",
		Port:     2525,
	}

	mta := mta.NewDefault(c, mta.HandlerFunc(mail))
	go func() {
		//<-sigc
		mta.Stop()
	}()
	err := mta.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
