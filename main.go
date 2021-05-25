package main

import (
	"context"
	"fmt"
	"github.com/kintegrate/kin-starter-go/kinsdk"
	"log"
	"time"

	"github.com/kinecosystem/kin-go/client"
)

func main() {
	// Set up Kin client
	k, err := kinsdk.New(client.EnvironmentTest)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare tokens for Alice and Bob
	privateKeyAlice, err := kinsdk.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	tokenAccountsAlice, err := k.CreateAccount(context.Background(), privateKeyAlice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🔑 Public Key Alice    %s\n", privateKeyAlice.Public().Base58())
	for _, tokenAccount := range tokenAccountsAlice {
		fmt.Printf("🗝  Token Account Alice %s\n", tokenAccount.Base58())
	}

	privateKeyBob, err := kinsdk.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	tokenAccountsBob, err := k.CreateAccount(context.Background(), privateKeyBob)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🔑 Public Key Bob    %s\n", privateKeyBob.Public().Base58())
	for _, tokenAccount := range tokenAccountsBob {
		fmt.Printf("🗝  Token Account Bob %s\n", tokenAccount.Base58())
	}

	// Helper method to sleep a bit, then print balance of Alice and Bob
	sleepAndPrintBalances := func() {
		fmt.Printf("😴 Sleeping for a bit...\n")
		time.Sleep(15 * time.Second)
		balanceAlice, _:= k.GetBalance(context.Background(), privateKeyAlice.Public())
		fmt.Printf("👛 Balance for Alice:  %d Kin\n", balanceAlice)
		balanceBob, _:= k.GetBalance(context.Background(), privateKeyBob.Public())
		fmt.Printf("👛 Balance for Bob:    %d Kin\n", balanceBob)
	}

	sleepAndPrintBalances()

	fmt.Printf("🙏 Request Airdrop for Alice\n")
	k.RequestAirdrop(context.Background(), tokenAccountsAlice[0], "10")

	fmt.Printf("🙏 Request Airdrop for Bob\n")
	k.RequestAirdrop(context.Background(), tokenAccountsBob[0], "10")

	sleepAndPrintBalances()

	fmt.Printf("💸 Submit P2P Payment from Alice to Bob\n")
	k.SubmitP2P(context.Background(), privateKeyAlice, privateKeyBob.Public(), "2", "My demo payment")

	sleepAndPrintBalances()

	fmt.Printf("✅ Done!")
}