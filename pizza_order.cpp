#include <iostream>
#include <queue>
#include <vector>
#include <functional>

#define max(a,b) (a>b)?a:b

using namespace std;


bool my_compare(pair<int,int> a, pair<int,int> b )
{
	return (a.second>b.second);
}

//find the min avg wait time given order time array anf cooking len array
//order time is sorted
double find_min_avg_wait_time(vector<int> order_time, vector<int> cook_time)
{
  //if no order return 0;
  if(order_time.size()==0)
    return 0;
  
  //idx is the current index of the orders 
  int idx=0;
  //set global time to the first order 
  int g_time=0;
  int total_wait_time=0;

  //define a priority queue i.e. min heap that stores the order index and cook times for the orders seen at time 
  //lambda function to calculate lesser time for cooking
  auto cmp = [](pair<int,int> a, pair<int,int> b){return (a.second > b.second);};
  //each element in the order_queue is a pair of idx of the order and cook time for that order
  priority_queue<pair<int,int>, vector<pair<int,int> >, decltype(cmp)> order_q(cmp);
  //priority_queue<pair<int,int>, vector<pair<int,int> >, function<bool(pair<int,int>, pair<int,int>)>> order_q;

  //now check orders which come before time t and add them to the priority queue
  //while there are orders left 
  while(!order_q.empty() || idx<order_time.size() )
  //do
    {
      //fast forward time such in case the order queue is empty, and orders are left that will come in future
      if(order_q.empty() && order_time[idx]>=g_time)
	{
	  g_time = order_time[idx];
	}

      //add orders in the queue; order_time is sorted
      while(order_time[idx]<=g_time && idx<order_time.size())
	{
	  //insert the pair
	  order_q.push(make_pair(idx,cook_time[idx]));
	  ++idx;
	}
      //now process the order with the least cooking time
      pair<int,int> min_time_order = order_q.top();
      cout << "Time = " << g_time << endl;
      cout << "processing order # " << min_time_order.first << " with cook len " << min_time_order.second << endl;
      //update global time with the current cooking time
      g_time += min_time_order.second;

      //update the waittime for this order
      total_wait_time += (g_time - order_time[min_time_order.first]);
      cout << "Waiting time: " << g_time << " from " << order_time[min_time_order.first] << endl;
      //remove the order from the queue
      order_q.pop();

          
    }

  return ((double)total_wait_time/order_time.size());
  //return total_wait_time;

}

int main()
{
  vector<int> orders={1,2,3,5,35};
  vector<int> cook_len={2,9,6,7,3};

  cout << "The avg time needed for waiting \n";
  cout << find_min_avg_wait_time(orders, cook_len);
  cout << endl;

  return 1;
}
