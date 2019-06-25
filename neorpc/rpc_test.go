package neorpc_test

import (
	"log"
	"testing"

	"github.com/o3labs/neo-rpc/neorpc"
)

var cliNode = "https://explorer.o3node.org"

func TestGetRawTransaction(t *testing.T) {
	client := neorpc.Client(cliNode)
	response, err := client.GetRawTransactionWithDetail("0x57cc08725e0114d8eb95c1799c57a92e57673828284f0276f22f3a712db7e664")
	if err != nil {
		t.Fail()
		return
	}
	log.Printf("%v", response)
}

func TestGetBlock(t *testing.T) {
	client := neorpc.Client(cliNode)
	response, err := client.GetBlock(3623800)
	if err != nil {
		t.Fail()
		return
	}
	log.Printf("%v", response)
}

func TestGetBlocks(t *testing.T) {
	client := neorpc.Client(cliNode)
	from := uint(3623500)
	to := uint(3623540)
	response, err := client.GetBlocks(from, to)
	if err != nil {
		t.Fail()
		return
	}
	log.Printf("%v", response)
}

func TestGetUnclaimed(t *testing.T) {
	client := neorpc.Client(cliNode)
	response, err := client.GetUnclaimed("AcjqHJapAXZbN9LzNkRzzWM1aZj3yuoztu")
	if err != nil {
		log.Printf("err %v", err)
		t.Fail()
		return
	}
	log.Printf("%v", response)
}
