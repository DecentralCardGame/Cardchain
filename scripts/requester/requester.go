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

func getAddr(logger *log.Logger, cosmos cosmosclient.Client, user string) sdktypes.AccAddress {
  address, err := cosmos.Address(user)
	if err != nil {
		logger.Fatal("Error:", err)
	}
  return address
}

func getClient() (cosmosclient.Client, error) {
  config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount("cc", "ccpub")

  return cosmosclient.New(context.Background(), cosmosclient.WithAddressPrefix("cc"))
}

func broadcastMsg(logger *log.Logger, cosmos cosmosclient.Client, creator string, msg sdktypes.Msg) {
  go logger.Println("Message:", msg)

  txResp, err := cosmos.BroadcastTx(creator, msg)
  if err != nil {
    logger.Fatal("Error:", err)
  }
  go logger.Println("Response:", txResp)
}

func getCardContent(rawContent string) []byte {
  logger := getLogger("get_card_content")
  contentString := fmt.Sprintf(`{"Content": "%s"}`, rawContent)

  var content Content
  err := json.Unmarshal([]byte(contentString), &content)
  if err != nil {
    logger.Fatal(err)
  }
  return content.Content
}

//export make_add_artwork_request
func make_add_artwork_request(creator *C.char, cardId int, image *C.char, fullArt bool) {
  logger := getLogger("make_add_artwork_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address := getAddr(logger, cosmos, C.GoString(creator))

  msg := types.NewMsgAddArtwork(
    address.String(),
		uint64(cardId),
		getCardContent(C.GoString(image)),
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

  address := getAddr(logger, cosmos, C.GoString(creator))
  artistAddr := getAddr(logger, cosmos, C.GoString(artist))

  cardobj, err := keywords.Unmarshal(getCardContent(C.GoString(content)))
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
    artistAddr.String(),
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

//export make_buy_card_scheme_request
func make_buy_card_scheme_request(creator *C.char, price *C.char) {
  logger := getLogger("make_buy_card_scheme_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address := getAddr(logger, cosmos, C.GoString(creator))

  // bid, err := sdktypes.ParseCoinNormalized(C.GoString(price))
  // if err != nil {
  //   logger.Fatal(err)
  // }

  msg := types.NewMsgBuyCardScheme(
    address.String(),
    C.GoString(price),
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

//export make_create_user_request
func make_create_user_request(creator *C.char, alias *C.char) {
  logger := getLogger("make_create_user_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address := getAddr(logger, cosmos, C.GoString(creator))
  useraddr := getAddr(logger, cosmos, C.GoString(alias))

  msg := types.NewMsgCreateuser(
    address.String(),
    useraddr.String(),
    C.GoString(alias),
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

//export make_transfer_card_request
func make_transfer_card_request(creator *C.char, cardId int, receiver *C.char) {
  logger := getLogger("make_transfer_card_request")
  cosmos, err := getClient()
	if err != nil {
		logger.Fatal("Error:", err)
	}

  address := getAddr(logger, cosmos, C.GoString(creator))
  receiverAddr := getAddr(logger, cosmos, C.GoString(receiver))

  msg := types.NewMsgTransferCard(
    address.String(),
    uint64(cardId),
    receiverAddr.String(),
	)

  broadcastMsg(logger, cosmos, C.GoString(creator), msg)
}

func main() {
}
