#include <iostream>
#include <vector>

#define max(a,b)  (a>b)?a:b
#define min(a,b)  (a<b)?a:b

using namespace std;
/**
int getIndex(int idx)
{
  if (idx<0)
    return 
}
***/

//no wraparound for matrix
int getNeighborSum(vector<vector<int> > S, int i, int j)
{
  int row = S.size();
  int col = S[0].size();
  int y = min((j+2),col);
  int x = min((i+2), row);
  //cout << x;
  int sum=0;
  for (int m=max(i-1,0); m<(min((i+2),row)); ++m)
    for(int n=max(j-1,0); n<(min((j+2),col)); ++n)
      sum += S[m][n];
  
  return (sum - S[i][j]);

}
      
/**** With wrap-around
	  int sum = S[max(i+1,m-1)][j] + S[][j] 
	  +    S[i][(j+1)%col] + S[i][(j-1)<0?col-1:j-1]
	  +    S[(i+1)%row][(j+1)%col] + S[(i-1)<0?row-1:i-1][(j-1)<0?col-1:j-1]
	  +    S[(i+1)%row][(j-1)<0?col-1:j-1] + S[(i-1)<0?row-1:i-1][(j+1)%col]
	  ;
****/

  


//calculate state for number of iterations
void calcNextState(vector<vector<int> > &S1,vector<vector<int> > &S2, int &iter)
{
  if (iter==0 || S1.size()==0 )
    return;
  
  for (int i=0 ; i<S1.size(); ++i)
    for(int j=0; j<S1[0].size(); ++j)
      {
	//update according to rules
	if (S1[i][j]) // value is 1
	  {
	    //if more that 3 alive
	    if (getNeighborSum(S1,i,j)>3 || getNeighborSum(S1, i,j)<2)
	      S2[i][j]=0;
	    else //no change in state
	      S2[i][j]=S1[i][j];
	  }
	else //value is 0
	  {
	    if (getNeighborSum(S1,i,j)==3)
	      S2[i][j] = 1;
	    else //no change in state
	      S2[i][j]=S1[i][j];
	  }
      }
  --iter;
  //recursively call the next state, by reversing the matrix
  calcNextState(S2, S1, iter);
  
}


int main()
{
  int row =3, col=3;
  vector<vector<int> > S, nextS;
  S.resize(row); 
  nextS.resize(row);
  for (int i=0; i<row; ++i){
    S[i].resize(col, 0); 
    nextS[i].resize(col);
  }
  int iter = 2;

  //add some intial values values
  S[1][0]=1; S[1][2]=1, S[0][2]=1;
  //S[0][1]=1;

  //print
  for(int i=0; i<row; i++)
    {
      for(int j=0; j<col; ++j)
	cout << S[i][j] << " ";
      cout << endl;
    }

  cout << "========" << endl;
  calcNextState(S, nextS, iter);

  

  //print
  for(int i=0; i<row; i++)
    {
      for(int j=0; j<col; ++j)
	cout << nextS[i][j] << " ";
      cout << endl;
    }

}
