/****
1. Write thread safe producer and consumer class.

class ProducerConsumer {

  // Define a queue with max capacity
  
  void put( const Object *obj );
  
  Object * get();
};
*****/

#include <iostream>
#include <mutex>
#include <condition_variable>
#include <queue>
#include <thread>
#include <future>
//#include <unistd.h>


#define MAX_SIZE 20

using namespace std;

class Object{
public:
  Object(int v): val(v)
  {};
  int getVal() const
  {
    return val;
  };
private:
  int val;
};


queue<const Object*> myQueue;  
mutex m;
condition_variable v;

class ProducerConsumer {
public:
  // Define a queue with max capacity
  static void put( const Object *obj );
  static const Object* get();
};

//Producer function
//void ProducerConsumer::put(const Object *obj)
void pput(const Object *obj)
{
  //define unque lock to wait on
  unique_lock<mutex> lck(m);

  //queue is full
  while(myQueue.size()>=MAX_SIZE)
    {
      //wait for queue to have space
      cout << "Waiting on consumer to dequeue obj..." << endl;
      v.wait(lck);
    }

  //queue is not full, insert object
  //m.lock();
  myQueue.push(obj);
  //m.unlock();

  //if first element after empty notify all threads
  if(myQueue.size()==1)
    v.notify_all();
  
  return;
}

//Consumer function
//const Object* ProducerConsumer::get()
const Object* cget()
{
  //define unque lock to wait on
  unique_lock<mutex> lck(m);

  //if Queue is empty wait
  while(myQueue.size()<=0)
    {
      //wait till element is produced
      cout << "Waiting on producer to enqueue obj..." << endl;
      v.wait(lck);
    }

  //queue is not empty acquire lock to consume
  //m.lock();
  
  const Object* result=myQueue.front();
  myQueue.pop();
  //m.unlock();

  //if Queue was full notify all producer threads
  if(myQueue.size()==MAX_SIZE-1)
    v.notify_all();

  return result;
  
}


int main()
{
  //ProducerConsumer pc;
  vector<thread> t_pool(50);

  /* Producer threads */
  for (int i=0; i<50; ++i)
    {
      const Object* temp = new Object(i);
      //t_pool[i]=thread(pc.put, &temp);
      t_pool[i]=thread(pput, temp);
    }

  for(int j=0; j<40; ++j)
  {
    auto f = async(launch::async, cget);
    const Object *res = f.get();
    cout << "value: " << res->getVal() << endl;
  }
  
  /***
  for(int j=20; j<50; ++j)
    {
      //auto tc = async(get);
      //Object *t;
      t_pool[j]=thread(get);
    }
  ****/

  for(int i=0; i<20; ++i)
    t_pool[i].join();


  /***/

}






