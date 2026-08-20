package main

import "fmt"

type Account struct {
	name string
}

func (a *Account) Check() error {
	fmt.Printf("[Account] Проверка аккаунта '%s'...\n", a.name)
	return nil
}

type SecurityCode struct {
	code int
}

func (s *SecurityCode) Check(code int) error {
	if s.code != code {
		return fmt.Errorf("[SecurityCode] Неверный PIN-код")
	}
	fmt.Println("[SecurityCode] PIN-код успешно подтвержден.")
	return nil
}

type Wallet struct {
	balance int
}

func (w *Wallet) Credit(amount int) {
	w.balance += amount
	fmt.Printf("[Wallet] Баланс пополнен на %d. Текущий баланс: %d\n", amount, w.balance)
}

func (w *Wallet) Debit(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("[Wallet] Недостаточно средств на балансе (%d из %d)", w.balance, amount)
	}
	w.balance -= amount
	fmt.Printf("[Wallet] Списано %d. Текущий баланс: %d\n", amount, w.balance)
	return nil
}

type Notification struct{}

func (n *Notification) SendSuccess(operation string, amount int) {
	fmt.Printf("[Notification] SMS: Операция '%s' на сумму %d выполнена успешно!\n", operation, amount)
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notif        *Notification
}

func NewWalletFacade(accountName string, pinCode int) *WalletFacade {
	return &WalletFacade{
		account:      &Account{name: accountName},
		wallet:       &Wallet{balance: 0},
		securityCode: &SecurityCode{code: pinCode},
		notif:        &Notification{},
	}
}

func (w *WalletFacade) DepositMoney(pinCode int, amount int) error {
	fmt.Println("---> Запрос: Пополнение баланса")

	if err := w.securityCode.Check(pinCode); err != nil {
		return err
	}

	if err := w.account.Check(); err != nil {
		return err
	}

	w.wallet.Credit(amount)

	w.notif.SendSuccess("Пополнение", amount)

	return nil
}

func (w *WalletFacade) DeductMoney(pinCode int, amount int) error {
	fmt.Println("---> Запрос: Списание средств")

	if err := w.securityCode.Check(pinCode); err != nil {
		return err
	}

	if err := w.account.Check(); err != nil {
		return err
	}

	if err := w.wallet.Debit(amount); err != nil {
		return err
	}

	w.notif.SendSuccess("Списание", amount)

	return nil
}

func main() {
	walletFacade := NewWalletFacade("Иван Иванов", 1234)

	err := walletFacade.DepositMoney(0000, 1000)
	if err != nil {
		fmt.Printf("Ошибка: %v\n\n", err)
	}

	err = walletFacade.DepositMoney(1234, 1000)
	if err != nil {
		fmt.Printf("Ошибка: %v\n\n", err)
	}

	err = walletFacade.DeductMoney(1234, 1500)
	if err != nil {
		fmt.Printf("Ошибка: %v\n\n", err)
	}

	err = walletFacade.DeductMoney(1234, 300)
	if err != nil {
		fmt.Printf("Ошибка: %v\n\n", err)
	}
}