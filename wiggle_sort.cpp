#include <iostream>
#include <vector>


using namespace std;

void swap(vector<int> &arr, int a, int b)
{
  int temp = arr[a];
  arr[a]=arr[b];
  arr[b]=temp;
}

void wiggle_sort(vector<int> &arr)
{
  bool less=true;
  for (int i=0; i<arr.size()-1; ++i)
    {
      //swap when the next element doesn't fllow the lesser or greater criteria
      if((less && arr[i]>arr[i+1]) ||(!less && arr[i]<arr[i+1]) )
	{
	  cout << "swapping " << arr[i] << " " << arr[i+1] << endl; 
	  swap(arr, i, i+1);
	}
      if(less)
	less=false;
      else
	less=true;
    }
}

int main()
{
  vector<int> mArr = {17,5,4,1,8,10,2};
  vector<int> mArr2 = {1,2,3,4,5,6,7,8};
  vector<int> mArr3 = {10,9,8,7,6,5,4};
  vector<int> mArr4 = {5,5,5};

  wiggle_sort(mArr4);
  for(int i=0; i<mArr4.size(); ++i)
    {
      cout << mArr4[i] << " " ;
    }
  cout << endl;

  return 0;
}

