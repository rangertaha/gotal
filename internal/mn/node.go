package mn


import (
	"time"
)

// Node represents a node in a Markov chain graph
type Node struct {
	// Value stored in this node
	Value interface{}
	
	// Map of transitions to other nodes with their probabilities
	Transitions map[*Node]float64
	
	// Count of times this node has been visited
	Visits int
	
	// Timestamp of last visit
	LastVisit time.Time
	
	// Additional metadata
	Metadata map[string]interface{}
}

// NewNode creates a new Markov chain node
func NewNode(value interface{}) *Node {
	return &Node{
		Value:       value,
		Transitions: make(map[*Node]float64),
		Visits:      0,
		LastVisit:   time.Time{},
		Metadata:    make(map[string]interface{}),
	}
}

// AddTransition adds or updates a transition to another node with given probability
func (n *Node) AddTransition(to *Node, probability float64) {
	n.Transitions[to] = probability
}

// RemoveTransition removes a transition to the specified node
func (n *Node) RemoveTransition(to *Node) {
	delete(n.Transitions, to)
}

// Visit increments the visit count and updates last visit time
func (n *Node) Visit() {
	n.Visits++
	n.LastVisit = time.Now()
}

// GetTransitionProbability returns the probability of transitioning to the given node
func (n *Node) GetTransitionProbability(to *Node) float64 {
	if prob, exists := n.Transitions[to]; exists {
		return prob
	}
	return 0.0
}

// GetProbabilities returns a map of all transition probabilities
func (n *Node) GetProbabilities() map[*Node]float64 {
	probs := make(map[*Node]float64)
	for node, prob := range n.Transitions {
		probs[node] = prob
	}
	return probs
}
