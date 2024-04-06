#include <stdio.h>

int main(void) {
    float cat_popular = 1, cat_geral = 5, cat_arquib = 10, cat_cadeiras = 20;
    int casos_testes, qtde_pessoas, popular, geral, arquib, cadeiras;

    printf("Informe o número de casos de testes: ");
    scanf("%d", &casos_testes);

    for(int i = 1; i <= casos_testes; i++) {
        float perc_popular = 0, perc_geral = 0, perc_arquib = 0, perc_cadeiras = 0;

        printf("Informe os seguintes dados abaixo para o caso de teste %d:\n\n", i);

        printf("Quantidade de pessoas que compraram ingressos: ");
        scanf("%d", &qtde_pessoas);
        printf("\n");

        printf("Percentual da categoria popular: ");
        scanf("%f", &perc_popular);
        printf("\n");
        
        printf("Pecentual da categoria geral: ");
        scanf("%f", &perc_geral);
        printf("\n");

        printf("Percentual da categoria arquibancada: ");
        scanf("%f", &perc_arquib);
        printf("\n");

        printf("Percentual da categoria cadeiras: ");
        scanf("%f", &perc_cadeiras);
        printf("\n");

        popular = qtde_pessoas * (perc_popular / 100);
        geral = qtde_pessoas * (perc_geral / 100);
        arquib = qtde_pessoas * (perc_arquib / 100);
        cadeiras = qtde_pessoas * (perc_cadeiras / 100);

        float renda_jogo = (popular * cat_popular) + (geral * cat_geral) + (arquib * cat_arquib) + (cadeiras * cat_cadeiras);

        printf("A renda do jogo N. %d é R$ %.2f\n\n\n\n", i, renda_jogo);
    }

    return 0;
}
