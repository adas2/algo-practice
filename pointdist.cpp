#include <iostream>
#include <queue>
//#include <cmath>

//#define dist(x,y) sqrt((x*x)+(y*y))
//Square of the distance is good enough, eliminates extra match operation
#define dist2(x,y) ((x*x)+(y*y))

using namespace std;

typedef struct point{
  double x;
  double y;
}pt;

//given N points vector find k closest from origin
void k_closest(vector<pt> pts_arr, int k)
{
  if (k > pts_arr.size())
    {
      //given array in closest
      return;
    }

  //create priority queue with K points [MAX-heap property by using 'less']
  priority_queue<double, vector<double>, less<int> > k_dist; 
  //calculate distance for first k points points
  vector<pt>::iterator it;
  //Time = O(k)
  for (it = pts_arr.begin(); it < pts_arr.begin()+k; ++it)  
    {
      double distance = dist2((*it).x, (*it).y); 
      k_dist.push(distance); //O(logk)
    }

  //for remaining N-k points choose if they can be included in the k_dist
  //O(N-k) 
  while(it!=pts_arr.end())  
    {
      double distance = dist2((*it).x, (*it).y);
      //distance is less than max dist in the priority queue
      if(distance < k_dist.top())
	{
	  //remove highest dist from queue and push this distance
	  k_dist.pop();   //Time: O(logK)
	  k_dist.push(distance); //Time: O(logK)
	}
      ++it;
    }

  //print the k closest distances
  while(!k_dist.empty())
    {
      cout << "Dist: " << k_dist.top() << endl;
      k_dist.pop();
    }

}


int main()
{
  vector<pt> my_points;
  int a, b;
  int num = 10;

  while (num){
    cout << "Enter points: " << endl;
    cin >> a >> b;
    pt p;
    p.x=a;p.y=b;
    my_points.push_back(p);
    --num;
  }

  k_closest(my_points, 3);

  return 0;
}


