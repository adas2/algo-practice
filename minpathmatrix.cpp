#include <iostream>
#include <vector>

class Solution {
public:
    int uniquePaths(int m, int n) {
      
      std::vector<std::vector<int> > minP(m, std::vector<int>(n));
      for(unsigned int i=0; i<m; i++ )
	{
	  for (unsigned int j=0; j<n; j++)
	    {
	      if(i==0 || j==0)
		{
		  minP[i][j]=1;
		  //continue;
		}
	      else
		{
		  minP[i][j]=minP[i-1][j]+minP[i][j-1];
		}
	      
	    }
	}
      return minP[m-1][n-1];
    }
};

int main()
{
  unsigned int m, n;
  std::cout<< "Enter the values for m and n" << std::endl;
  std::cin >> m;
  std::cin >> n;
  Solution s;
  std::cout << "The result is: " << s.uniquePaths(m,n) << std::endl;

}
