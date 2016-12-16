#include <iostream>
#include <vector>

class Solution {
public:
  int uniquePathsWithObstacles(std::vector<std::vector<int> >& obstacleGrid) {
        int num_row = obstacleGrid.size();
        int num_col = obstacleGrid[0].size();
        
      if(obstacleGrid[0][0]==1)
        return 0;
     
      std::vector<std::vector<unsigned int> > minP(num_row, std::vector<unsigned int>(num_col, 0));
      //base case if no obstacle at 0,0
      minP[0][0]=1;
      std::cout << "Hey0" << num_row << num_col << std::endl;
      for(unsigned int i=0; i<num_row; i++ )
        {
          for (unsigned int j=0; j<num_col; j++)
            {
	      //std::cout << "Hey1" << std::endl;
	      if(i==0 && j==0)
		continue;
              //case: if obstacle no path 
              if(obstacleGrid[i][j]==1)
		{
                  minP[i][j]=0;
		  
		}
	      else
                {
		  if(i==0)
		    {
		      minP[i][j]=minP[i][j-1];
		    }
		  else if (j==0)
		    {
		      minP[i][j]=minP[i-1][j];
		    }
		  else
		    minP[i][j]=minP[i-1][j]+minP[i][j-1];
		  //minP[i][j]+= (i-1>=0)?minP[i-1][j]:0;
		  //minP[i][j]+= (j-1>=0)?minP[i][j-1]:0;
		}
              
            }
        }
      return minP[num_row-1][num_col-1];
    }
  
};

int main()
{
  unsigned int m, n;
  std::cout<< "Enter the values for m and n" << std::endl;
  std::cin >> m;
  std::cin >> n;
  std::vector<std::vector<int> > obstacleGrid(m, std::vector<int>(n, 0));
  
  obstacleGrid[1][0]=1;
  obstacleGrid[0][1]=1;
  //while(m)

  Solution s;
  std::cout << "The result is: " << s.uniquePathsWithObstacles(obstacleGrid) << std::endl;

}
