#include <iostream>
#include <vector>
#include <stdlib.h>
#include <string.h>

using namespace std;

template<class T, class Y>
class A{
public:
  A(T n, Y r);
  T triple();
  Y value();

private:
  //~A() {};
  T num;
  Y ratio;
};

template<class T>
class A<T, T*>{
public:
  vector<T> values;
  void sort_numbers(vector<int> v);
  bool isempty() {return true;}
  int triple();
};

template<>
class A<bool, char>{
public:
  char* triple(bool p);
};


template<class T, class Y> 
A<T,Y>::A(T n, Y r) : num(n),ratio(r)
{
}
template<class T, class Y> 
T A<T,Y>::triple()
{
  return 3*num;
}
template<class T, class Y> 
Y A<T,Y>::value()
{
  return num*ratio;
}


char* A<bool, char>::triple(bool flag)
{
  char *t;
  if(flag)
    {
      t = (char*)calloc(10, sizeof(char));
      strcpy(t, "TRIPLE");
    }
  return t;
}


int main()
{
  int n;
  double r; 
  cout << "Enter an integer and ratio" << endl;
  cin >> n >> r;
  A<int, double> a(n,r);
  //a.num = n;
  cout << "Output: " << a.triple() << ", " << a.value() << endl;
  A<int, int*> b;
  if(b.isempty())
    cout << "EMPTY" << endl;
  A<bool, char> c;
  cout << c.triple(true) << endl;

}
