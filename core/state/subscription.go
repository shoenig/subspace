// Author hoenig

package state

// A CreateSubscription represents the creation request of a new Subscription that
// agents can subscribe to and recieve generations of torrents from.
type CreateSubscription struct {
	Name  string
	Owner string
}

func ValidateCreateSubscription(c CreateSubscription) error {
}
