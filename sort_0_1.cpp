#include <iostream>
#include <vector>


using namespace std;

/****
template<class T>
void swap(T &a, T &b )
{
	T temp;
	temp=a;
	a=b;
	b=temp;
	return;
}
*****/

void swap(short &a, short &b )
{
	short temp;
	temp=a;
	a=b;
	b=temp;
	return;
}


void sort_0_1(vector<short> &arr)
{
	// array size
	int n=arr.size();
	//empty array
	if(!n)
		return;

	// Solution 1. First pass count the number of 0's and then Second pass replace first count elements to zeros and rest as 1.
	// solution becomes two passes and also lot of writes

	//Solution 2. single pass only,
	//maintain lft and right indices; increment left till it hits a 1 and decrement right till it hits a 0
	// if left < right swap the elements and continue
	int left =0;
	int right =n-1;

	//O(n) operations
	while(left<right)
	{
		while(arr[left]!=1)
			++left;
		while(arr[right]!=0)
			--right;
		if(left<right)
			swap(arr[left], arr[right]);
		else
			return; //already sorted


	}

	return;

}


int main()
{
	vector<short> arr={1,0,0,1,1,1};
	sort_0_1(arr);

	cout << "Sorted array " << endl;
	for(auto x: arr)
	{
		cout << x << " ";
	}
	cout << endl;
}

