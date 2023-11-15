# cardchain
Client library for crowdcontrol cardchain

## Usage
```c#
using CardchainCs.CardchainClient;
using Cosmcs.Crypto.Secp256k1;
using Google.Protobuf;

byte[] hex = new byte[32]
{
	176, 139, 181, 219, 136, 19, 183,
	176, 135, 218, 199, 231, 70, 249,
	129, 238, 212, 107, 207, 147, 217,
	51, 43, 217, 82, 136, 182, 245,
	189, 104, 186, 17
};  // Place byte privateKey here

var privateKey = new PrivateKey(hex);
var accoutAddress = privateKey.PublicKey().AccountId("cc");
Console.Out.WriteLine(accoutAddress);

var ccClient = new CardchainClient("http://localhost:9090", "Testnet3", hex);
var resp = ccClient.SendMsgBuyCardScheme("10000000000000000000", "ucredits").Result;
Console.Out.WriteLine(resp.RawResponse);
if (resp.ResponseMessage != null)
{
	Console.Out.WriteLine(JsonFormatter.Default.Format(resp.ResponseMessage));
}
```

## Development
### Build proto
```bash
git submodule sync
./gen_proto.sh
```