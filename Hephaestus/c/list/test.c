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
	char  *f= "f";
	list_push(ls,a);
	list_push(ls,b);
	list_push(ls,c);
	list_push(ls,d);
	list_push(ls,f);
	printf("index 0 is a =%s\n",list_index(ls,0));
	printf("index 1 is b =%s\n",list_index(ls,1));
	printf("index 2 is c =%s\n",list_index(ls,2));
	printf("index 3 is d =%s\n",list_index(ls,3));
	printf("index 4 is f =%s\n",list_index(ls,4));
	printf("index 5 is null =%s\n",list_index(ls,5));
	void * tmp_node;
	while(1){
		printf("%d\n",list_len(ls));
		tmp_node=list_pop(ls);
		if(!tmp_node) break ;
		printf("%s\n",tmp_node);
	}
	return 0;
}
