package app

type ServiceTypes map[string]string

func (services *ServiceTypes) InitValidServices() {
	// TODO: externalise this into a config file
	services = &ServiceTypes{
		"PubSub":    "pubsub.googleapis.com",
		"Container": "run.googleapis.com",
	}
}

// TODO: Implement a case insensitive search for keys
