//Find the number of trailing zeros in n! (factorial n)
//Solution by Adas2

#include <iostream>

using namespace std;

unsigned long findTrailingZeros(unsigned long n)
{
  //return 1 if zero
  if(n==0)
    {
      return 1;
    }
  int p=5;
  int power = 1;
  unsigned long z_count = 0;

  while (p<=n)
    {
      z_count += (n/p)*power;
      //suntract the xtra ones counted twice
      if(power>1)
	z_count -= (n/p)*(power-1);
      p *= 5;
      power++;
	
    }

  return z_count;

}

int main()
{
  cout << "Enter the number n" << endl;
  unsigned long  n;
  cin >> n;
  cout << "The number of trailing zeros = "  << findTrailingZeros(n) << endl;
  
}
