package main

type TXOutput struct {
	Value       int
	ScripPubKey string
}

type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// func NewCoinbaseTX(to, data string) *Transaction {
// 	if data == "" {
// 		data = fmt.Sprintf("Reward to '%s'", to)
// 	}

// 	txin := TXInput{[]byte{}, -1, data}
// 	txout := TXOutput{10, to}
// 	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
// 	tx.SetID()

// 	return &tx
// }
