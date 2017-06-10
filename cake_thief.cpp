#include <iostream>
#include <vector>
#include <climits>
//#include <limits>
//#include <cstddef>
#include <algorithm>

#define max(a,b) (a>b)?a:b

using namespace std;

class CakeType
{
public:
  size_t weight_;
  long long value_;

  CakeType(size_t weight = 0, long long value = 0) :
    weight_(weight),
    value_(value)
  {
  }
};


//function that implements Knapsack apprroach to get max possible value, when each item type has count of one 
long long maxDuffelBagValue(vector<CakeType> cakeTypes, size_t capacity)
{
  //Methodology: Let MaxV[i][j] be the max possible value with items 1..i and capacility j for the knapsack or dufflebag in this case:
  // i. MaxV[i][j] = MaxV[i-1][j] is the i-th item is not chosen or if it is more than capacity 
  //  ii. MaxV[i][j] = MaxV[i-1][j-weight_of_i] + (value_of_i) i.e. ith item is chosen
  // Oprimzal MaxV[i][j] = Max(case i and case ii) 

  // this depicts the (i) optimal substructure  and (ii) overlapping subproblem properties and hence DP can be used to solve it  

  //number of cakes
  int n=cakeTypes.size();

  //edge case
  if(n<=0)
    return LLONG_MIN;

  //initialize DP matrix
  vector<vector<long long> > MaxV(n+1);
  for(int i=0; i<=n; ++i)
    MaxV[i].resize(capacity+1, 0); //initialize all entries to zero

  //traverse the matrix to fill it out completely uisng above formulae
  //T=O(m*n) 
  for(int i=1; i<=n; ++i)
    {
      for(int j=1; j<=capacity; ++j)
	{
	  //if weight of cake larger than capacity j
	  if(cakeTypes[i-1].weight_>j)
	    MaxV[i][j]=MaxV[i-1][j];
	  else if(cakeTypes[i-1].weight_!=0)
	    {
	      //weight is less than or equal to capacity
	      long long val1=MaxV[i-1][j];
	      long long val2=MaxV[i-1][j-cakeTypes[i-1].weight_] + cakeTypes[i-1].value_;
	      MaxV[i][j]=max(val1,val2);
	      //cout << i << " " << j << " " << MaxV[i][j] << endl;
	    }
	}
    }


  return MaxV[n][capacity]; 
}

bool myCompare(pair<double,int> a, pair<double,int> b)
{
  return (a.first > b.first);
}

//Non-optimal version (might not work for all cases
//if the cake of each type has unlimited supply, different approach
long long maxDuffelBagValue2(vector<CakeType> cakeTypes, size_t capacity)
{
  //num caketypes
  int num = cakeTypes.size();

  if(num<=0)
    return LLONG_MIN;
  
  //create double vector for value/weight for each cakeType, along with the index of cakeType
  vector<pair<double, int> > valueTypes(num);
  //T=O(n)
  for (int i=0; i<num; ++i)
    {
      double val;
      if(cakeTypes[i].value_==0)
	val= 0;
      else if(cakeTypes[i].weight_!=0)
	val = (double)(cakeTypes[i].value_)/(cakeTypes[i].weight_);
      else //weight is zero
	val = (double)INT_MAX; //divide by zero is inf
      valueTypes[i] = make_pair(val,i);
    }
  //sort the value vector in descending order of value/weight; note indices get arranged in that order
  //T = O(nlogn)
  sort(valueTypes.begin(),valueTypes.end(), myCompare);
  
  long long maxV=0;
  size_t space_left=capacity;

  //now choose one by one from the valueTypes array till the capacity is full
  //start with highest value and move to lower values when you can't fit anymore
  //T=O(n)
  for(int i=0; i<num;++i)
    {
      int k=valueTypes[i].second;
      //item will fit and has not zero weight
      if(cakeTypes[k].weight_!=0 && cakeTypes[k].weight_ <= space_left)
	{
	  //num items of type k taken
	  int count = (int)space_left/cakeTypes[k].weight_;
	  maxV += (count*cakeTypes[k].value_);
	  //remaining space after taking item k
	  space_left %= cakeTypes[k].weight_;
	}
      //else move on to the next item;
    }

  return maxV;
  
}

//Dynamic programming solution wiht unbounded Knapsack
//T=O(m*n) where m is the items and n is the capavity
long long maxDuffelBagValue3(vector<CakeType> cakeTypes, size_t capacity)
{
  //total number of items
  int num=cakeTypes.size();

  //edge case capacity zero or less
  if(num<=0)
    return LLONG_MIN;

  //define the MaxV vector , initialize to zero 
  vector<long long> MaxV(capacity+1, 0);
  
  //calculate  max value for a given capacity
  for(int i=0; i<=capacity; ++i)
    {
      //temp varibale to calculate max over all items
      long long curr_max=0;

      //iterate over all items so that we can calculate max value
      for(int j=0; j<num; ++j)
	{
	  //Max value if jth item is taken; if the weight of the cake is less than curr capacity and not zero
	  if(cakeTypes[j].weight_<=i && cakeTypes[j].weight_!=0 )
	    curr_max = max((MaxV[i-cakeTypes[j].weight_] + cakeTypes[j].value_), curr_max); 
	  
	}
      //curr_max is the max value across all items for this caapcity
      MaxV[i]=curr_max;
      
    }
  
  return MaxV[capacity];
  
}


int main()
{
  const vector<CakeType> cakeTypes = {
    CakeType(8, 0),
    CakeType(0, 2000),
    CakeType(7, 160),
    CakeType(3, 90),
    CakeType(2, 15),
  };

  //size_t capacity = 20;
  size_t capacity = 20;
  
  cout << "Max value that thief can have in Duffle Bag with each item with count 1 =  ";
  cout << maxDuffelBagValue(cakeTypes, capacity);
  cout << endl;
  cout << "Max value that thief can have in Duffle Bag with unlimited item count each item type = ";
  //cout << maxDuffelBagValue2(cakeTypes, capacity);
  cout << maxDuffelBagValue3(cakeTypes, capacity);
  cout << endl;
}
