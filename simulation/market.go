package simulation

import (
	"sync"
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
}

// NewMarket initializes a Market struct
func NewMarket() (market *Market) {
	return &Market{}
}

// Getter returns all transactions
func (m *Market) Transactions() []Transaction {
	return m.transactions
}

// Sell takes an Agent that will sell a ressource for money in exchange
func (m *Market) Sell(simu *Simulation, ship *Ship, ressource Ressource) error {
	//Lock the market when a transaction is currently happening so the price changes after every transaction
	m.Lock()
	defer m.Unlock()

	//Get agent ressource quantity [[REPLACE WITH THE GETTER]]
	quantitySold := ship.ressources[ressource.name]

	// Calculate the price for the quantity of the ressource
	ressourcePrice := m.CalculateRessourcePrice(simu, ressource)
	totalPrice := ressourcePrice * float32(quantitySold)

	// Deduct the sold quantity from the agent's resources [[REPLACE WITH THE SETTER]]
	ship.ressources[ressource.name] = 0

	// Add the money to the agent's account [[REPLACE WITH THE SETTER]]
	ship.money += totalPrice

	// Create and record a new transaction
	transaction := NewTransaction(simu.parameters.step, ressource, quantitySold, ship, totalPrice)
	m.transactions = append(m.transactions, transaction)

	return nil
}

// CalculateRessourcePrice returns the price an agent will sell it's ressource
func (m *Market) CalculateRessourcePrice(simu *Simulation, ressource Ressource) float32 {
	// Get transactions
	transactions := m.Transactions()

	// Current simulation step
	currentStep := simu.parameters.step

	// Only get the transactions in the last 31 steps (31 days)
	var quantitySoldBefore int
	for _, transaction := range transactions {
		if transaction.ressource.name == ressource.name && transaction.step >= currentStep-31 {
			quantitySoldBefore += transaction.quantity
		}
	}

	// Calculate the resource price
	price := ressource.minPrice + (ressource.initPrice-ressource.minPrice)/(1+ressource.hyperbolicDecay*float32(quantitySoldBefore))

	// Return the price for the requested quantity
	return price
}

// RessourcePrice returns the price and the difference in percentage with previous ressource price of a ressource
func (m *Market) RessourcePrice(simu *Simulation, ressource Ressource) (float32, float32) {
	// Get transactions
	transactions := m.Transactions()

	// Current and previous simulation step for the percent calculation
	currentStep := simu.parameters.step
	previousStep := currentStep - 1

	// Get the quantity sold previous step and current step
	var quantitySoldCurrentStep, quantitySoldPreviousStep, quantitySoldBefore int

	for _, transaction := range transactions {
		if transaction.ressource.name == ressource.name {
			if transaction.step >= currentStep-31 {
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
