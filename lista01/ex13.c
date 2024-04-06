#include <stdio.h>

int main(void) {
    float n=0;
    char c;

    printf("Diga a nota: ");
    scanf("%f", &n);

    if (n >= 9.0){
        c = 'A';
    }
    else if (n >= 7.5){
        c = 'B';
    }
    else if (n >= 6.0){
        c = 'C';
    }
    else if(n >= 0){
        c = 'D';
    }

    printf("NOTA = %.2f CONCEITO = %c\n", n, c);

    return 0;
}
