//Buy mcnugget in packs of a1, a2, a3 sizes, can you buy N?
#include <iostream>
#include <vector>

using namespace std;

bool is_mcnugget_number(int N, int a, int b, int c)
{
  //bool mNum[N];
  vector<bool> mNum(N+1, false);
  mNum[a]= true;
  mNum[b]= true;
  mNum[c]= true;

  //find 4
  int min;
  if(b<a && b<c)
    min=b;
  else if (c<a && b>a)
    min=c;
  else
    min=a;
  cout << min << endl;
  int cnt=0;
  int i=1;
  //cnt till min consequtive entries are true, or N is reached
  while(cnt!=min && i<=N)
    {
      cout << "Before i: " << i << " Val: " << mNum[i] << endl;
      mNum[i] = mNum[i]|| //important, if entry is true keep it 
	((i>=a)&&(mNum[i-a]))||
	((i>=b)&&(mNum[i-b]))||
	((i>=c)&&(mNum[i-c]));
      cout << "i: " << i << " Val: " << mNum[i] << endl;
      if(mNum[i]) //count consecutive 1's
	++cnt;
      else //entry is zero reset
	cnt=0;
      ++i;
    }

  if(cnt==min)
    {
      //largest non-McNugget encountered already#
      cout << "Larget non attainable num: " << i-cnt-1 << endl;
      return true;
    }
  return mNum[N];
  
}

int main()
{
  int N=150;
  if(is_mcnugget_number(N,6,9,20))
    cout << "True" << endl;
  else
    cout << "False" << endl;
  return 1;
}


