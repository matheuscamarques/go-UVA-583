
/**
 * Fatores primos
 * Crivo de Erestóstenes, todos numeros primos em um intervalo inteiro
 * Otimização de O(NloglogN) -> o(sqrt(N)loglogsqrt(N))
 * 
 * build 
 *     g++ main.cpp -o main
 */

#include <bits/stdc++.h>
#define MAX 112345
using namespace std;


int p[MAX],pi,q[MAX],m[MAX],k;


void sieve(int n) {
    p[0] = 2; pi = 1;
    for (int i = 3; i <= n; i++) {
        int prime = 1 , r = sqrt(i);
         for(int j = 0; j < pi && p[j] <= r && prime; j++) {
             prime &= (i % p[j] != 0);
             if(prime) p[pi++] = i;
         }
    }
}

void fact(int n) {
  k = 0;
  for(int i = 0; i < pi && p[i] * p[i] <= n; i++) 
      if(n % p[i] == 0) {
          q[k] = p[i];
          m[k] =0;
          while (n % p[i] == 0) {
              m[k]++;
              n /= p[i];
          }
          k++;
      }
    
    if(n > 1){ q[k] = n; m[k] = 1; k++; }
  
}



int main(){
    int n,first;
    sieve(MAX);

    while(scanf("%d",&n),n){
        fact(abs(n));
        printf("%d =%s",n,n < 0 ? " -1" : "");
        first = n > 0;
        for(int i = 0; i < k; i++) {
            for (int j = 0; j < m[i]; j++) {
                printf("%s%d", first ? " " : " x ",q[i]);
                first = 0;
            }
        }

         printf("\n");

    }

    return 0;
}