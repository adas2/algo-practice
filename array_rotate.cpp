#include <iostream>
#include <vector>


using namespace std;
/****
void swap(int *a, int *b)
{
  cout << "Swapping: " << *a << " " << *b << endl;
  int temp = *a;
  *a = *b;
  *b = temp;
  return;
}
****/

void reverse(vector<int> &arr, int start, int end)
{
  int i=start, j=end;
  while (i<j)
    {
      int temp = arr[i];
      arr[i] = arr[j];
      arr[j]=temp;
      i++;j--;
    }
  return;
}

void rotate_array(vector<int> &arr, int rotate_len)
{
  //int i, k=rotate_len;
  if(arr.size()==0)
    return;
  
  //check if rotation lenght is larger than array
  if (rotate_len > arr.size())
    rotate_len %= arr.size();

  int n = arr.size();
  int d = rotate_len;

  //1. Reverse array [0,n-1];
  reverse(arr, 0, n-1);
  //2. Reverse array [0, n-d-1];
  reverse(arr, 0, n-d-1);
  //3. Reverse array [n-d, n-1];
  reverse(arr, n-d, n-1);
  
  return;

}

int main()
{
  vector<int> array;
  array.push_back(1);
  array.push_back(2);
  array.push_back(3);
  array.push_back(4);
  array.push_back(5);
  array.push_back(6);
  array.push_back(7);

  vector<int>::iterator it;
  int k;

  for(it=array.begin(); it !=array.end(); ++it)
    {
      cout << *it << " ";
    }
  cout << endl;

  cout << "Rotate Len: " << endl;
  cin >> k;
 
  rotate_array(array, k);
  
  cout << "After rotation: " << endl;
  
  for(it=array.begin(); it !=array.end(); ++it)
    {
      cout << *it << " ";
    }
  cout << endl;
  
 
}
