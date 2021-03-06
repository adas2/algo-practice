/** test implementation of segment tree **/

#include <iostream>
#include <vector>
#include <cmath>
#include <climits>

using namespace std;

class SegTree{
//public methods
public:
  //constructor
  SegTree(vector<int> arr, int a_size);
  //destructor
  ~SegTree();
  //print values
  void printTree();
  int getDepth();
  //update(value at index of original array)
  void update_sum(int location, int val, int dpth, int idx);
  //search sum from range( i to j )
  int get_sum (int start, int end, int dpth, int idx);
  

private:
  int s_size; //size of original array
  vector<vector<int> > s_arr; //copy of original array
  int depth;
  //??KEEP A STRUCT FOR KEEPING RANGES
};

//constructor
SegTree::SegTree(vector<int> arr, int a_size)
{
    //resize the main array
    s_size=a_size;
    //calculate depth of tree
    depth = ceil(log2(s_size));
    //create multi-level array, lowest level original array, highest level is the cumulative sum
    s_arr.resize(depth+1);
    s_arr[0]=arr;
    
    cout << "Depth of seg tree: " << depth << endl;
    
    int i=1; int cnt = s_size; 
    //cout << "main array size " << ceil((double)cnt/2) << endl;
        
    //resize the auxiliary arrays
    //i: current depth of the tree
    while(i<=depth)
      {
	cnt = ceil((double)cnt/2);
	//cout << "Depth at "<< i << " "<< cnt << endl;
	s_arr[i].resize(cnt);
	//calculate the cum sum for this level
	//j : index of previous level, k: index of current level
	for(int j=0, k=0; k<cnt; ++k,j+=2){
	  //if last element add it else sum of pairs
	  s_arr[i][k] = (j==(s_arr[i-1].size()-1))?s_arr[i-1][j]:s_arr[i-1][j]+s_arr[i-1][j+1];
	}
	++i;
      }
    //return;
}

//destructor
SegTree::~SegTree()
{
  for(int i=0 ; i<=depth; ++i)
    s_arr[i].clear();
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

// s: start e: end such that s->e is thus query range; d: current depth; idx: starting index for current level
//O(logN) operations to get sum for a range
int SegTree::get_sum(int s, int e, int d, int idx)
{
  //corner case check
  if(s<0 || e>=s_arr[0].size())
    return INT_MIN;

  //start with the deepest array (root node)
  //coverage range for this depth
  int coverage = (1<<d); //i.e. idx to idx+coverage-1
  // start and end of range for this depth
  int s_i = idx*coverage; //this is the starting index of the actual array covered by this range
  int s_j = s_i + coverage - 1; //this is the ending index of the actual array covered by this range
  //if the range covered by this level is completely within the query range then return value;
  if(s <= s_i && e >= s_j){
    //cout << "covered at depth, indx " << d << " " << idx << " returning: " << s_arr[d][idx] << endl;
    return s_arr[d][idx];
  }
  //if out of range return zero
  if(s > s_j || e < s_i )
    {
      //cout << "Out of range at depth, idx  " << d << " " << idx << " returning: 0" << endl;
      return 0;
    }
  //else if there is a partial overlap between the query range and range of the tree
  //    go one level up and send the sum of two halves
  return get_sum(s, e, d-1, 2*idx) + get_sum(s, e, d-1, 2*idx+1); 
}

//update the value at the index in the original array by adding val
//NOTE: val is the delta
//loc: index of the increment, val: new delta to add to, d: current depth, idx: starting index for the current depth
void SegTree::update_sum(int loc, int val, int d, int idx)
{
  //start at root node lowest depth
  int coverage = 1<<d;
  int s_i = idx*coverage;
  int s_j = s_i + coverage-1;
  //if the index lies in the range of the node update the sum by val
  if(loc >= s_i && loc <= s_j)
    {
      s_arr[d][idx] += val;
      //recursively call halves of the one level above array
      if(d!=0){
	update_sum(loc, val, d-1, 2*idx);
	update_sum(loc, val, d-1, 2*idx+1);
      }
    }
  //else
  return; 
  
}

int SegTree::getDepth()
{
  return this->depth;
}



int main()
{
  vector<int> arr= {1,7,-3,4,8,9,15,-31,2};
  int s=arr.size();
  SegTree st(arr, s); 
  st.printTree();
  
  cout << "Sum from 2 to 6= " << st.get_sum(2,6,st.getDepth(),0) << endl;
  
  st.update_sum(8, 15, st.getDepth(), 0);
  cout << "After update\n" ;
  st.printTree();
  return 1;
}


