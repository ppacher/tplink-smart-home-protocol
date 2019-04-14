package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	tpshp "github.com/ppacher/tplink-smart-home-protocol"
)

var deviceIP = flag.String("d", "", "The IP address of the device")
var devicePort = flag.Int("p", 9999, "The port the TP-Link device is listening. Defaults to 9999")

func main() {
	flag.Parse()

	command := flag.Arg(0)
	if command == "" {
		log.Fatalf("Missing command name")
	}

	parts := strings.Split(command, ".")
	if len(parts) < 2 {
		log.Fatal("Command must have the format <namespace>.<command>")
	}

	var payload interface{}
	if flag.Arg(1) != "" {
		if err := json.Unmarshal([]byte(flag.Arg(1)), &payload); err != nil {
			log.Fatal(err)
		}
	}

	ns := parts[0]
	name := strings.Join(parts[1:], ".")

	req := tpshp.NewRequest()

	var res interface{}
	req.AddCommand(ns, name, payload, &res)

	cli := tpshp.NewWithPort(*deviceIP, uint16(*devicePort))
	if err := cli.Call(context.Background(), req); err != nil {
		log.Fatal(err)
	}

	responseBlob, _ := json.Marshal(req.Responses())
	fmt.Println(string(responseBlob))
}
