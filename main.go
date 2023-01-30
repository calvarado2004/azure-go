package main

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"log"
	"os"
)

func main() {
	var (
		token     azcore.AccessToken
		publicKey []byte
		err       error
	)

	if publicKey, err = getPublicKey(); err != nil {
		log.Printf("Error getting public key: %s", err)
		os.Exit(1)
	}

	if token, err = getToken(); err != nil {
		log.Printf("Error getting token: %s", err)
		os.Exit(1)
	}

	if err = launchInstance(token, publicKey); err != nil {
		log.Printf("Error launching instance: %s", err)
		os.Exit(1)
	}

}

func getPublicKey() (publicKey []byte, err error) {

	publicKey, err = os.ReadFile("~/.ssh/id_rsa.pub")
	if err != nil {
		return nil, err
	}

	return publicKey, nil

}

func getToken() (azToken azcore.AccessToken, err error) {

	ctx := context.Background()

	opts := policy.TokenRequestOptions{}

	azCLI, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		log.Printf("Error getting Azure CLI credential: %s", err)
		return azToken, err
	}

	azToken, err = azCLI.GetToken(ctx, opts)
	if err != nil {
		log.Printf("Error getting token: %s", err)
		return azToken, err
	}

	return azToken, nil

}

func launchInstance(token azcore.AccessToken, publicKey []byte) (err error) {

	return nil
}
