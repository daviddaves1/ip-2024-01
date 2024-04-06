#include <stdio.h>

int main(void) {
    float fahrenheit, poleg;

    printf("Informe a temperatura em F: ");
    scanf("%f", &fahrenheit);
    printf("Informe a quantidade de chuva em polegada: ");
    scanf("%f", &poleg);

    float celsius = (5 * fahrenheit - 160) / 9;
    float milim = poleg * 25.4;

    printf("O VALOR EM CELSIUS = %.2f\n", celsius);
    printf("A QUANTIDADE DE CHUVA E = %.2f\n", milim);


}
