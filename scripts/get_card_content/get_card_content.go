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
  "github.com/DecentralCardGame/cardobject/keywords"
)

type Content struct {
  Content []byte
}

func getLogger(name string) (*log.Logger) {
  logger := log.Default()
  logger.SetPrefix("\033[1m[" + name + "]\033[0m ")
  return logger
}

func getClient() (cosmosclient.Client, error) {
  config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount("cc", "ccpub")

  return cosmosclient.New(context.Background(), cosmosclient.WithAddressPrefix("cc"), cosmosclient.WithKeyringServiceName("Cardchain"))
}

func broadcastMsg(logger *log.Logger, cosmos cosmosclient.Client, creator string, msg sdktypes.Msg) {
  logger.Println("Message:", msg)

  txResp, err := cosmos.BroadcastTx(creator, msg)
  if err != nil {
    logger.Fatal("Error:", err)
  }
  logger.Println("Response:", txResp)
}

//export get_card_content
func get_card_content(rawContent *C.char) *C.char {
  logger := getLogger("get_card_content")
  contentString := fmt.Sprintf(`{"Content": "%s"}`, C.GoString(rawContent))

  var content Content
  err := json.Unmarshal([]byte(contentString), &content)
  if err != nil {
    logger.Fatal(err)
  }
  return C.CString(string(content.Content))
}

//export make_add_artwork_request
func make_add_artwork_request(creator *C.char, cardId int, image *C.char, fullArt bool) {
  logger := getLogger("make_add_artwork_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address, err := cosmos.Address(C.GoString(creator))
	if err != nil {
		logger.Fatal("Error:", err)
	}

  msg := types.NewMsgAddArtwork(
    address.String(),
		uint64(cardId),
		[]byte(C.GoString(image)),
    fullArt,
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

//export make_save_card_content_request
func make_save_card_content_request(creator *C.char, cardId int, content *C.char, notes *C.char, artist *C.char) {
  logger := getLogger("make_save_card_content_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address, err := cosmos.Address(C.GoString(creator))
	if err != nil {
		logger.Fatal("Error:", err)
	}

  cardobj, err := keywords.Unmarshal([]byte(C.GoString(content)))
  if err != nil {
    logger.Fatal("Error:", err)
  }

  cardbytes, err := json.Marshal(cardobj)
  if err != nil {
    logger.Fatal("Error:", err)
  }

  msg := types.NewMsgSaveCardContent(
    address.String(),
		uint64(cardId),
		cardbytes,
    C.GoString(notes),
    C.GoString(artist),
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

func main() {
}
