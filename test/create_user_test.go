package test

import (
	"context"
	"github.com/nuntiodev/mercury-proto/go_mercury"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"testing"
)

var (
	mercuryClientOne go_mercury.ServiceClient
	mercuryClientTwo go_mercury.ServiceClient
)

func TestMain(m *testing.M) {
	// setup first client
	apiOne := "localhost:9000"
	mercuryConnOne, err := grpc.Dial(apiOne, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	mercuryClientOne = go_mercury.NewServiceClient(mercuryConnOne)
	// setup second client
	apiTwo := "localhost:9001"
	mercuryConnTwo, err := grpc.Dial(apiTwo, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	mercuryClientTwo = go_mercury.NewServiceClient(mercuryConnTwo)
	os.Exit(m.Run())
}

// TestCreateUser creates a user by connecting to one client and fetches the user from another client
func TestCreateUser(t *testing.T) {
	firstName := "Test"
	lastName := "User"
	// create user in client one
	respOne, err := mercuryClientOne.CreateUser(context.Background(), &go_mercury.MercuryRequest{
		User: &go_mercury.User{
			FirstName: firstName,
			LastName:  lastName,
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, respOne)
	// get user using client two
	respTwo, err := mercuryClientTwo.GetUser(context.Background(), &go_mercury.MercuryRequest{
		User: respOne.User,
	})
	assert.NoError(t, err)
	assert.NotNil(t, respTwo)
	assert.Equal(t, firstName, respTwo.User.FirstName)
	assert.Equal(t, lastName, respTwo.User.LastName)
}
