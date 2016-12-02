#include "list.h"

list_t * 
list_new(){
	list_t *self;
	self=malloc(sizeof(list_t));
	if(self==NULL){
		return self;
	}
	self->head = NULL;
	self->tail = NULL;
	self->len  = 0;
	return self;
}


list_node_t * 
list_push_tail( list_t *self, list_node_t *node){
	if (!node) return NULL;
	if (self->len){
		node->prev = self->tail;
		node->next = NULL;
		self->tail->next = node;
		self->tail = node;
	}
	else{
		self->head = self->tail = node;
		node->prev = node->next = NULL;
	}
	++self->len;
	return node;
};

list_node_t * 
list_push_head( list_t *self, list_node_t *node){
	if (!node) return NULL;
	if (self->len){
		node->next = self->head;
		node->prev = NULL;
		self->head->prev = node;
		self->head = node;
	}else{
		self->head = self->tail = node;
		node->prev = node->next = NULL;
	}
	++self->len;
	return node;
};

list_node_t *
list_pop_tail( list_t *self){
	if(!self->len) return NULL;
	list_node_t *node = self->tail;
	if(--self->len){
		self->tail = node->prev;
		node->prev=NULL;
	}else{
		self->head = self->tail = NULL;
	}
	return node;
};

list_node_t *
list_pop_head( list_t *self){
	if(!self->len) return NULL;
	list_node_t *node = self->head;
	if(--self->len){
		self->head = node->next;
		node->next=NULL;
	}else{
		self->head = self->tail = NULL;
	}
	return node;
};

list_node_t *
list_node_new(void *val) {
  list_node_t *self;
  if (!(self = malloc(sizeof(list_node_t))))return NULL;
  self->prev = NULL;
  self->next = NULL;
  self->val = val;
  return self;
}