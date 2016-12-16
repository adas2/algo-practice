#include <iostream>
#include <vector>

using namespace std;

//prototype
void merge(int arr[], int low, int mid, int high);

void msort(int arr[], int low, int high)
{
  //base case end of recursion
  if(low==high)
    return;

  int mid = low+(high-low)/2;
  msort(arr, low, mid);
  msort(arr, mid+1, high);
  merge(arr, low, mid, high);
 
}

void merge(int arr[], int low, int mid, int high)
{
  /****
  cout << "Array merge:"  << endl;
  for(int m=low; m<=mid; m++)
    cout << arr[m] << " ";
  cout << endl;
  for(int m=mid+1; m<=high; m++)
    cout << arr[m] << " ";
  cout << endl;
  *****/
  vector<int> left;
  vector<int> right;
  
  //left.clear();
  //right.clear();

  //copy the two halves in vectors
  for(int i=low; i<=mid; i++)
    left.push_back(arr[i]);
   for(int i=mid+1; i<=high; i++)
    right.push_back(arr[i]);

   //compare the elemnets left to right and change it in main array
   int i=0;
   int j=0;
   int k=low;
   while (i<left.size() && j<right.size())
     {
       if(left[i] < right[j])
	 {
	   //cout << "left["<< i << "]: " << left[i]<< endl;
	   arr[k] = left[i];
	   i++;
	   k++;
	 }
       if(left[i]>right[j])
	 {
	   //cout << "right["<< j << "]: " << right[j]<< endl;
	   arr[k] = right[j];
	   j++;
	   k++;
	 }
       if(left[i]==right[j])
	 {
	   //cout << "left["<< i << "]: " << left[i]<< endl;
	   //cout << "right["<< j << "]: " << right[j]<< endl;
	   arr[k] = left[i];
	   arr[k+1] = right[j];
	   k+=2;
	   i++;
	   j++;
	 }
     }

   //if elements still left
   while (i<left.size())
     {
       arr[k] = left[i];
       i++;
       k++;
     }

   while (j<right.size())
     {
       arr[k] = right[j];
       j++;
       k++;
     }
   
   return;
}


int main()
{
  //enter the array
  int myArr[] = {110, 5, 10, 3, 22, 100, 1, 23, 10, -9};
  int arrSize = sizeof(myArr)/sizeof(int);
  cout << "Size: " << arrSize << endl;
  for (int i=0; i<arrSize; i++)
    cout << myArr[i] << " ";
  cout << endl;
  
  //call quicksort
  msort(myArr, 0, arrSize-1);

  //print the result
   for (int i=0; i<arrSize; i++)
    cout << myArr[i] << " ";
  cout << endl;
  
  return 1;
}
