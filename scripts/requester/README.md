# requester
Get's a clear string card content from a as string marshaled byte array. Has to be installed first to be able to run import_old_genesis.py

## Install
```
go build -buildmode=c-shared -o ./requester.so ./requester.go
```

## Usage
Import `requester`.
Note that "Cooler" Typ needs to have enought funds
```
Cardchaind tx bank send alice $(Cardchaind keys show "Cooler Typ" --address) 2000000000000000000000000000000000000000000000000000000ucredits
```
