#include <iostream>
#include <queue>
#include <cmath>

#define dist(x,y) sqrt((x*x)+(y*y))

using namespace std;

typedef struct point{
  double x;
  double y;
}pt;

//given N point vector find k cloests from origin
void k_closest(vector<pt> pts_arr, int k)
{
  if (k > pts_arr.size())
    {
      //given array in closest
      return;
    }

  //create prioity queue with K points
  priority_queue<double, vector<double>, less<int> > k_dist; 
  //calculate distance for first k points points
  vector<pt>::iterator it = pts_arr.begin();
  for (int i=0; i<k; ++it, ++i)  //O(k)
    {
      double distance = dist((*it).x, (*it).y); 
      k_dist.push(distance);
    }

  //for remaining N-k points choose if they can be included in the k_dist
  while(it!=pts_arr.end())  //O(N)
    {
      double distance = dist((*it).x, (*it).y);
      if(distance<k_dist.top())
	{
	  //remove highest dist from queue and push this distance
	  k_dist.pop();   //Time: O(logK)
	  k_dist.push(distance); //Time: O(logK)
	}
      ++it;
    }

  //print k closest points
  //priority_queue<double>::iterator qt;
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
  int num = 5;

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


