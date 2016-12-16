#include <iostream>
#include <vector>

using namespace std;

class Solution{
public:
  void mrgsort(vector<int> &arr, int start, int end)
  {
  
    if(end>start)
      {
	int mid = (end-start)/2;
	//vector<int> arr1(arr.begin(), arr.begin()+len/2);
	//vector<int> arr2((arr.begin()+len/2)+1, arr.end());
	cout << "Start End Mid " << start << end << mid <<endl;
	mrgsort(arr, start, mid);
	mrgsort(arr, mid+1, end);
	merge(arr, start, end, mid);
      }

    return;
    //return merge(mrgsort(arr1), mrgsort(arr2));

  }
  
  //merge operation
  void merge(vector<int> &arr, int start, int end, int mid)
  {
    //define two subarrays for arrA, arrB
    //result arrC and clear it contents just in case
    vector<int> arrA(mid-start+1);
    vector<int> arrB(end-mid);
    vector<int> arrC(end-start);
    //arrC.clear();
    
    //copy the contents into two arrays
    for (int k=0; k<=mid; k++)
      arrA.push_back(arr[k]);
    for (int k=mid+1; k<=end; k++)
      arrB.push_back(arr[k]);
   
    
    int i=0;
    int j=0;
    int k=0;
    cout << "Inside Merge() arrA size " << arrA.size() << " arrB size "<< arrB.size() << endl;

    ///while arrA or arrB is empty
    while(i<arrA.size() && j<arrB.size())
      {
	cout << "Inside while" << endl;
	if(arrA[i]>arrB[j])
	  {
	    arrC.push_back(arrB[j]);
	    //arrC[k] = arrB[j];
	    j++;
	    //k++;
	  }
	else
	  {
	    cout << "inside else" << end;
	    arrC.push_back(arrA[i]);
	    //arrC[k] = arrA[i];
	    i++;
	    //k++;
	  }
      }
    //arrA has elements
    if(i<arrA.size() && j==arrB.size())
      {//accC.insert();
	while(i!=arrA.size())
	  arrC.push_back(arrA[i]);
      }
  
    //arrB has elements
    if (j<arrB.size() && i==arrA.size())
      {
	while(j!=arrB.size())
	  arrC.push_back(arrB[j]);
      }

    cout << "Finished copying" << endl;

    //copy the contents back to arr
    for(int m=start; m<=end; m++)
      arr[m]=arrC[m];

    //return arrC;
  }
};

int main()
{
  cout << "Enter the elemnts of the vector" << endl;
  int num;
  cin >> num;
  cout << "Enter values" << endl;
  vector<int> testArr;
  int val;
  while(num>0)
    {
      cin >> val;
      testArr.push_back(val);
      num--;
    }

  cout << "Enetered nums" << endl;
  for(int i=0; i < testArr.size(); i++)
    cout << testArr[i] << " ";
  cout << endl;

  Solution s;
  s.mrgsort(testArr, 0, testArr.size()-1);
  //vector<int>::const_iterator i;
  cout << "Sorted vector : " << endl;
  for(int i=0; i < testArr.size(); i++)
    cout << testArr[i] << " ";
  cout << endl;
  
}




/****
class Soltuion{
public:





  int *mergesort(int *arr, int len)
  {
    //mergesort(arr, 


    return 0;
  }

};
****/

