package testingfunc

import (
	"fmt"
	"testing"
)

func TestIntmin(t *testing.T) {
	ans := Intmin(2,-2)
	if ans != -2{
		t.Errorf("Intmin(2,-2)=%d,want -2",ans)
	}
}

func TestIntminTableDriven(t *testing.T) {
	var tests = []struct{
		a,b int
		ans int
	}{
		{1,2,1},
		{4,5,4},
		{9,2,2},
	}

	for _,test := range tests{
		testname := fmt.Sprintf("%d-%d",test.a,test.b)
		t.Run(testname,func(t *testing.T){
			ans := Intmin(test.a,test.b)
			if ans != test.ans{
				t.Errorf("Intmin(%d,%d)=%d want%d",test.a,test.b,ans,test.ans)
			}
		})
	}
}