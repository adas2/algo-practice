/***
Suppose that you have a sorted array of integers (positive or negative). You want to apply a function of the form f(x) = a * x^2 + b * x + c to each element x of the array such that the resulting array is still sorted. Implement this in Java or C++. The input are the initial sorted array and the function parameters (a, b and c).
****/

#include <vector>
#include <iostream>
#include <algorithm>

using namespace std;


int bin_search(vector<int> arr, int start, int end, int val)
{
  int len=(end-start)+1;
  int mid = start + (len/2);
  //cout << "start: " << start << " mid: " << mid << " end: " << end << endl; 
  //terminating condition
  if (start==end)
    return start;

  if(arr[mid]> val)
    return bin_search(arr, start, mid-1, val);
  else if (arr[mid] < val)
    return bin_search(arr, mid, end, val);
  else //equal
    return mid;
  //else
  //return mid;
}

int calculate_value(vector<int> &arr, int i, int a, int b , int c){
  return ((arr[i]*arr[i]*a)+ (arr[i]*b) + c);
}

bool mycompare(int p, int q)
{
  return (p<q);
}

//O(1) storagea O(n^2) time : Naive approach
// argument backwards tells whether top or bottom needs to be traveresed backwards:
//            1: for top backwards, 2: for bottom backwards, 0: Both ascending order
void inPlaceMerge(vector<int> &arr,int top_start, int top_end, int bottom_start, int bottom_end, int backwards )
{
  int i=top_end, j=bottom_end;
  //only j moves down and while it reaches the end of top
  while(j>top_end){
    //cout << arr[i] << " " << arr[j] << endl;
    if(arr[i] > arr[j]) //top last is greater than bootk last
      {
	//swap the largest element to the bottom 
	int temp = arr[i];
	arr[i] = arr[j];
	arr[j] = temp;
	
	//important: change ponter for bootm
	//--j;
	
      }
    //decrement to next element
    --j;
    //cout << arr[i] << " " << arr[i-1] << endl;
    //if the new top last element is smaller than previos change top index, else keep same
    
    while ((arr[i] < arr[i-1]) && (i> top_start))
      {
	//cout << "SWapped " <<arr[i] << " " << arr[i-1] << endl;
	//swap a[i] and a[i-1]
	int temp = arr[i-1];
	arr[i-1]=arr[i];
	arr[i]=temp;
	--i;
      }
    //restore i to top_end
    i=top_end;
    
  }
}

//Uses O(n) storage in linear time i.e. O(n); 
//   assumes top part from start to mid is sorted and bootm from mid to end is sorted
void merge(vector<int> &arr, int start,int mid, int end)
{
  //make a new array of lenght n
  vector<int> result;
  
  //merge copy the contents one by one
  int i=start, j=mid;
  while(i<mid && j<=end ){
  if(arr[i]<arr[j])
    {
      result.push_back(arr[i]);
      ++i;
    }
  else
    {
      result.push_back(arr[j]);
      ++j;
    }
  }
  while(j<=end)
    {
      result.push_back(arr[j]);
      j++;
    }
  while (i<mid)
    {
      result.push_back(arr[i]);
      i++;
    }

  i=0;
  //copy the new array into the old one
  vector<int>::iterator it = result.begin();
  while(it!=result.end())
    {
      arr[i++]= (*(it++));
    }
  return;
}

void sort_quadratic_arr(vector<int> &arr, int a, int b, int c)
{
  int min_max = ((-1)*b)/(2*a);
  int idx = bin_search(arr, 0, arr.size()-1, min_max);
  //evaluate the resltant array 
  for(int i=0; i < arr.size(); ++i )
    {
      arr[i]=calculate_value(arr, i, a, b, c);
    }   
  cout << "original array after function:" << endl;
  for(int i=0; i<arr.size(); ++i)
    {
      cout << arr[i] << " ";
    } 
  cout << endl;
  
  //check the concavity
  if(a<0)
    //concave downwards
    {
      //evaluate and sort right side
      sort(arr.begin()+idx+1, arr.end(), mycompare);
    }
  else
    //concace upwards
    {
      //evaluate and sort left side
      sort(arr.begin(), arr.begin()+idx);
    }
  
  //merge the two sorted half arrays (can you do inplace?)
  merge(arr, 0, idx+1, arr.size()-1);
}


int main()
{
  vector <int> arr;
  /***/
  arr.push_back(-3);
  arr.push_back(1);
  arr.push_back(3);
  arr.push_back(7);
  arr.push_back(10);
  arr.push_back(12);
  /****/
  /***
  arr.push_back(1);
  arr.push_back(3);
  arr.push_back(5);
  arr.push_back(7);
  arr.push_back(15);
  arr.push_back(2);
  arr.push_back(4);
  arr.push_back(6);
  arr.push_back(11);
  ****/


  int a=-2, b =16, c=5;
  
  //cout << bin_search(arr, 0, 5, 8) << endl;

  /****/
  sort_quadratic_arr(arr, a, b, c);
  /****/
  cout << "Original array: " << endl;
  for(int i=0; i<arr.size(); ++i)
    {
      cout << arr[i] << " ";
    } 
  cout << endl;
  //inPlaceMerge(arr, 0, 4, 5, 8, 0);
  cout << "Sorted array: " << endl;
  for(int i=0; i<arr.size(); ++i)
    {
      cout << arr[i] << " ";
    } 
  cout << endl;
  

  

  
}


