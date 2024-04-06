#include <stdio.h>
#include <math.h>

int main(void) {
    int n1=0, n2=0, n3=0;

    printf("Digite os tres numeros inteiros separados por espaco: ");
    scanf("%d%d%d", &n1, &n2, &n3);

    if (n1 > 10 && n1 <= 0 && n2 > 10 && n2 <= 0 && n3 > 10 && n3 <= 0){

        printf("DIGITO INVALIDO\n");
    } else{

        int result = n1 * 100 + n2 * 10 + n3;
        int quadrado = pow(result, 2);
        printf("Concatenacao: %d,           Quadrado: %d\n", result, quadrado);
    }

    return 0;
}
