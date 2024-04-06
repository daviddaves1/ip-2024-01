#include <stdio.h>

int main(void) {
    int horas, minutos, segundos;

    printf("Informe as horas, minutos e segundos: ");
    scanf("%d%d%d", &horas, &minutos, &segundos);

    int tempo_segun = horas * 3600 + minutos * 60 + segundos;

    printf("O TEMPO EM SEGUNDOS E = %d\n", tempo_segun);

    return 0;
}
