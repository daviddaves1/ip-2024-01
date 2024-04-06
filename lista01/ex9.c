#include <stdio.h>
#include <math.h>

int Bhaskara(float a, float b, float c, float *r1, float *r2) {
    float delta;

    delta = b*b-4*a*c;

    if (delta > 0) {
        *r1 = (-b + sqrt(delta)) / (2*a);
        *r2 = (-b - sqrt(delta)) / (2*a);
        return 2;
    } else if (delta == 0) {
        *r1 = -b /(2*a);
        return 1;
    } else {
        return 0;
    }

    return delta;
}

int main(void) {
    float a=0, b=0, c=0, r1=0, r2=0;

    printf("\n\n");
    printf("Digite os valores de A, B e C (0 0 0): ");
    scanf("%f%f%f", &a, &b, &c);

    int resul = Bhaskara(a, b, c, &r1, &r2);

    printf("O VALOR DE DELTA E = %.2f\n", (b*b-4*a*c));

    if (resul == 2) {
        printf("A equacao possui 2 raizes reais distintas --> Raiz 1: %.2f     Raiz 2: %.2f\n", r1, r2);
    } else if (resul == 1){
        printf("A equacao possui 2 raizes reais iguais --> Raiz 1: %.2f     Raiz 2: %.2f\n", r1, r1);
    } else {
        printf("Nao existe raiz real nessa equacao.\n");
    }

    return 0;
}
