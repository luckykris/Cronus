#include "list.h"
#include <stdio.h>
int main(int argc, char const *argv[])
{
	list_t *ls;
	ls = list_new();
	char *a= "a";
	char *b= "b";
	char *c= "c";
	char *d= "d";
	int   e= 1;
	int  *f= &e;
	list_push(ls,a);
	list_push(ls,b);
	list_push(ls,c);
	list_push(ls,d);
	list_push(ls,f);
	void * tmp_node;
	while(1){
		printf("%d\n",list_len(ls));
		tmp_node=list_pop(ls);
		if(!tmp_node) break ;
		printf("%s\n",tmp_node);
	}
	return 0;
}
