#include <iostream>
#include <vector>


using namespace std;

//given denomination vector and total value
int findCoinChangeWays( vector<int> denom, int value )
{
	int n_coins = denom.size();

	//2D matrix to store the num ways for j value uisng i coins
	vector<vector<int> > W;
	W.resize(n_coins+1); //extra size so that 1 to n can be indexed
	for(int i=0;i<=n_coins; ++i)
		W[i].resize(value+1,0); //all value initialized to 0
	//if the value is zero then take 0 coins of each type to make it; hence 1 way
	for(int i=0; i<=n_coins; ++i)
		W[i][0]=1;

	//Using DP solution with (i) Optimal Sunstructure and (ii) Overlapping subproblems
	//Let W[i][j] --> max num ways in which with first i coins value j can be constructed
	//	W[i][j] =  i.  W[i-1][j] //i.e. without the ith coin at all
	//		 +  ii. W[i][j-denom[i-1]] //i.e. with ith coin used x number of times 
	// W[i][j] =  (i. + ii.);

	//Traverse the matrix and fill the optimal ways
	for(int i=1; i<=n_coins; ++i)
	{
		for(int j=1; j<=value; ++j)
		{
			//case i. coin i is not used; 
			W[i][j] = W[i-1][j];
			//case ii. coin i is used, i.e. coin with value denom[i-1], if j-denom[i-1] is negative, then zero ways
			W[i][j] += ((j-denom[i-1])<0?0:W[i][j-denom[i-1]]);

		}

	}

	//return the right most corner ofmatrix
	return W[n_coins][value];

}

int main()
{
	vector<int> coins={2,5,3,6};
	int V = 10;
	cout << "Number of ways value " << V << " can be generated: " << findCoinChangeWays(coins, V) << endl;
	return 1;
}
