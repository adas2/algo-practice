/* Given the vector of IDs, which contains many duplicate integers and one unique integer, find the unique integer.  ***/

#include <iostream>
#include <vector>
#include <climits>

using namespace std;

int find_unique_int_arr(vector<int> arr)
{
  //empty array or 1 size
  int n = arr.size();
  if (n==0)
    return INT_MIN;
  if (n==1)
    return arr[0];

  //define temprary varibale and iterate through the array xoring every element 
  int result=0;
  for (int i=0; i<n; ++i)
    {
      result ^= arr[i];
    }
  
  return result;

}


//find 2 uniqye elements except which all are repeated
pair<int,int> find_2_unique_int_arr(vector<int> arr)
{
  //empty array or 1 size
  int n = arr.size();
  if (n==0 || n==1)
    return make_pair(-1,-1); 
  if (n==2)
    {
      return make_pair(arr[0],arr[1]);
    }

  //let the elements be a and b
  //pass 1: traverse the array to detremin a xor b
  int exor=0;
  for (int i=0; i<n ; ++i)
    {
      exor ^= arr[i];
    }
  //bits set in exor are the bits that are different in a and b; e.g. a=010 and b=011 a xor b = 001 
  // i.e. last bit is the one that differs
  //  create the bitmask with last bit set in exor
  int mask = exor & ~(exor-1); //this just sets the last 1 bit and rest 0
  
  //pass 2: now scan the arrya nd divide the elements in two sets where this bit is set and where not
  int num1=0, num2=0;
  for(int i=0; i<n; ++i)
    {
      if(arr[i] & mask)
	{
	  //bit set
	  num1 ^=arr[i];
	}
      else
	{
	  //bit not set
	  num2 ^=arr[i];
	}
    }
  
  return make_pair(num1,num2);

}


//find unique element in array when all other elements are repeated 3 times
int find_3_unique_element_arr()
{


}


int main()
{
 // vector<int> test={2,4,8,3,4,8,2};
  vector<int> test={2,3,4,2,3,5,5,7};

  //cout << "Unique element in array: " << find_unique_int_arr(test);
  cout << "Unique element in array: " << find_2_unique_int_arr(test).first << " " << find_2_unique_int_arr(test).second ;
  cout << endl;

  return 1;
}

