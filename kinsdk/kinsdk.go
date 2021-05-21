package kinsdk

import (
	"context"
	"github.com/kinecosystem/agora-common/kin"
	"github.com/kinecosystem/kin-go/client"
	"log"
)

type Kin struct {
	kinClient client.Client
}

func GenerateKey() (kin.PrivateKey, error) {
	return kin.NewPrivateKey()
}

func New(env client.Environment) (*Kin, error) {
	c, err := client.New(env, client.WithAppIndex(1), client.WithMaxRetries(0))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Kin{
		kinClient: c,
	}, nil
}

func (k *Kin) CreateAccount(ctx context.Context, key kin.PrivateKey) ([]kin.PublicKey, error) {
	err := k.kinClient.CreateAccount(ctx, key)
	if err != nil {
		log.Fatal(err)
	}

	return k.kinClient.ResolveTokenAccounts(ctx, key.Public())
}

func (k *Kin) GetBalance(ctx context.Context, account kin.PublicKey) (int64, error) {
	return k.kinClient.GetBalance(ctx, account)
}

func (k *Kin) RequestAirdrop(ctx context.Context, publicKey kin.PublicKey, amount string) ([]byte, error) {
	Quarks, err := kin.ToQuarks(amount)
	if err != nil {
		log.Fatal(err)
	}
	return k.kinClient.RequestAirdrop(ctx, publicKey, uint64(Quarks))
}

func (k *Kin) SubmitPayment(
	ctx context.Context,
	sender kin.PrivateKey,
	destination kin.PublicKey,
	amount string,
	transactionType kin.TransactionType,
	memo string) ([]byte, error) {

	return k.kinClient.SubmitPayment(ctx, client.Payment{
		Sender:      sender,
		Destination: destination,
		Type:        transactionType,
		Quarks:      kin.MustToQuarks(amount),
		Memo:        memo,
	})
}

func (k *Kin) SubmitEarn(
	ctx context.Context,
	sender kin.PrivateKey,
	destination kin.PublicKey,
	amount string,
	memo string) ([]byte, error) {
	return k.SubmitPayment(ctx, sender, destination, amount, kin.TransactionTypeEarn, memo)
}

func (k *Kin) SubmitSpend(
	ctx context.Context,
	sender kin.PrivateKey,
	destination kin.PublicKey,
	amount string,
	memo string) ([]byte, error) {
	return k.SubmitPayment(ctx, sender, destination, amount, kin.TransactionTypeSpend, memo)
}

func (k *Kin) SubmitP2P(
	ctx context.Context,
	sender kin.PrivateKey,
	destination kin.PublicKey,
	amount string,
	memo string) ([]byte, error) {
	return k.SubmitPayment(ctx, sender, destination, amount, kin.TransactionTypeP2P, memo)
}