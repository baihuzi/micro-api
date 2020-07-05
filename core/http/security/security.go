package security

type Security struct {
	//policies
	//firewalls
	Firewalls map[string]FirewallInterface
	AccessTokenGenerator map[string]AccessTokenGeneratorInterface
}
