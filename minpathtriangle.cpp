#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
  int minimumTotal(vector<vector<int> >& triangle) {
    int n = triangle.size();
    //single element case
    //if(n==0)
    //return traingle[0][0]
    
    //start from last but one row
    for(int j=n-2; j>=0; j--)
      {
	for(int i=0; i<=j; i++)
	  {
	    //evaluate last but one from last
	    triangle[j][i]+=min(triangle[j+1][i], triangle[j+1][i+1]);
	  }
      } 

    return triangle[0][0];
    
  }
 
};

int main()
{
  Solution s;
  int row;
  vector<vector<int> > triangle;
  cout << "Enter num rows" << endl;
  cin >> row;
  cout << "Enter value by rows" << endl;
  vector<int> v[row];
  
  for (int i=0; i<row; i++)
    {
      for(int j=0; j<(i+1); j++)
	{
	  int input;
	  cin >> input;
	  v[i].push_back(input);
	}
      triangle.push_back(v[i]);
    }

  cout << "Min Path = " << s.minimumTotal(triangle) << endl;

  return 1;
}
