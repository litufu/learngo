package funding

type Fund struct {
	balance int
}

func NewFund(balance int) *Fund{
	return &Fund{balance:balance}
}

func (f *Fund)Balance() int {
	return f.balance
}

func(f *Fund)WithDraw(amount int){
	f.balance-=amount
}


