#include <iostream>
#include <algorithm>
#include <stack>
#include <vector>

using namespace std;

class Meeting
{
private:
        // number of 30 min blocks past 9:00 am
        unsigned int startTime_;
        unsigned int endTime_;

public:
  Meeting() :
    startTime_(0),
    endTime_(0)
  {
  }
  
  Meeting(unsigned int startTime, unsigned int endTime) :
    startTime_(startTime),
    endTime_(endTime)
  {
  }

  unsigned int getStartTime() const
  {
    return startTime_;
  }
  
  void setStartTime(unsigned int startTime)
  {
    startTime_ = startTime;
  }
  
  unsigned int getEndTime() const
  {
    return endTime_;
  }
  
  void setEndTime(unsigned int endTime)
  {
    endTime_ = endTime;
  }
  
  bool operator==(const Meeting& other) const
  {
    return
      startTime_ == other.startTime_
      && endTime_ == other.endTime_;
    }
  };


//compare the calendar event times for sorting
bool myCompare(pair<int,bool> a, pair<int,bool> b)
{
  if (a.first < b.first)
    return true;
  else if (a.first==b.first && a.second)
    return true;
  else 
    return false;
}

//given a list of meetings find the times when everybody is free
vector<pair<int,int>> findFreeSlots(vector<Meeting> meetings)
{
  vector<pair<int, bool> > events; 

  //Read all create data structure event from start and end times 
  for(auto& x: meetings)
    {
      int s = x.getStartTime();
      int e = x.getEndTime();
      //if event is start mark as true, false otherwise
      events.push_back(make_pair(s,true));
      events.push_back(make_pair(e,false));
    }

  //sort the events array such that start times always come before end for the same time stamp
  sort(events.begin(), events.end(), myCompare);

  //Use a count variable to track ongoing events, increment when something starts and decrement when something ends
  int count=0, prev_count=0;
  int s_time=0, e_time=0;

  //Array to hold the free slots with start time, end time
  vector<pair<int,int>> result;

  //iterate through the array
  vector<pair<int, bool> >::const_iterator it;
  for(it=events.begin(); it!=events.end(); ++it)
    {
      if((*it).first)
	{
	  count++;
	}
      else
	{
	  count--;
	}

      //first time counter moves from 0->1 (indicates there was no conflict before this)
      //only exception is first meeting, if it indeed has a conflict, s_time will get rewritten 
      if(count==1 && prev_count==0)
	{
	  //mark start time
	  s_time=(*it).first;
	  cout << "Start Time: " << s_time << endl;
          //result.push_back(make_pair(s_time,true));
     	}
      //if the counter moves from 1->0 implies meeting with no conflicts ended
      if(count==0)
	{
	  //mark end time
	  e_time=(*it).first;
	  cout << "End Time: " << e_time << endl;
	  //end of a event push to result array
	  //Meeting m(s_time, e_time);
          
	  result.push_back(make_pair(s_time, e_time));
	}

      prev_count=count;

    }
  return result;
}


int main()
{
  //Test These 2 cases
  //[ Meeting(0, 1), Meeting(3, 5), Meeting(4, 8), Meeting(10, 12), Meeting(9, 10) ]
  //  [ Meeting(1, 10), Meeting(2, 6), Meeting(3, 5), Meeting(7, 9) ]
    vector<Meeting> test1={ Meeting(0, 1), Meeting(3, 5), Meeting(4, 8), Meeting(10, 12), Meeting(9, 10)  };
    vector<pair<int,int>> result;
    
    result=findFreeSlots(test1);
    /***/
    for(auto x: result)
    {
        cout << x.first << " to " << x.second << endl;
        
    }
     /**/ 
  return 0;
}
