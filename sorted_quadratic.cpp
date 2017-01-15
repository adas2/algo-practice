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

//Uses O(n) storage
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
  arr.push_back(-3);
  arr.push_back(1);
  arr.push_back(3);
  arr.push_back(7);
  arr.push_back(10);
  arr.push_back(12);

  int a=-2, b =16, c=5;
  
  //cout << bin_search(arr, 0, 5, 8) << endl;
  sort_quadratic_arr(arr, a, b, c);
  cout << "Sorted array: " << endl;
  for(int i=0; i<arr.size(); ++i)
    {
      cout << arr[i] << " ";
    } 
  cout << endl;
  
}


