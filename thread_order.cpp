/** Execute threads functions in their thread id order ***/
#include <iostream>
#include <thread>
#include <condition_variable>
#include <mutex>
#include <vector>
#include <unistd.h>

#define MAX_THREADS 4

using namespace std;

mutex mtx;
condition_variable cnd;
//shared variable
int curr=0;

class Foo{
public:
  void first(){cout << "I'm first" << endl;}
  void second(){cout << "I'm second" << endl;}
  void third(){cout << "I'm third" << endl;}
};

//
void c_func(int id, Foo f)
{
  //acquire lock and wait of cond var
  unique_lock<mutex> lk(mtx);
  cout << id <<" Waiting..... " << curr <<endl;
  cnd.wait(lk, [id]{return (curr==id);});
  
  if(id==1){
    f.first();
    //curr=2;
  }
  else if (id==2){
    //cnd.wait(lk, [id]{return (curr==id);});
    f.second();
    //curr=3;
  }
  else if (id==3){
    //cnd.wait(lk, [id]{return (curr==id);});
    f.third();
    //curr++;
    //return;
  }
  cout << "Finished " << id << "..."<< curr <<  endl; 
  //else
  //cout << "ERROR" << endl;
  curr++;
  cnd.notify_all();
  //lk.unlock();
  //return;
}

void notify()
{
  unique_lock<mutex> lk(mtx);
  curr=1;
  cnd.notify_all();
}

int main()
{
  Foo f;
  //itrations to test
  for (int i=0; i<100; ++i){
    thread t1=thread(c_func, 1, f);
    thread t2=thread(c_func, 2, f);
    thread t3=thread(c_func, 3, f);

    notify();

    t1.join();
    t2.join();
    t3.join();
  }
  
  return 1;
}
