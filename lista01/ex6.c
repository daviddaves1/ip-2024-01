#include <stdio.h>

int main(void) {
    int n;
    float fahrenheit;

    printf("Informe o numero de temparaturas em F a serem convertidas: ");
    scanf("%d", &n);

    for (int i=0; i<n; i++){
        printf("Informe a temperatura em F: ");
        scanf("%f", &fahrenheit);

        float celsius = (5.0 /9.0) * (fahrenheit - 32.0);

        printf("%.2f FAHRENHEIT EQUIVALEM A %.2f CELSIUS\n", fahrenheit, celsius);
    }

    return 0;
}
