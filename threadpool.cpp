// Smaple program to generate a thread pool
//   Each thread waits on an infinte loop for a job to come to the queue
// Use a condition variable to signal the waking of a thread  

#include <iostream>
#include <thread>
#include <condition_variable>
#include <mutex>

#define MAX_THREADS 4

using namespace std;
/**
class Threadpool {
public:
  //constructor
  Threadpool(int num_thrds);
  bool Initialize_Threads();

private:
  
};

***/
mutex m;

void foo(int id)
{
   m.lock();
   cout << "Thread id " << id << endl;
   m.unlock();  	
}



int main()
{
  //int num_jobs, num_thrds;
  //cout << "Enter size of the job" << endl;
  //cin >> num_jobs;
  //cout << "Enter size of the job" << endl;
  //cin >> num_thrds;

  //thread t1, t2;
  thread t1(foo, 1);
  thread t2(foo, 2);

  t1.join();
  t2.join();

}


