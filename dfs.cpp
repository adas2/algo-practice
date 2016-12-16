#include <iostream>
//#include <cstdbool>
#include <list>
#include <queue>

#define MAX_NODES 10

using namespace std;

typedef struct gnode{
  int adj_v;
  int weight;
  //struct gnode *next;
}g_node;


class Graph{
public:
  //add edge from node u to node v 
  void addEdge(int u, int v);
  //initialize()
  void initGraph(int num_v, const bool gdirected);
  //start DFS from node v
  void DFSinit();
  void DFSvisit(int v);
  //start BFS from v
  void BFSvisit(int v);
  void printGraph();
  int numNodes();
  int numEdges();
  void printParents();
  void topSort();


private:
  int V; //num nodes
  int E; // edges
  int parent[MAX_NODES]; //parent nodes
  list<g_node> edges[MAX_NODES];
  bool directed;
  int degree[MAX_NODES];
  int in_degree[MAX_NODES];
  bool visited[MAX_NODES];
  bool processed[MAX_NODES];
  //stack used for DFS
  list<int> st;
  
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
void Graph::addEdge(int x, int y)
{
  //g_node p = (g_node*)malloc(sizeof(g_node));
  //p->
  g_node t;
  t.adj_v = y;
  t.weight = 0;
  edges[x].push_front(t);
  degree[x]++;
  in_degree[y]++;
  
  if (!directed)
    {
      //addEdge(y, x);
      g_node q;
      q.adj_v = x;
      q.weight = 0;
      edges[y].push_front(q);
      degree[y]++;
      //E++;
    }
  
  E++;
  
  
}


//print grapg as read
void Graph::printGraph()
{
  //for each vertex
  for (int i = 0; i< V; i++)
    {
      cout << "Node: " << i << " Degree: "<< degree[i] << endl;
      if(!edges[i].empty())
	{
	  cout << "Edges: " << endl;
	  list<g_node>::iterator it;
	  for(it=edges[i].begin(); it != edges[i].end(); it++)
	    {
	      cout << (*it).adj_v << " " ;
	    }
	  cout << endl;
	}
    }
  if (directed)
    cout << "Graph is directed!" << endl;
  else
    cout << "Undirected Graph! " << endl;
}

void Graph::DFSinit()
{
  st.clear();
  for(int i=0; i<V; i++)
    {
      visited[i]=false;
      processed[i]=false;
      parent[i]=-1; // initialize parents to -1 
    }
}


//Depth First Search: recursive
void Graph::DFSvisit(int s)
{
  //insert first node in stack
  //st.push_front(s);
  visited[s] = true;
  //process(s)

  //while(!st.empty())
  //{
  //int x=st.front();
  cout << "Strated processing vertex: " << s << endl;
  list<g_node>::iterator it;
  for(it=edges[s].begin(); it!=edges[s].end(); it++)
    {
      int y = (*it).adj_v;
      if(!visited[y])
	{
	  //st.push_front(y);
	  parent[y]=s;
	  cout << "Process edge: " << s << y << endl; 
	  //visited[y] = true;
	  DFSvisit(y);
	}
      else if (!processed[y])
	{
	  //this is a back edge visit
	  //process edge s->y
	  cout << "Process edge: " << s << y << endl; 
	  if(parent[y]!=s)
	    cout << "Back edge causing cycle: " << s << y << endl;
	}
	  
    }
  processed[s]=true;
  //st.pop_front();
  cout << "Finished processing vertex: " << s << endl;
  //}
  
  
}

//Bread First Search: iterative
void Graph::BFSvisit(int s)
{
  //initialize stakc, visited, processed, same as DFS
  DFSinit();

  //insert first node in stack
  st.push_front(s);
  visited[s] = true;
  //process(s)

  while(!st.empty())
    {
      int x=st.front();
      cout << "Strated processing vertex: " << x << endl;
      list<g_node>::iterator it;
      for(it=edges[x].begin(); it!=edges[x].end(); it++)
	{
	  int y = (*it).adj_v;
	  if(!visited[y])
	    {
	      //queue implementation, always push back
	      st.push_back(y);
	      parent[y]=x;
	      cout << "Process edge: " << x << y << endl; 
	      visited[y] = true;
	      //DFSvisit(y);
	    }
	  else if (!processed[y])
	    {
	      //disocvered but not processed
	      cout << "Process edge: " << x << y << endl; 
	
	    }
	}
      processed[x]=true;
      st.pop_front();
      cout << "Finished processing vertex: " << x << endl;
    }
  
  
}


void Graph::printParents()
{
  cout << "Parent Graph:" << endl;
  for (int j=0; j<V; j++)
    cout << parent[j] << " --> " << j << endl;
}

//topological sort of Directed Graph
void Graph::topSort()
{
  cout << "Printing top sort" << endl;
  
  //if undirected return
  if(!directed)
    {
      cout << "Undirected graph!!" << endl;
      return;
    }
  //maintain queue of vertices
  list<int> tsq;
  tsq.clear();
  int count = 0;
  
  // push vertices with in degree 0
  for (int i=0; i<V; i++)
    {
      if (in_degree[i] == 0)
	{
	  tsq.push_back(i);
	}
    }

  if(tsq.empty())
    {
      cout << "Graph has cycles!" << endl;
      return;
    }

  //traverse all vertices
  while(!tsq.empty())
    {
      int temp = tsq.front();
      cout << tsq.front() << " , ";
      tsq.pop_front();
      //modify in-degree for adj edges in-degree array
      list<g_node>::iterator it;
      for(it = edges[temp].begin(); it!=edges[temp].end(); it++)
	{
	  if((--in_degree[(*it).adj_v])==0)
	    {
	      tsq.push_back((*it).adj_v);
	    }
	  //if()
	}
    }
  
  cout << endl;

}

int main()
{
  cout << "Enter the num vertices and num edges" << endl;
  int num_v, num_e;
  cin >> num_v >> num_e;
  int x,y;
  //cin >> x >> y;
  Graph G;
  const bool g_directed = true;
  G.initGraph(num_v, g_directed);
  cout << "Enter graph by its edges " << endl;
  int i=1;
  while (i<=num_e)
    {
      cin >> x >> y;
      G.addEdge(x,y);
      i++;
    }

  cout << "Graph Initialized with " << G.numNodes() << " nodes and " << G.numEdges() << " edges" << endl;
 
  
  G.printGraph();
  /*******
  cout << "DFS traversal of graph:" << endl;
  G.DFSinit();
  G.DFSvisit(0);
  G.printParents();

  G.BFSvisit(0);
  G.printParents();
  ****/

  G.topSort(); 
  
}

