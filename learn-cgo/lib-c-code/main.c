#include <stdio.h>
#include "main.h"
void hello_world(int a) {
    printf("hello world %d\n", a);
}

/*
gcc -c lib-c-code/main.c
gcc -shared -o main.so main.o
*/
int main(int argc, char** argv){
    hello_world(42);
}