// Smaple program to generate a thread pool
//   Each thread waits on an infinte loop for a job to come to the queue
// Use a condition variable to signal the waking of a thread  

#include <iostream>
#include <thread>
#include <condition_variable>
#include <mutex>
#include <vector>
#include <unistd.h>

#define MAX_THREADS 4

using namespace std;

mutex chmtx;
condition_variable cnd;
//int global_var=0;
vector <int> chopstick(10,1);


void grab(int id)
{
  //chmtx.lock();
  int left=id;
  int right =(id>0)?(id-1):(chopstick.size()-1);
  unique_lock<mutex> lck(chmtx);
  
  //cout << "Waiitng... " << id << endl; 
  //cnd.wait(lck, [left, right]{return (chopstick[left] && chopstick[right]); });
  /***/
  while(chopstick[left]==0 || chopstick[right]==0)
    {
      cnd.wait(lck);
      cout << "Chopstick taken..." << id << " is waiting" << endl;
      //sleep and wait
      sleep(3);
    }

  /***/
  //if not taken
  chopstick[left]=0;
  chopstick[right]=0;
  cout << "Chopstick taken by " << id << endl; 
  
  //chmtx.unlock();  	
}

void release(int id)
{
  //chmtx.lock();
  int left=id;
  int right =(id>0)?(id-1):(chopstick.size()-1);
  unique_lock<mutex> lck(chmtx);
 
  chopstick[left]=1;
  chopstick[right]=1;
  cout << "Chopstick released by " << id << endl; 
  //notify all philosophers
  cnd.notify_all();

  //notify any one philosopher
  //cnd.notify_one();
  
  //chmtx.unlock();  	
}


void dine(int id)
{
  //grab
  grab(id);
  //eat()
  //release
  release(id);
  return;
}

int main()
{
  //define a thread vector
  vector<thread> philosophers(10);
  //vector<thread> thp2(10);
  
  //philosophers start dining one by one
  for(int i=0; i<10; ++i)
    {
      philosophers[i]= thread(dine, i);
      //philosophers[i]= thread(release, i);
    }
  
  for (auto& ph: philosophers)
    {
      ph.join();
    }
  
  return 0;
}


