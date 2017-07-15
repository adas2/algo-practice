/****
Dijkstra's shortest path sample code
*****/
#include <climits>
#include <iostream>
#include <list>
#include <vector>

#define MAX_NODES 10
#define min(a,b)  (a<b?a:b)

using namespace std;


typedef struct gnode{
  int adj_v;
  int weight;
  //struct gnode *next;
}g_node;


class Graph{
public:
  //add edge from node u to node v
  void addEdge(int u, int v, int wght);
  //initialize()
  void initGraph(int num_v, const bool gdirected);
  void printGraph();
  int numNodes();
  int numEdges();
  void printParents();
  //calculate shorted path distances from src to all vertices
  vector<int> shortestPathDist(int src);

private:
  int V; //num nodes
  int E; // edges
  int parent[MAX_NODES]; //parent nodes
  list<g_node> edges[MAX_NODES];
  bool directed;
  int degree[MAX_NODES];
  int in_degree[MAX_NODES];
};

//initialize this graph
void Graph::initGraph(int num_v, const bool gdirected)
{
  V=num_v;
  E=0;
  directed = gdirected;
  for(int i=0; i<MAX_NODES; i++)
    {
      //parent[i] = -1;
      degree[i] = 0; //undirected case
      in_degree[i] = 0; //directed case
      edges[i].clear();
    }

}

int Graph::numNodes()
{
  return V;
}

int Graph::numEdges()
{
  return E;
}

//manually add each edge from user
void Graph::addEdge(int x, int y, int w)
{
  //g_node p = (g_node*)malloc(sizeof(g_node));
  //p->
  g_node t;
  t.adj_v = y;
  t.weight = w;
  edges[x].push_front(t);
  degree[x]++;
  in_degree[y]++;

  if (!directed)
    {
      //addEdge(y, x);
      g_node q;
      q.adj_v = x;
      q.weight = w;
      edges[y].push_front(q);
      degree[y]++;
      //E++;
    }

  E++;

}

//print grapg as read
void Graph::printGraph()
{
  if (directed)
    cout << "\n\nDirected Graph:" << endl;
  else
    cout << "\n\nUndirected Graph:" << endl;

  //for each vertex
  for (int i = 0; i< V; i++)
    {
      cout << "Node: " << i;
      //cout << " In-Degree="<< in_degree[i] << " Out-Degree=" << degree[i] << endl;
      if(!edges[i].empty())
        {
          cout << " Edges -->" ;
          list<g_node>::iterator it;
          for(it=edges[i].begin(); it != edges[i].end(); it++)
            {
              cout << " ("  << (*it).adj_v ;
              cout << ", " << (*it).weight << ")";
            }
        }
        cout << endl;
    }

}

//Dijkstra first cut : not optimized
vector<int> Graph::shortestPathDist(int src)
{
  //vector to idicate vertices included in shortest path, initialized 
  vector<bool> sptSet(this->numNodes(), false);
  //vector to store shortest path dist for each, init to infinity
  vector<int> sDist(this->numNodes(), INT_MAX); 
  //count for nodes added to sptSet
  int cnt =0;

  //set the src dist to zero
  
  sDist[src]=0;
  //++cnt;

  //while all nodes have not been discovered and added to sptSet 
  while(cnt < this->numNodes())
    {
      //choose the vertex with min dist
      int min_dist=INT_MAX, min_v=-1;
      for(int i=0; i<this->numNodes(); ++i)
	{
	  //if vertex not added and has value less that min
	  if(!sptSet[i] && min_dist>sDist[i])
	    {
	      min_dist=sDist[i];
	      min_v=i;
	    }
	}//min v finding complete 
      cout << "min_v: " << min_v << " min_dist: " << min_dist << endl;
 
      //add min i to the sptSet, increment cnt
      sptSet[min_v]=true;
      ++cnt;

      //update the adjacent vertices for min_v
      list<g_node>::iterator it;
      for(it=edges[min_v].begin(); it!=edges[min_v].end(); ++it)
	{
	  //if the value is sDist is greate than min_v+weight update
	  if(sDist[(*it).adj_v] > sDist[min_v]+(*it).weight)
	    {
	      sDist[(*it).adj_v]=sDist[min_v]+(*it).weight;
	    }
	}//update complete
    }//while loop complete
  
  return sDist;
}

int main()
{
  int num_v=9;
  Graph G;
  //undirected
  const bool g_directed = false;
  G.initGraph(num_v, g_directed);
  G.addEdge(0,1, 4);
  G.addEdge(0,7, 8);
  G.addEdge(1,7, 11);
  G.addEdge(1,2, 8);
  G.addEdge(2,8,2);
  G.addEdge(7,8,7);
  G.addEdge(7,6,1);
  G.addEdge(8,6,6);
  G.addEdge(2,5,4);
  G.addEdge(2,3,7);
  G.addEdge(6,5,2);
  G.addEdge(3,5,14);
  G.addEdge(3,4,9);
  G.addEdge(5,4,10);

  /*
  G.addEdge(5,0, 0);
  G.addEdge(5,2, 0);
  G.addEdge(4,1, 0);
  G.addEdge(4,0, 0);
  G.addEdge(2,3, 0);
  G.addEdge(3,1, 0);
  */

  G.printGraph();

  vector<int> sPDist=G.shortestPathDist(0);
  for(int x=0; x<num_v; ++x )
    cout << x <<":" << sPDist[x] << " ";
  cout << endl;
}


