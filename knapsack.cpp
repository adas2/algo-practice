//Implmentation of (0/1) Knapsack Problem with example

#include <iostream>
#include <vector>
#include <list>

#define MAX 100
#define max(a,b) (a>b)?a:b

using namespace std;

//given list of items and a total weight for the knapsack W
//each item = <value, wieght> pair, both > 0 (+ve integers)
void find_optimal_value(vector<pair<int, int> > items, int W)
{
  //create the matrix of n x W rows columns
  //dummy element first to use matrix indeices 1 to n
  int n=items.size();
  //int mval[n][W];
  vector<vector<int> > mval(n+1);
  //store item list of each weight
  vector<list<int> > item_tracker(n+1);
  //item_tracker.resize();

  //inialize to all zeros
  for(int i=0; i<=n; ++i)
    mval[i].resize(W+1); //this initializes 0 in columns

  //sweep over items and update
  for(int i=1; i<=n; ++i)
  {
    //start from weight 1
    for(int j=1; j<=W; ++j){
      //case 1: element i has w_i > j (current weight): reject item
      if(items[i-1].second > j){
        mval[i][j] = mval[i-1][j];
        item_tracker[i]=item_tracker[i-1];
      }
      //case 2: w_i <= j
      else { //i is taken or i is dropped
        //mval[i][j]= max(mval[i-1][j], mval[i-1][j-items[i-1].second]+items[i-1].first);
        if((mval[i-1][j-items[i-1].second]+items[i-1].first) > mval[i-1][j]){
          mval[i][j]=mval[i-1][j-items[i-1].second]+items[i-1].first;
          //somehow remember the item i-1 added: push into a list
          item_tracker[i]=item_tracker[i-1];
          item_tracker[i].push_back(i-1);
        }
        else {
          mval[i][j]=mval[i-1][j];
          item_tracker[i]=item_tracker[i-1];
        }
      }
    }
  }
  
  //find the lowest right element in matrix i.e. n items with weight W
  cout << "Max possible value of Knapsack: " << mval[n][W] << endl;
  //print list of items
  for(auto x: item_tracker[n])
    {
      cout << x << endl;
    }
}

int main()
{
  vector<pair<int, int> > items;
  //dummy element first to use matrix indeices 1 to n
  //items.push_back(make_pair(0,0));
  items.push_back(make_pair(9,5));
  items.push_back(make_pair(3,2));
  items.push_back(make_pair(4,7));
  items.push_back(make_pair(6,1));
  
  int t_weight = 8; 

  find_optimal_value(items, t_weight);

  return 1;
}


