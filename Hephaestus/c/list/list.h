#ifndef LIST_H
#define LIST_H


#ifdef __cplusplus//cpp support
extern "C" {
#endif

#define LIST_VERSION "0.0.1"
#include <stdlib.h>


typedef enum {
	LIST_HEAD,
	LIST_TAIL
} list_direction_t;

typedef struct list_node{
	struct list_node *prev;
	struct list_node *next;
	void *val;
} list_node_t;

typedef struct{
	list_node_t *head;
	list_node_t *tail;
	unsigned int len;
} list_t;

list_t * 
list_new();

//list_node_t * 
void
list_push( list_t *self, void *val);

void
list_push_head( list_t *self, void *val);

void *
list_pop( list_t *self);

void *
list_pop_head( list_t *self);

list_node_t *
list_node_new(void *val);

unsigned int
list_len(list_t *self);

void *
list_index(list_t *self,unsigned int index);

#ifdef __cplusplus
}
#endif
#endif