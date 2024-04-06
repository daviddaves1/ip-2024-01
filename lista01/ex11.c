#include <stdio.h>

int main(void) {      
    int number;

    printf("Digite o numero: ");
    scanf("%d", &number);

    if (number % 3 != 0 && number % 5 != 0) {
        printf("O NUMERO NAO E DIVISIVEL\n");
    } else {
        printf("O NUMERO E DIVISIVEL\n");
    }
}
