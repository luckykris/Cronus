#include "list.h"
#include <stdio.h>
int main(int argc, char const *argv[])
{
	list_t *ls;
	ls = list_new();
	list_node_t *a = list_node_new("a");
	list_node_t *b = list_node_new("b");
	list_node_t *c = list_node_new("c");
	list_node_t *d = list_node_new("d");
	list_push_tail(ls,a);
	list_push_tail(ls,b);
	list_push_head(ls,c);
	list_push_head(ls,d);
	list_node_t * tmp_node;
	while(1){
		tmp_node=list_pop_head(ls);
		if(!tmp_node) break ;
		printf("%s\n",tmp_node->val);
	}
	return 0;
}