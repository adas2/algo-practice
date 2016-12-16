#include <iostream>
#include <vector>

class Solution {
public:
    int integerBreak(int n) {
        std::vector<int> intBr(n+1, 0);
        //base cases:
        intBr[2]=1;
        
        int max = 0;
        //look back historical data
        for (int i=3; i<=n; i++)
        {
	  for(int j=1; j<=(i/2); j++)
            {
	      if((i-j)>intBr[i-j])
                {
		  max = (i-j)*j;
		  intBr[i]=max;
                }
	      else if(max<intBr[i-j]*j)
                {
		  //case when max is hit
		  max=intBr[i-j]*j;
		  intBr[i]=max;
                }
            }
        }
        return intBr[n];
    }
};


int main()
{
  int num;
  std::cout<< "Enter the values for n" << std::endl;
  std::cin >> num;
  //std::cin >> n;
  Solution s;
  std::cout << "The result is: " << s.integerBreak(num) << std::endl;
}
