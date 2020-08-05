package funding

import (
	"fmt"
	"sync"
	"testing"
)

const workers  = 10

func BenchmarkFund(b *testing.B)  {

	if (b.N<workers){
		return
	}


	//f := NewFund(b.N)
	num := int(b.N/10)*10
	fmt.Println(num)
	server := NewFundServer(num)

	var wg sync.WaitGroup

	fundPerWoker := num / workers

	pizzaTime := false
	for i:=0;i<workers;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for j:=0;j<fundPerWoker;j++{
				server.Transact(func(fund *Fund){
					if fund.Balance()<=10{
						pizzaTime=true
						return
					}
					fund.WithDraw(1)
				})
				if pizzaTime{
					break
				}
			}
		}()

	}

	wg.Wait()


	balance:= server.Balance()

	if balance != 10{
		b.Error("Balance wasn't zero:",balance)
	}
}