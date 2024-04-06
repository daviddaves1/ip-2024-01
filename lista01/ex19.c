#include <stdio.h>

int main(void) {
    int n=0;
    double soma=0;

    do{
        printf("Informe o numero: ");
        scanf("%d", &n);
        printf("Numero invalido!\n");
    }while(n <= 1);

    for (int k = 1; k <= n; k++){
        soma += 1.0 / k;
    }

    printf("SAIDA: %.6lf\n", soma);

    return 0;
}
