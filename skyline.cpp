#include <iostream>
#include <vector>
#include <climits>

#define max(a,b) (a>b)?a:b

using namespace std;

class Solution {
public:
  vector<pair<int, int> > getSkyline(vector<vector<int> >& buildings) {
    //result vector
    vector<pair<int,int> > res;

    if(!buildings.size())
      return res;

    //find the max of end point for all bulinds
    int max_s = INT_MIN;
    for (auto& b: buildings)
      {//0: start 1: end 2: height
	if (b[1] > max_s)
	  max_s = b[1];
      }
    //exception case
    if(max_s==INT_MAX)
      return res; 

    //create auxillary array and initialize with zeros
    vector<int> sky(max_s+1, 0);
          
    //for each building chnage the aux array with max (height, arr[i])
    for (auto& b: buildings)
      {//0: start 1: end 2: height
	//note: start from start position , end at (end-1)th position
	for (int k=b[0]; k<b[1]; ++k )
	  sky[k]=max(b[2],sky[k]);
      }
    
    //traverse array to see and report change of heights 
    //edge case when builing starts from 0, change in heoght starts from 0
    if(sky[0]!=0)
      res.push_back(make_pair(0,sky[0]));    
    for(int i=0; i<max_s; ++i)
      {
	cout << sky[i] << " ";
	if (sky[i+1]!=sky[i]) 
	  res.push_back(make_pair(i+1,sky[i+1]));
      }
    //for the end adjust 
    //res.push_back(make_pair(max_s,0));
    cout << endl;
     
    //return the result vector
    return res;

  }
 
 
};



int main()
{
  vector<int> b1={0,2147483647,2147483647};

  //vector<int> b1 = {0,1,3};
  //vector<int> b1 = {2,9,10};
  //vector<int> b2 = {3, 7, 15};
  //vector<int> b3 = {5, 12, 12};
  //vector<int> b4 = {15, 20, 10};
  //vector<int> b5 = {19, 24, 8};
  
  vector<vector<int> > buildings;
  buildings.push_back(b1); 
  //buildings.push_back(b2); 
  //buildings.push_back(b3); 
  //buildings.push_back(b4);
  //buildings.push_back(b5);

  Solution s;
  vector<pair<int, int> > result = s.getSkyline(buildings);
  for (auto& x : result)
    {
      cout << x.first << ", " << x.second << endl;
    }
  cout << endl;

  return 1;
}
