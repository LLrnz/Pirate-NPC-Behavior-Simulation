package simulation

import (
	"sync"

	"gitlab.utc.fr/arsaintm/ia04-projet/constants"
)

/*
----------------------------------------
**************** Market ****************
----------------------------------------
*/
// The Market is what is used to calculate ressource prices based on supply and handles transactions

type Market struct {
	sync.RWMutex
	transactions []Transaction

	france  float32
	spain   float32
	england float32
}

// NewMarket initializes a Market struct
func NewMarket() (market *Market) {
	return &Market{}
}

// Getter returns France money
func (m *Market) France() float32 {
	m.RLock()
	defer m.RUnlock()

	return m.france
}

// Getter returns Spain money
func (m *Market) Spain() float32 {
	m.RLock()
	defer m.RUnlock()

	return m.spain
}

// Getter returns England money
func (m *Market) England() float32 {
	m.RLock()
	defer m.RUnlock()

	return m.england
}

// Getter returns England money
func (m *Market) AddMoneyToCountry(ship *Ship, tax float32) {
	m.Lock()
	defer m.Unlock()

	nationality := ship.Nationality()

	if nationality == constants.France {
		m.france += tax
	} else if nationality == constants.Spain {
		m.spain += tax
	} else if nationality == constants.England {
		m.england += tax
	}
}

// Getter returns all transactions
func (m *Market) Transactions() []Transaction {
	m.RLock()
	defer m.RUnlock()

	return m.transactions
}

// Getter returns all transactions
func (m *Market) AddTransactions(transaction Transaction) {
	m.Lock()
	defer m.Unlock()

	m.transactions = append(m.transactions, transaction)

}

// Sell takes an Agent that will sell a ressource for money in exchange
func (m *Market) Sell(ship *Ship, ressource Ressource, step int) {
	//Get agent ressource quantity
	quantitySold := ship.Ressources()[ressource.name]

	// Calculate the price for the quantity of the ressource
	ressourcePrice := m.CalculateRessourcePrice(step, ressource)
	totalPrice := ressourcePrice * float32(quantitySold)
	tax := totalPrice * 0.2

	// Deduct the sold quantity from the agent ressource
	ship.RemoveRessources(ressource.name, quantitySold)

	m.AddMoneyToCountry(ship, tax)

	// Add the money to the Agent account
	ship.AddMoney(totalPrice - tax)

	// Create and record a new transaction
	transaction := NewTransaction(step, ressource, quantitySold, ship, totalPrice)
	m.transactions = append(m.transactions, transaction)

}

// CalculateRessourcePrice returns the price an agent will sell it's ressource
func (m *Market) CalculateRessourcePrice(step int, ressource Ressource) float32 {
	// Get transactions
	transactions := m.Transactions()

	//Lock the market when a transaction is currently happening so the price changes after every transaction
	m.Lock()
	defer m.Unlock()

	// Only get the transactions in the last 365 steps (365 days)
	var quantitySoldBefore int
	for _, transaction := range transactions {
		if transaction.ressource.name == ressource.name && transaction.step >= step-365 {
			quantitySoldBefore += transaction.quantity
		}
	}

	// Calculate the resource price
	price := ressource.minPrice + (ressource.initPrice-ressource.minPrice)/(1+ressource.hyperbolicDecay*float32(quantitySoldBefore))

	// Return the price for the requested quantity
	return price
}

// RessourcePrice returns the price and the difference in percentage with previous ressource price of a ressource
func (m *Market) RessourcePrice(step int, ressource Ressource) (float32, float32) {
	// Get transactions
	transactions := m.Transactions()

	//Lock the market for calculations
	m.Lock()
	defer m.Unlock()

	// Current and previous simulation step for the percent calculation
	currentStep := step
	previousStep := currentStep - 1

	// Get the quantity sold previous step and current step
	var quantitySoldCurrentStep, quantitySoldPreviousStep, quantitySoldBefore int

	for _, transaction := range transactions {
		if transaction.ressource.name == ressource.name {
			if transaction.step >= currentStep-365 {
				quantitySoldBefore += transaction.quantity
			}
			if transaction.step == currentStep {
				quantitySoldCurrentStep += transaction.quantity
			}
			if transaction.step == previousStep {
				quantitySoldPreviousStep += transaction.quantity
			}
		}
	}

	// Calculate the resource price at the current step
	currentPrice := ressource.minPrice + (ressource.initPrice-ressource.minPrice)/(1+ressource.hyperbolicDecay*float32(quantitySoldBefore))

	// Calculate the resource price at the previous step
	previousPrice := ressource.minPrice + (ressource.initPrice-ressource.minPrice)/(1+ressource.hyperbolicDecay*float32(quantitySoldBefore-quantitySoldCurrentStep+quantitySoldPreviousStep))

	// Calculate the percentage difference between previous price and new price
	var percentChange float32
	if previousPrice != 0 {
		percentChange = ((currentPrice - previousPrice) / previousPrice) * 100
	}

	// Return the current price and the percent change
	return currentPrice, percentChange
}
