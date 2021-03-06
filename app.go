package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/optimizely/go-sdk/pkg/client"
	"github.com/optimizely/go-sdk/pkg/entities"
	"github.com/optimizely/go-sdk/pkg/logging"
)

const optimizeSdkKey = "XWrur15ApwRCYRASkqUmA"

func randString() string {
	return strconv.Itoa(int(rand.Float64() * 100000))
}

func main() {

	//rand.Seed(time.Now().UTC().UnixNano()) - without seeding the random generator
	//we get the same set of emails every time so we can nicelty see how feature values
	//change for the same users depending 
	logging.SetLogLevel(logging.LogLevelError)

	optimizelyFactory := &client.OptimizelyFactory{
		SDKKey: optimizeSdkKey,
	}

	// Instantiates a client that syncs the datafile in the background
	// this should be done on handler level in the lambda, so that the client lives and is not
	// created for every request! 
	optimizelyClient, err := optimizelyFactory.Client()

	if err != nil {
		panic(err)
	}

	for i := 0; i <= 10; i++ {
		user := entities.UserContext{
			ID: "michal+" + randString() + "@zavamed.com",
		}

		isEnabled, err := optimizelyClient.IsFeatureEnabled("zava_de_email_verification", user)

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s: %t\n", user.ID, isEnabled)
	}

}
