package main

import (
	"kin-starter-go/kinsdk"
	"context"
	"fmt"
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

	fmt.Printf("🔑 Public Key Alice    %s", privateKeyAlice.Public().Base58())
	for _, tokenAccount := range tokenAccountsAlice {
		fmt.Printf("🗝  Token Account Alice %s", tokenAccount.Base58())
	}

	privateKeyBob, err := kinsdk.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	tokenAccountsBob, err := k.CreateAccount(context.Background(), privateKeyBob)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🔑 Public Key Bob    %s", privateKeyBob.Public().Base58())
	for _, tokenAccount := range tokenAccountsBob {
		fmt.Printf("🗝  Token Account Bob %s", tokenAccount.Base58())
	}

	// Helper method to sleep a bit, then print balance of Alice and Bob
	sleepAndPrintBalances := func() {
		fmt.Printf("😴 Sleeping for a bit...")
		time.Sleep(15 * time.Second)
		balanceAlice, err:= k.GetBalance(context.Background(), privateKeyAlice.Public())
		fmt.Printf("👛 Balance for Alice:  %x Kin, err: %v", balanceAlice, err)
		balanceBob, err:= k.GetBalance(context.Background(), privateKeyBob.Public())
		fmt.Printf("👛 Balance for Bob:    %x Kin, err: %v", balanceBob, err)
	}

	sleepAndPrintBalances()

	fmt.Printf("🙏 Request Airdrop for Alice")
	k.RequestAirdrop(context.Background(), tokenAccountsAlice[0], "10")

	fmt.Printf("🙏 Request Airdrop for Bob")
	k.RequestAirdrop(context.Background(), tokenAccountsBob[0], "10")

	sleepAndPrintBalances()

	fmt.Printf("💸 Submit P2P Payment from Alice to Bob")
	k.SubmitP2P(context.Background(), privateKeyAlice, privateKeyBob.Public(), "2", "My demo payment")

	sleepAndPrintBalances()

	fmt.Printf("✅ Done!")
}