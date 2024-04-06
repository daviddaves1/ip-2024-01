#include <stdio.h>

int main(void) {
    float salar;

    printf("Informe o salario do funcionario: ");
    scanf("%f", &salar);


    float salar_reajust;
    if (salar <= 300.00) {
        salar_reajust = salar * 1.50;
    } else {
        salar_reajust = salar * 1.30;
    }

    printf("SALARIO COM REAJUSTE = %.2f\n", salar_reajust);

    return 0;
}
