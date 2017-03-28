/** test implementation of segment tree **/

#include <iostream>
#include <vector>
#include <cmath>

using namespace std;

class SegTree{
//public methods
public:
  //constructor
  SegTree(vector<int> arr, int a_size);
  //destructor
  ~SegTree();
  
  void printTree();

  //update(value at index of original array)
  void update(int idx, int val);
  //search sum from range( i to j )
  int sum_range(int i, int j);
private:
  int s_size; //size of original array
  vector<vector<int> > s_arr; //copy of original array
  int depth;

};

//contructor
SegTree::SegTree(vector<int> arr, int a_size)
{
    //resize the main array
    s_size=a_size;
    //calulcate depth of tree
    depth = ceil(log2(s_size));
    s_arr.resize(depth+1);
    s_arr[0]=arr;
    
    cout << "Depth of seg tree: " << depth << endl;
    
    int i=1; int cnt = s_size; 
    
    //resize the auxillary arrays
    while(i<=depth)
      {
	cnt = ceil(cnt/2);
	cout << "Depth at "<< i << " "<< cnt << endl;
	s_arr[i].resize(cnt);
	//calculate the cum sum for this level
	for(int j=0; j<cnt-1; j+=2){
	  //if last elmemnt add it else sum of pairs
	  s_arr[i][j] = (j==cnt-1)?s_arr[i-1][j]:s_arr[i-1][j]+s_arr[i-1][j+1];
	}
	++i;
      }
    //return;
}

//destructor
SegTree::~SegTree()
{
  s_arr.clear();
  s_size=0;
}

void SegTree::printTree()
{
  for(int i=0; i<=this->depth; ++i)
    {
      for(int j=0; j<this->s_arr[i].size(); ++j)
	cout << this->s_arr[i][j] << " ";
      cout << endl;
    }
  return;
}


int main()
{
  vector<int> arr= {1,7,-3,4,8,9,15,-31,2};
  int s=arr.size();
  SegTree st(arr, s); 
  st.printTree();
  return 1;
}


