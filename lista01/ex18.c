#include <stdio.h>

int main(void) {
    int a1, r, n;
    long int soma=0;

    printf("Informe o valor inicial da progressao , a razao e o numero de elementos: ");
    scanf("%d%d%d", &a1, &r, &n);

    for (int i=0; i<n; i++) {
        soma += a1+i*r;
    }

    printf("%ld\n", soma);

    return 0;
}
