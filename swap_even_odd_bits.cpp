#include <iostream>
#include <string>

using namespace std;

//swap the even and odd bits of a 32 bit integer
int swapEvenOdd(int num)
{
  //two masks
  int x_mask=1; //position 0
  int y_mask=2; //position 1

  int result =0;
  int count=1;
  //check if the pos 0 and 1 differ

  while(count<=16)
    {
      //xor of the two bits is 0
      if(!((n&x_mask)^((n&y_mask)>>1)))
	{
	  //result is same as input
	  result |= (n&x_mask);
	  result |= (n&y_mask);
	}
      else //xor is 1 i.e. bits differ
	{
	  //result bits swapped
	  result |= ((n&x_mask)<<1);
	  result |= ((n&y_mask)>>1);
	}
  
      //shift the masks by two bits
      x_mask>>=2;
      y_mask>>=2;
    }
  return result;
}


//Round a number to the next power of 16 i.e. 2^5
int roundNum(int num, int power)
{
  //check power
  if(power%2!=0)
    {
      cout << "invalid power value: " << power << endl; 
      return   -1;
    } 

  //power is a multple of 2
  int mask = power;
  //if(num==mask)
  //return mask;

 while(num>mask)
   {
     //shift left mask
     mask<<=1;
   }
 //mask is >= num
 return mask;
 
}



int main()
{
  int test=0x00000094;
  int power = 16;

  cout << "Even Odd bits swapped for " << hex << num << " is = " << hex << swapEvenOdd(test);
  cout << endl; 
  
  cout << "value of num " << test << " rounded to  next lasrgest power of 16 = " << roundNum(test, power);
  cout << endl;



  return 1;

}
