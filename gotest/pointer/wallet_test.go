package pointer

import (
	"errors"
	"fmt"
	"testing"
)

func TestWAllet(t *testing.T) {
	// refactor a bit, to reduce duplications
	assertBalance := func(t testing.TB, wlt Wallet, wnt Bitcoin) {
		t.Helper()
		got := wlt.Balance()
		if got != wnt {
			t.Errorf("got %s want %s", got, wnt)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Error("error occurred", err)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)} // start with a balance of 30 BTC

		wallet.Withdrawl(Bitcoin(10))
		// assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(20))
	})

	t.Run("Insufficient Balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)} // start with a balance of 30 BTC
		err := wallet.Withdrawl(Bitcoin(100))

		assertError(t, err)
		assertBalance(t, wallet, Bitcoin(30))
	})
}

// creating custom types
type Bitcoin int // to make wallet field type more descriptive

type Stringer interface { // let we define our string for the type Bitcoin, since go works on hierarchy of modules,
	//naming issues with std pkgs will not happen.
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdrawl(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("Insufficient balance")
	}
	w.balance -= amount
	return nil
}
