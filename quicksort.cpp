#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int partition(int arr[], int low, int high);

void quicksort(int arr[], int low, int high)
{
  //base case end of recursion
  if (high <= low)
    return;
  
  int p = partition(arr, low, high);
  quicksort(arr, low, p-1);
  quicksort(arr, p+1, high);
}

int partition(int arr[], int low, int high)
{
  int p;
  int mid = low+(high-low)/2;
  //initialize the pivot to the array mid value
  p = arr[mid];
  //bring to the front
  swap(arr[low], arr[mid]);
  
  //traverse the array from left and right to align the elemnts < and > p
  int i=low+1;
  int j=high;

  //terminate when j>i
  while(i<=j)
    {
      //cout << "Inside outer while" << endl;
      //find the elemnt larger than p
      while(i<=j && arr[i]<=p)
	{
	  //cout << "Inside inner while i" << endl;
	  i++;
	}

      //find the element smaller than p
      while(i<=j && arr[j]>p)
	{
	  //cout << "Inside inner while j" << endl;
	  j--;
	}
      if(i<j)
	{
	  //swap to align their positions 
	  swap(arr[i], arr[j]);
	}
    }
  //traversal complete swap to insert the pivot in the right position
  swap(arr[i-1],arr[low]);

  return (i-1);
  
}



int main()
{
  //enter the array
  int myArr[] = {110, 5, 10, 3, 22, 100, 1, 23, 10, -2};
  int arrSize = sizeof(myArr)/sizeof(int);
  cout << "Size: " << arrSize << endl;
  for (int i=0; i<arrSize; i++)
    cout << myArr[i] << " ";
  cout << endl;
  
  //call quicksort
  quicksort(myArr, 0, arrSize-1);

  //print the result
   for (int i=0; i<arrSize; i++)
    cout << myArr[i] << " ";
  cout << endl;
  
  return 1;
}
