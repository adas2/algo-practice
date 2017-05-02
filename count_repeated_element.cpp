#include <iostream>

using namespace std;

int cnt_repeated_elem(int arr[], int elem, int start, int end)
{

  //cout << "Start: " << start << " End: " << end << endl;
  //base case:
  if (start>end)
    return 0;
  //if element doesn't lie in the array return 
  if(arr[start]>elem || arr[end]<elem)
    return 0;
  
  //int cnt=0;
  int mid = (start + end)/2;
  
  //case when first and last elements are equal
  if(arr[end]==elem && arr[start]==elem)
    {
      //count is equal to size of the array
      //if(arr[start]==arr[end])
      //  {
      //    cnt += (end-start)+1;
      //    cout << "Match start end" << endl;
      //    return cnt;
      //  }
      return (end-start)+1;
      
    }

  //find which half has the elem
  if (arr[mid] > elem )
    {
      //count only from start to mid
      //splt array in half and iterate
      return cnt_repeated_elem(arr, elem, start, mid-1);     
    }
  else if (arr[mid] < elem )
    {
      //middle elemnt is less, elem could be in the right half
      return cnt_repeated_elem(arr, elem, mid+1, end);
    }
  else
      return cnt_repeated_elem(arr, elem, start, mid)+ cnt_repeated_elem(arr, elem, mid+1, end);     
  //return cnt;
  
}

int main()
{
  int array[10] = {-2, 1,2,2,2,4,5,6,6, 9};
  int num = -6;
  cout << "There element: " << num << " has count: " << cnt_repeated_elem(array, num, 0, 9) << endl;

}



