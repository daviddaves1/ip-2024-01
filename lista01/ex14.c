#include <stdio.h>
#include <math.h>

int main(void){
    float alt, arest, vol;
    const float raiz3sobre2 = sqrt(3) / 2;

    printf("Informe a altura e comprimento da aresta: ");
    scanf("%f%f", &alt, &arest);

    float area_base = 3 * arest * arest * raiz3sobre2;
    vol = (1/3) * area_base * alt;

    printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", vol);
}
