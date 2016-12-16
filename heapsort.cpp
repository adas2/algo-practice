// range heap example
#include <iostream>     // std::cout
#include <algorithm>    // std::make_heap, std::pop_heap, std::push_heap, std::sort_heap
#include <vector>       // std::vector

int main () {
  int myints[] = {10,20,30,5,15};
  std::vector<int> v(myints,myints+5);

  std::cout << "Initial array :";
  for (unsigned i=0; i<v.size(); i++)
    std::cout << ' ' << v[i];
  std::cout << std::endl;

  std::make_heap (v.begin(),v.end(), [](int a, int b){return a>b;});
  std::cout << "initial min heap   : " << v.front() << '\n';

  std::pop_heap (v.begin(),v.end(), [](int a, int b){return a>b;}); v.pop_back();
  std::cout << "min heap after pop : " << v.front() << '\n';

  v.push_back(2); std::push_heap (v.begin(),v.end(), [](int a, int b){return a>b;});
  std::cout << "min heap after push: " << v.front() << '\n';

  std::sort_heap (v.begin(),v.end(), [](int a, int b){return a>b;}); 

  std::cout << "final sorted range :";
  for (unsigned i=0; i<v.size(); i++)
    std::cout << ' ' << v[i];

  std::cout << '\n';

  return 0;
}
