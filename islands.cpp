#include <iostream>
#include <vector>
#include <string>

using namespace std;

class Solution {
public:
  //helper function find is there is a  neighbor for grid[i][j]
  void visitNeighbors(vector<vector<char> >& grid, int i, int j, vector<vector<bool> >& visited )
  {
    int r_size = grid.size();
    int c_size = grid[0].size();
    //is any element is this island visited before?

    //visit node if not visited
    if(!visited[i][j])
      visited[i][j]=true;
    
    //visit vertical neighbors
    if((i<r_size-1) && (grid[i+1][j]=='1') && (!visited[i+1][j]))
      {
	visitNeighbors(grid,i+1,j,visited);
      }
    if((i>0) && (grid[i-1][j]=='1') && (!visited[i-1][j]))
      {
	visitNeighbors(grid,i-1,j,visited);
      }
    //visit  horizontal 
    if(j<c_size-1 && (grid[i][j+1]=='1') && (!visited[i][j+1]))
      {
	visitNeighbors(grid,i,j+1,visited);
      }
    if(j>0 && (grid[i][j-1]=='1') && (!visited[i][j-1]))
      {
	visitNeighbors(grid,i,j-1,visited);
      }
    return;
  }

  int numIslands(vector<vector<char> >& grid) {
    vector<vector<bool> > visited ;
    int r_size=grid.size();
    if(!r_size)
      return 0;
    int c_size=grid[0].size();

    visited.resize(r_size);
    for(int i=0; i<r_size; ++i)
      visited[i].resize(c_size, false);

    //count of islands
    int count=0;

    for(int i=0; i<r_size; ++i)
      for(int j=0; j<c_size; ++j)
      {
	if(grid[i][j]=='1' && !visited[i][j] )
	  {
	    //start of a new island
	    count++;
	    //DFS/recursively visited the neighbors for this isalnd
	    visitNeighbors(grid, i,j,visited);
	  }
      }
    return count;
  }
};


int main()
{
  /***/
  vector<vector<char> > test= {{'1','0','1', '1', '1'},
			       {'1', '0','1','0', '1'},
			       {'1','1','1','0','1'}
			       //{'0','1','0'}
  };
  /***/
 //vector<vector<char> > test= {{"10111"},{"10101"}, {"11101"}}; 

  Solution s;
  cout << "Num islands: " << s.numIslands(test) << endl;

}

