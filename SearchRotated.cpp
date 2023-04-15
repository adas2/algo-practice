/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

#include <iostream>
#include <vector>

using namespace std;

// given an sorted array with arbitrary number of rotations, find the index of a given element
// start and end indices are given
// assumes all unique elements
int search_rotated_array(vector<int> arr, int start, int end, int num)
{
    // null or case not found
    if (end < start)
        return -1;

    // central element
    int mid = (end + start) / 2;

    cout << "start: " << start << " mid: " << mid << " end: " << end << endl;
    // if equal to mid
    if (arr[mid] == num)
        return mid;

    // find if the inflection point is left or right of mid
    // left half sorted and num in left half
    if (arr[start] < arr[mid])
    {
        if (arr[start] < num && num < arr[mid])
        {
            // recursively call to search function
            return search_rotated_array(arr, start, mid - 1, num);
        }
        else
        {
            // num or right half but with inflection point
            return search_rotated_array(arr, mid + 1, end, num);
        }
    }
    // right half is sorted and num is in right half
    else if (arr[mid] < arr[end])
    {
        if (arr[mid] < num && num < arr[end])
        {
            // recursively call search on right half
            return search_rotated_array(arr, mid + 1, end, num);
        }
        else
        {
            // num in left half but left half is unsorted
            return search_rotated_array(arr, start, mid - 1, num);
        }
    }
    else // error case
        return -1;

    // NOTE:
    // if repeated element check if one half has no distinct element at start/end
    // binary search in that half
    // if mid==start==end linear search is needed
}

int main()
{
    vector<int> test = {4, 5, 6, 7, 1, 2};

    cout << "Index found: " << search_rotated_array(test, 0, 5, 1) << endl;
}