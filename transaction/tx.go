package transaction

type Tx struct {
	From string 		`json:"from,omitempty"`
	Nonce string 		`json:"nonce,omitempty"`
	GasPrice string		`json:"gasPrice,omitempty"`
	Gas string			`json:"gas,omitempty"`
	To string 			`json:"to,omitempty"`
	Value string 		`json:"value,omitempty"`
	Data string 		`json:"data,omitempty"`
	
	V string  			`json:"v,omitempty"`
	R string 			`json:"r,omitempty"`
	S string 			`json:"s,omitempty"`
}