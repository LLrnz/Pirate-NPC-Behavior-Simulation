package simulation

/*
----------------------------------------
************* Transaction **************
----------------------------------------
*/
// A transaction is created every time a ship sell ressources

type Transaction struct {
	step        int
	ressource   Ressource
	quantity    int
	agentSeller *Ship
	price       float32
}

// NewTransaction initializes a Transaction
func NewTransaction(step int, ressource Ressource, quantity int, agentSeller *Ship, price float32) (transaction Transaction) {
	return Transaction{
		step:        step,
		ressource:   ressource,
		quantity:    quantity,
		agentSeller: agentSeller,
		price:       price,
	}
}
