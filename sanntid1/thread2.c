#include <pthread.h>
#include <stdio.h>

int i = 0;
pthread_mutex_t m1;

void* threadOne(){
	int x;
	for(x = 0 ; x <= 1000000 ; x++){
		pthread_mutex_lock(&m1);
		i ++;
		pthread_mutex_unlock(&m1);
	}
	return NULL;
}

void* threadTwo(){
	int y;
	for(y = 0 ; y <= 1000000 ; y++){
		pthread_mutex_lock(&m1);
		i--;
		pthread_mutex_unlock(&m1);
	}
	return NULL;
}

int main(){
	pthread_t thread_1;
	pthread_create(&thread_1, NULL, threadOne, NULL);
	
	pthread_t thread_2;
	pthread_create(&thread_2, NULL, threadTwo, NULL);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);
	printf("%d \n", i);

	return 0;
}
