#include <iostream>
#include <vector>
#include <climits>

#define min(a,b) (a<b)?a:b

using namespace std;

int minNeighbor(vector<vector<int> > m, int i, int j)
{
  int _min=INT_MAX;
  int row_s = m.size();
  int col_s = m[0].size();

  int A = min(((i>0)?m[i-1][j]:INT_MAX), ((i<row_s-1)?m[i+1][j]:INT_MAX));
  int B = min(((j>0)?m[i][j-1]:INT_MAX), ((j<col_s-1)?m[i][j+1]:INT_MAX));
  _min = min(A,B);

  cout << "min for " << i  <<" "<< j << " " << _min<< endl;
  return _min;
}

vector<vector<int> > updateMatrix(vector<vector<int> >& matrix) 
{
  int m=matrix.size();
  int n=matrix[0].size();

  vector<vector<int> > result;
  //initialize result matrix with INT_MAX
  result.resize(m);
  for(int i=0;i<m; ++i)
    result[i].resize(n,INT_MAX);

  //1. traverse input matrix top down  
  for(int i=0; i<m; ++i)
    for(int j=0; j<n; ++j)
      {
	if(matrix[i][j]==0)
	  result[i][j]=0;
	else
	  {
	    int temp=minNeighbor(result, i, j); 
	    if(temp!=INT_MAX) //avoids overflow situation
	      result[i][j]=temp + 1;
	  }
      }
  
  /****/
  //2. traverse again bottom up
  for (int i=m-1; i>=0; --i)
    for(int j=n-1; j>=0; --j)
      {
      if(matrix[i][j]==0)
	  result[i][j]=0;
	else
	  {
	    int temp=minNeighbor(result, i, j); 
	    if(temp!=INT_MAX) //avoids overflow situation
	      result[i][j]=temp + 1;
	  }

      }
  /****/

  return result;
}

int main()
{
  vector <vector<int> > test = {{1,2,3}, {4,5,6}, {7,0,8}};
  vector <vector<int> > result = updateMatrix(test);
  
  for(int i=0; i<test.size(); ++i)
    {
      for(int j=0; j<test[0].size(); ++j)
      { 
	cout << result[i][j] << " ";
      }
    cout << endl;
    } 

  return 0;
}
