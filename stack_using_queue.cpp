#include <queue>
#include <iostream>

using namespace std;

class qStack{
public:
  qStack()//default constructor
  {
    cnt=0;
    //q1.clear();
    //q2.clear();
  }
  void push(int val);
  int pop();
  int top();
private:
  queue<int> q1,q2;
  int cnt;
};

void qStack::push(int val)
{
  //alternate between empty and non-empty queue
  
  if(q1.empty())
    {
      //1. enqueue the element in the empty queue
      q1.push(val);    
      //2. for each element in second queue dequeue and enqueue in first queue
      while(!q2.empty())
	{
	  q1.push(q2.front());
	  q2.pop();
	}
      return;
    }
  else //if(q2.empty())
    {
      //use second queue to push
      q2.push(val);
      while(!q1.empty())
	{
	  q2.push(q1.front());
	  q1.pop();
	}
      return;
    }
    
}

int qStack::pop()
{
  int val;
  if(q1.empty())
    {
      val = q2.front();
      q2.pop();
    }
  else // q2.empty()
    {
      val = q1.front();
      q1.pop();
    }
  return val;
}

int main()
{
  qStack my_stack;
  my_stack.push(1);
  my_stack.push(2);
  my_stack.push(3);
  
  //pop ops
  cout << my_stack.pop() << endl;
  cout << my_stack.pop() << endl;
  cout << my_stack.pop() << endl;
}
