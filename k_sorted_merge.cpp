/***Problem: Mreged k sorted lists into a sorted list**/

#include <list>
#include <iostream>
#include <queue>
#include <vector>
#include <functional>

using namespace std;
//entry has list-id and corresponding front value
typedef pair<int,int> entry;

//comparator function
bool my_compare (entry a, entry b){
  return (a.second > b.second);
}

void k_sorted_merge(vector<list<int>> lists, int k, list<int> &result)
{
  if (lists.empty())
    return;
  //initialize the list
  result.clear();
  
  //define empty priority queue
  //priority_queue<entry, vector<entry>, decltype(&my_compare)> pq;
  priority_queue<entry, vector<entry>, function<bool(entry, entry)>> pq(my_compare);

  //fill the min queue with first elements of the k lists
  for (int i=0; i<k; ++i)
    {
      if(!lists[i].empty()){
	//int val  = lists[i].front();
	//cout << lists[i].front() << endl;
	entry ent = make_pair(i, lists[i].front());
	pq.push(ent);
      }
    }

  //now that queue is created access min element and add to result
  while(!pq.empty())
    {
      //find the min element and pop the lost with the min element
      result.push_back(pq.top().second);
      int j = pq.top().first;
      //cout << "Queue top: " << pq.top().second << endl;
      //cout << j << endl;
      pq.pop();
      //insert new entry to the queue from the popped list
      lists[j].pop_front();
      //check if list has elements after popping
      if(!lists[j].empty()){
	entry new_ent = make_pair(j, lists[j].front());
	pq.push(new_ent);
      }
    }
  
  return;
}

int main()
{
  list<int> l1 = {10,15,20,25};
  list<int> l2 = {2,12,23};
  list<int> l3 = {30,35,40};
  list <int> l4 = {};
  list<int> result;

  vector<list<int>> my_lists; //= {l1, l2, l3};
  my_lists.push_back(l1);;
  my_lists.push_back(l2);;
  my_lists.push_back(l4);;

  k_sorted_merge(my_lists, 3, result);
  
  //entry e1 = make_pair(2,3);
  //entry e2 = make_pair(5,2);
  //priority_queue<entry, vector<entry>, function<bool(entry, entry)>> pq(my_compare);
  //pq.push(e1);
  //pq.push(e2);

  //print result
  list<int>::iterator it;
  for (it=result.begin(); it!=result.end(); ++it)
    {
      cout << *it << "  "; 
    }
  cout << endl;
  return 1;
}
