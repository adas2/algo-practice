#include <iostream>
#include <unordered_map>
#include <list>
#include <set>
#include <string>


using namespace std;

//#define MAX_SIZE 10000
#define MAX_SIZE 4
#define SUCCESS true
#define FAILURE false 

//cache entry: index, value
typedef pair<string, string> EntryPair; 

class TestCache
{
public:
  static TestCache& getInstance();
  //insertEntry()
  bool insertEntry(string uuid, string val);
  //deleteEntry()
  bool deleteEntry(string uuid);
  bool readEntry(string uuid, string &buffer);
  bool modifyEntry(string uuid, string add);
  //OR writeEntry()
  //printCache()
  void printCache();
  //evict entry based on LRU
  bool evictEntry();

private:
  //initCache()
  TestCache(unsigned long max_size);
  //map of  index --> pointer for each entry
  static unordered_map<string, list<EntryPair>::const_iterator> cacheMap;
  //list of cache entries
  static list<EntryPair> cacheList; 
  //count of cache entries
  static unsigned long nEntries;
  //max size of cache
  static unsigned long  mSize; //in  number of entries
  
};

unsigned long TestCache::nEntries = 0;
unsigned long TestCache::mSize = MAX_SIZE;
list<EntryPair> TestCache::cacheList;
unordered_map<string, list<EntryPair>::const_iterator> TestCache::cacheMap;


//single object creation
TestCache& TestCache::getInstance()
{
  static TestCache tcache(MAX_SIZE);
  return tcache;

}


//constructor
TestCache::TestCache(unsigned long max_size) 
{
  mSize = max_size;
  cacheList.clear();
  cacheMap.clear();
  cout << "Cache Initialized" << endl;
}

//rewrite to check eviction condition first
//if full, evict, insert
//else insert
bool TestCache::insertEntry(string uuid, string data)
{
  //check if size has reached max
  if(nEntries >= mSize)
    {
      //cache is full you need to evict entries
      cout << "cache is full" << endl;
      //invoke LRU eviction policy
      if(!evictEntry())
      {
        return FAILURE;
      }
    }
  
  //check case: when an existing entry is inserted
  if(cacheMap.find(uuid)!=cacheMap.end())
    {
      cout << "Data Present and must be overwritten" << endl;
      return FAILURE;
    }
  //push the data to front of cList
  cacheList.push_front(make_pair(uuid, data));
  //create entry in the map
  //error case check if entry is already present?
  cacheMap[uuid] = cacheList.begin();
  //increment num entries
  nEntries++;
  return SUCCESS;
  
}

bool TestCache::readEntry(string uuid, string &buffer)
{
  //see if entry exists; failure if not
  //unordered_map<string,list<EntryPair>::const_iterator>::iterator mit;
  if(cacheMap.find(uuid)==cacheMap.end())
    return FAILURE;
  
  //bring entry to front of list ; splice() list to bring (*it) to front;
  list<EntryPair>::const_iterator it = cacheMap[uuid];
  cacheList.splice(cacheList.begin(),cacheList, it);
  //update map with new pointer
  cacheMap[uuid] = cacheList.begin();
  //copy contents in buffer
  buffer = cacheList.front().second;
  return SUCCESS;
}

//incremental change to entry, instead of full write
bool TestCache::modifyEntry(string uuid, string add)
{
  //check if uuid exists 
  if(cacheMap.find(uuid)==cacheMap.end())
    return FAILURE;
  //return FAUiLURE if no data found
  //bring entry to front of list
  list<EntryPair>::const_iterator it = cacheMap[uuid];
  cacheList.splice(cacheList.begin(),cacheList, it);
  //update entry in list with "add"
  cacheList.front().second += add;
  //update map with new pointer
  cacheMap[uuid] = cacheList.begin();
  return SUCCESS;
}


bool TestCache::deleteEntry(string uuid)
{
  //search index
  unordered_map<string,list<EntryPair>::const_iterator>::iterator it;
  it = cacheMap.find(uuid);
  if(it==cacheMap.end())
    return FAILURE;
  //delete item in list
  cacheList.erase(cacheMap[uuid]);
  //erase item in map
  cacheMap.erase(uuid);
  //reduce num entries
  nEntries--;
  return SUCCESS;
}

//LRU eviction
/***/
bool TestCache::evictEntry()
{
  //find last entry of list
  
  //delete Map entry
  cacheMap.erase(cacheList.back().first);
  //delete last element in list
  cacheList.pop_back();

  //reduce num entries
  nEntries--;
  return SUCCESS;
}
/****/

void TestCache::printCache()
{
  list<EntryPair>::iterator it;
  //print entries
  for (it = TestCache::cacheList.begin(); it!=TestCache::cacheList.end();it++)
    {
      cout << ">> Value: " << it->second << endl;
    }
  cout << endl;
  /****/
  unordered_map<string, list<EntryPair>::const_iterator>::iterator itm;
  for(itm = cacheMap.begin(); itm!=cacheMap.end(); itm++)
    {
      cout << "Index: " << itm->first << " --> Pointer: " << (itm->second)->second << endl;
    }
  /****/
  //pritn map?
  
}



int main()
{
  string uuid, val;
  bool rv;
  //cout << "Enter Sixe of cache" << endl;
  //unsigned long cSize;
  //cin >> cSize;
  //cout << "Initializing cache with max size " << cSize << endl; 
  TestCache *tc = &(TestCache::getInstance());
  cout << "Enter some values for cache" << endl;
  cin >> uuid;
  cin >> val;
  while(uuid!="end")
    { 
      rv = tc->insertEntry(uuid, val);
      cin >> uuid >> val;
    }
  cout << "Printing Cache Contents...." << endl;
  tc->printCache();
  cout << "Modify value for index: " << endl;
  string idx, newdata;
  cin >> idx >> newdata;
  tc->modifyEntry(idx, newdata);
  cout << "Printing Cache Contents...." << endl;
  tc->printCache(); 
  /***
  cout << "Delete some values for index" << endl;
  cin >> uuid;
  //cin >> val;
  rv = tc.deleteEntry(uuid);
  ***/
  //rv = tc.evictEntry();
  cout << "Read value for index: " << endl;
  string test;
  cin >> idx;
  if(tc->readEntry(idx, test))
    cout << "DATA: " << test << endl; 
  cout << "Printing Cache Contents...." << endl;
  tc->printCache(); 
  return 0;
}

