#include <iostream>
#include <string>
#include <cstdlib>


// Attribute types
#define BOOLEAN_ATTR                    0x1
#define ULONG_ATTR                      0x2
#define BYTESTR_ATTR                    0x3
#define ARRAY_ATTR                      0x4

static unsigned long pA, oA;

bool processArray(std::string str, unsigned int cnt);

int processLine(std::string lhs, std::string rhs)
{

   std::cout << "Received args LHS: "<< lhs << " RHS: " << rhs  << std::endl;
  //convert RHS to UL
  if(lhs == "UUID")
    {
      std:: string uuid(rhs);
      std::cout << uuid << std::endl;
    }

  if(lhs == "p11AttrType")
    {
      std::cout << "LHS: "<< lhs << std::endl << "RHS: " << rhs << " UL val: " << strtoul(rhs.c_str(), NULL, 16)  << std::endl;

      pA = strtoul(rhs.c_str(), NULL, 16);
      std::cout << "P11 Attribute " << pA << std::endl;
    }
  if (lhs == "osAttrType")
    {
      //convert to UL
      oA = strtoul(rhs.c_str(), NULL, 16);
      std::cout << "OS Attribute " << oA << std::endl;
    }
  if (lhs == "value")
    {
      //convert the vlaue into the appropriate stuct
      if (oA == BOOLEAN_ATTR)
	{
	  if(rhs == "true")
	    bool val = true;
	  else
	    bool val = false;
	  
	}
      else if (oA == ULONG_ATTR)
	{
	  unsigned long val = strtoul(rhs.c_str(), NULL, 16);
	  std::cout << "Val: " << val << std::endl;
	}
      else if (oA == BYTESTR_ATTR)
	{
	  std::cout << "Byte Stirng" << std::endl;
	}
      else if (oA == ARRAY_ATTR)
	{
	  std::cout << "Array with val: " << rhs.substr(1, rhs.length()-2)<< std::endl;
	  processArray(rhs.substr(1, rhs.length()-2), 6);
	}
      else
	{
	  std::cout << "Unknoen Attr Type: " << rhs << std::endl;
	  return 0;
	}
    }

  return 1;
}

bool processArray(std::string str, unsigned int cnt)
{	
	std::string delim = ",";
	std::string egal = "=";
	size_t start = 0;
	size_t mid = 0;
	size_t end = 0;

	if(cnt%3!=0)
		return false;

	// number of elements in the array (including Type, Kind)
	while (cnt!=0)
	{
		//1.
		mid = str.find(egal, start);
		end = str.find(delim, start);
		std::string aType = str.substr(mid+1, end-mid-1);		
		std::cout << "Attr Type " << aType << std::endl;
	 	start = end+1;	

		//2.
		mid = str.find(egal, start);
		end = str.find(delim, start);
		std::string aKind = str.substr(mid+1, end-mid-1);		
		std::cout << "Attr Kind " << aKind << std::endl;
	 	start = end+1;	


		//3.
		mid = str.find(egal, start);
		end = str.find(delim, start);
		std::string aVal = str.substr(mid+1, end-mid-1);		
		std::cout << "Attr Val " << aVal << std::endl;
	 	start = end+1;	

		cnt-=3;
	}		
	return true;
}

int main()
{
  //std::string obj_str = KVAgent::obj_cache[uuid];
  //std::string obj_str = "p11AttrType=0000000000000000;osAttrType=0000000000000002;value=0000000000000000;UUID=9e9e4ff3-2ce5-1d1e-6a23-b27559196b82;";
  std::string obj_str = "p11AttrType=0000000000000086;osAttrType=0000000000000004;value={attrType=286,attrKind=2,value=false,attrType=152,attrKind=3,attrVal=DEADBEEF,};"; 	
  std::string delim = ";";
  std::string egal = "=";

  size_t start = 0;
  size_t end = obj_str.find(delim) ;
  size_t mid = 0;
  
  while (end != std::string::npos)
    {
      std::string str_line = obj_str.substr(start, end - start) ;
      //std::cout << obj_str.substr(start, end - start) << std::endl;
      start = end + delim.length();
      end = obj_str.find(delim, start);

      //process str_line	
      //std::cout << str_line << std::endl;
      mid = str_line.find(egal);
      std::string lhs = str_line.substr(0,mid);
      std::string rhs = str_line.substr(mid+1,str_line.length());
      std::cout << "LHS: "<< lhs << std::endl << "RHS: " << rhs << std::endl;
      processLine(lhs, rhs);
      
    }
 /* 
  //std::cout << obj_str.substr(start, end) << std::endl;
  std::string last_line = obj_str.substr(start, end);
  mid = last_line.find(egal);
  std::string lhs = last_line.substr(0,mid);
  std::string rhs = last_line.substr(mid+1,last_line.length());
  processLine(lhs, rhs);
  std::cout << "LHS: "<< lhs << std::endl << "RHS: " << rhs << std::endl;
*/
  return true;
}


