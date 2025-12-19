package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/avinxshKD/Sahara/internal/p2p"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p/core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	peerAddr := flag.String("peer-address", "", "peer address")
	port := flag.String("port", "9000", "port to listen on")

	flag.Parse()
	host, CONNECTION_STRING, err := p2p.CreateHost(*port)

	fmt.Println("Connection string: ", CONNECTION_STRING)

	if err != nil {
		log.Fatal("Error creating host node")
		return
	}

	// If we received a peer address, we should connect to it.
	if *peerAddr != "" {
		// Parse the multiaddr string.
		peerMA, err := ma.NewMultiaddr(*peerAddr)
		if err != nil {
			panic(err)
		}
		peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMA)
		if err != nil {
			panic(err)
		}

		// Connect to the node at the given address.
		if err := host.Connect(context.Background(), *peerAddrInfo); err != nil {
			panic(err)
		}
		log.Println("Connected to", peerAddrInfo.String())
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	<-sigCh

}
