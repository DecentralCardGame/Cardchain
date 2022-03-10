package main

import (
  "context"
  "log"
  "encoding/json"
  "fmt"
  "C"

  sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
)

type Content struct {
  Content []byte
}

//export get_card_content
func get_card_content(rawContent *C.char) *C.char {
  contentString := fmt.Sprintf(`{"Content": "%s"}`, C.GoString(rawContent))

  var content Content
  err := json.Unmarshal([]byte(contentString), &content)
  if err != nil {
    panic(err)
  }
  return C.CString(string(content.Content))
}

//export make_add_artwork_request
func make_add_artwork_request(addr *C.char, cardId int, image *C.char, fullArt bool) {
  config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount("cc", "ccpub")

  cosmos, err := cosmosclient.New(context.Background(), cosmosclient.WithAddressPrefix("cc"), cosmosclient.WithKeyringServiceName("Cardchain"))
	if err != nil {
		log.Fatal(err)
	}

  address, err := cosmos.Address(C.GoString(addr))
	if err != nil {
		log.Fatal(err)
	}
  log.Println(address.String())

  msg := types.NewMsgAddArtwork(
    address.String(),
		uint64(cardId),
		[]byte(C.GoString(image)),
    fullArt,
	)
  log.Println(msg)

  txResp, err := cosmos.BroadcastTx(C.GoString(addr), msg)
	if err != nil {
		log.Fatal(err)
	}
  log.Println(txResp)
}

func main() {
}
