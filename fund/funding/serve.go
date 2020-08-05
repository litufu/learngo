package funding

type FundServer struct {
	commands chan TransactionComand
	fund Fund
}

type Transactor func(fund *Fund)

type TransactionComand struct {
	Transactor Transactor
	Done chan bool
}


func NewFundServer(initialNumer int) *FundServer{
	server := &FundServer{
		commands: make(chan TransactionComand),
		fund:     *NewFund(initialNumer),
	}
	go server.loop()
	return server
}


func (s *FundServer)Balance() int{
	var balance int
	s.Transact(func(fund *Fund){
		balance = fund.Balance()
	})
	 return balance
}

func (s *FundServer)WithDraw(amount int){
	s.Transact(func(fund *Fund){
		fund.WithDraw(amount)
	})
}

func (s *FundServer)Transact(transactor Transactor){
	command := TransactionComand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands<-command
	<-command.Done
}

func (s *FundServer)loop(){
	for transction := range s.commands{
		transction.Transactor(&s.fund)
		transction.Done<-true
	}
}
