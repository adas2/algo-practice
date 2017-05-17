#include <iostream>
#include <vector>
#include <cstdlib>

using namespace std;

void print_arr(vector<int> arr);

void swap(int &a, int &b)
{
	int temp = a;
	a=b;
	b=temp;
}

int partition(vector<int> &arr, int start, int end)
{
	if (start>end)
		return -1;
	int n = end-start+1;
	//randomized partition
	int pivot = rand()%n;
	swap(arr[start+pivot], arr[end]);
	
	pivot = end;
	int firsthigh = start;
	for(int i=start; i<end; ++i)
	{
		//iterate till any elemnt is greater than pivot
		if(arr[i] < arr[pivot])
		{
			swap(arr[i], arr[firsthigh]);
			firsthigh++;
		}
	}
	//
	//cout << "firsthigh : " << firsthigh << endl;
	swap(arr[firsthigh], arr[pivot]);
	//print_arr(arr);
	return firsthigh;
	
}

//utlity ot find the index of the kth smallest element i.e. if k=n/2 its median index.
int quickSelect(vector<int> &arr, int start, int end, int k)
{
	if(start>end)
	{
		return arr[k];
	}
	//else
	int idx = partition(arr, start, end);
	//cout << idx << " " <<k <<  " "  <<start <<endl; 
	if(idx-start == k)
	{
		//cout << arr[start] << " " << arr[idx] << endl;
		return arr[idx];
	}
	else if (idx-start>k)
		return quickSelect(arr, start, idx-1, k);
	else //idx<k
		return quickSelect(arr, idx+1, end, k-idx-1);
		//return quickSelect(arr, idx+1, end, k-idx+start);
}


int minMoves2(vector<int>& nums) 
{
	int moves=0;
       	//find median of the array
	int median_idx = (nums.size()/2);
	cout << "Median idx = " << median_idx << endl;
	int median_val = quickSelect(nums, 0, nums.size()-1, median_idx);
	cout << "Median val = " << median_val << endl;
	//iterate through the array to calculate the moves 
	for (int i=0; i<nums.size(); ++i)
	{
		moves += abs(nums[i]-median_val);
	}
	return moves;

}

void print_arr(vector<int> arr)
{
	for(auto &x: arr)
		cout << x << " ";
	cout << endl;
}


int main()
{
	//vector<int> test = {10,2,11,1,3 };
	vector<int> test = {1,2,5,8,0 };
	cout << "Num moves: " << minMoves2(test);
	cout << endl;
	return 1;
}
