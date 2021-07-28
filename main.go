package main
import (
	"fmt"
	"math"
)

const MAX = 12345
var p [MAX]int
var q [MAX]int
var m [MAX]int
var pi int
var k int



func main(){
	
	var n int 
	sieve(MAX);
	for {	
		value ,err := fmt.Scanf("%d", &n)
		if err != nil {
			//panic(err)
		}

		if value == 1 {
			// é uma character não numerico
		}

		if n == 0 {
			break
		}
		fact(int(math.Abs(float64(n))))
		 menosum  := ""
		if(n < 0){
			menosum = " -1"
		}

		fmt.Printf("%d =%s",n,menosum)
		first := n > 0
		
		for i := 0; i < k; i++ {
			for j := 0; j < m[i]; j++ {
				firstStr := " "
				if(!first){
					firstStr = " x "
				}	
				fmt.Printf("%s%d", firstStr ,q[i])
				first = false
			}
		}
		fmt.Println()
	}

	

}


func sieve(n int){
	p[0] = 2; pi = 1;
	for i := 3; i <= n; i++ {
		prime := true 
		r := int(math.Sqrt(float64(i)))
		for j := 0; j < pi && p[j] <= r && prime &&  pi < len(p); j++ {
			
			prime = i % p[j] != 0;

			if(prime){
				p[pi] = i
				pi++
			}

		}

	}
}

func fact(n int){
	k = 0
	for i := 0; i < pi && p[i] * p[i] <= n; i++ {
		if n%p[i] == 0 {
			q[k] = p[i]
			m[k] = 0
			for n % p[i] == 0 {
				n /= p[i]
				m[k]++
			}
			k++;
		}
	}
	if n > 1 {
		q[k] = n
		m[k] = 1
		k++
	}
}