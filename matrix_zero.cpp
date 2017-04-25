#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  void setZeroes(vector<vector<int> >& matrix);
};

void Solution::setZeroes(vector<vector<int> >& matrix) 
{
  if(matrix.size()==0)
    return;

  bool row_start=false;
  bool col_start=false;

  //scan the mayrix to check for zeros
  for(int i=0; i<matrix.size(); ++i)
    for(int j=0; j<matrix[0].size(); ++j)
      {
	//mark the column and row starting with zero
	if(matrix[i][j]==0){
	  if(j==0)
	    col_start=true;
	  if(i==0)
	    row_start=true;
	  //row and column markers
	  matrix[i][0]=0;
	  matrix[0][j]=0;
	}

      }

  /** Better way: Traverse backwards to avoid tampering the 0th row and 0th column first 
  ***/
  

  // set the zeros by traversing the rows and columns;
  // NOTE: traverse the 0th column and row last otherwise markers will be overwritten
  //set rows 1 to m
  for(int i=1; i<matrix.size(); ++i)
    {
      
      if(matrix[i][0]==0)
	for(int j=0; j<matrix[0].size(); ++j)
	  matrix[i][j]=0;
    }
  
  //set columns 1 to n
  for(int j=1; j<matrix[0].size(); ++j)
    {
     
      if(matrix[0][j]==0)
	{
	  for(int k=0; k<matrix.size(); ++k)
	    matrix[k][j]=0;
	}
    }
  //set the 0th column and 0th row
  if(matrix[0][0]==0 && col_start){
    for(int k=0; k<matrix.size(); ++k)
      matrix[k][0]=0;
  }
  if (matrix[0][0]==0 && row_start){
    for(int j=0; j<matrix[0].size(); ++j)
      matrix[0][j]=0;
  }
 
	
}

int main()
{
  Solution s;
  vector <vector<int> > test= {{1,1,1}, {1,1,1}, {0,1,0}};
  //print
  for(int i=0; i<test.size(); ++i)
    {
      for (int j=0; j<test[0].size(); ++j)
	cout << test[i][j] << " ";
      cout << endl;
    }
  cout <<"========" << endl;  
  s.setZeroes(test);
  //print
  for(int i=0; i<test.size(); ++i)
    {
      for (int j=0; j<test[0].size(); ++j)
	cout << test[i][j] << " ";
      cout << endl;
    }
    
  return 0;
}
