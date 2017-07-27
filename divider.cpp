#include <iostream>
#include <stdlib.h>
#include <climits>

//#define MAX_INT 2147483647
//#define MIN_INT –2147483648


using namespace std;

class Solution {
public:
    //Divide without using the division operator
  int divide(int dividend, int divisor) {
    
    int remainder = abs(dividend);
    int dvsr = abs(divisor);
    //set sign bit XOR of dividend -ve? and divsr -ve?
    int sign = ((dividend < 0) ^ (divisor < 0)) ? -1:1;
   

    cout << remainder << " " << dvsr << endl;
    //overflow condition
      if (!divisor)
	return INT_MAX;
      if(dividend==0)
	return 0;
     
      if(dividend == INT_MIN) 
	{
	  if (divisor == -1)
	    return INT_MAX;
	  else if (divisor == 1)
	    return INT_MIN;
	  else
	    {
	      if((divisor&1)==1) //divisor odd, dividend even
		return divide(dividend+1, divisor);
	      else //divisor even, dividend even
		return divide (dividend>>1, divisor>>1);
	    } 
	}
      if(divisor == INT_MIN)
	return 0;

      //if(divisor == 1)
      //return dividend;
    
      int quotient = 0; //
      
      //subtract till remainder is less that divisor
      while(remainder >= dvsr)
	{
	  //increase the dvisior size by a factor 2 by shifting
	  int temp = dvsr;
	  int factor = 1; //starts with 1
	  // O(log-n) steps
	  while(remainder >= (temp << 1) && ((temp<<1)> dvsr))
	    {
	      temp <<= 1;
	      factor <<= 1;
	      // cout << "hello" << endl;
	    }

	  remainder -= temp;
	  quotient += factor;
	  //cout << "Quotient: " << quotient << endl;
	}
      //cout << "Remainder: " << remainder << endl;
      return quotient*sign;
    }
};

int main()
{
  cout << "Enter the dividend and divisor" << endl;
  int dividend, divisor;
  cin >> dividend >> divisor ;
  cout << "Dividend: " << dividend << " Divisor: " << divisor << endl;
  Solution s;
  cout << "Answer is : " << s.divide(dividend, divisor) << endl; 
  return 0;
}
