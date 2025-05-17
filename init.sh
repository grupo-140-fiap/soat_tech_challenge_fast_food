#!/bin/bash

API_URL="http://localhost:8080/api/v1/products"
HEADER="Content-Type: application/json"

declare -A ITENS=(
    ["Sanduíche Natural"]="Pão integral, peito de peru, alface, tomate e maionese de ervas|12.50|burger"
    ["Coxinha de Frango"]="Tradicional salgadinho de frango com massa crocante|8.00|burger"
    ["Pão de Queijo"]="Delicioso pão de queijo mineiro, quentinho e macio|5.00|burger"
    ["Empadão de Frango"]="Massa amanteigada recheada com frango desfiado e temperado|15.00|burger"
    ["Wrap de Frango"]="Tortilha recheada com frango grelhado, alface, cenoura ralada e molho de iogurte|18.00|burger"
    ["Hambúrguer Artesanal"]="Pão brioche, carne artesanal, queijo, alface, tomate e molho especial|22.00|burger"
    ["Pastel de Carne"]="Massa crocante recheada com carne moída bem temperada|10.00|burger"
    ["Esfirra de Queijo"]="Massa leve com recheio de queijo derretido|7.00|burger"
    ["Tapioca com Queijo e Presunto"]="Tapioca recheada com queijo derretido e presunto|12.00|burger"
    ["Açaí na Tigela"]="Açaí batido com frutas e granola para uma opção mais saudável|20.00|burger"

    ["Batata Frita"]="Crocante por fora e macia por dentro, com opção de molhos como ketchup e maionese|25.50|side"
    ["Onion Rings"]="Anéis de cebola empanados e fritos, crocantes e saborosos|35.90|side"
    ["Salada Coleslaw"]="Mistura refrescante de repolho, cenoura e molho cremoso|15.90|side"
    ["Mandioca Frita"]="Palitos de mandioca fritos, com toque crocante e interior macio|22.90|side"
    ["Salada Verde"]="Mix de folhas, tomate cereja, pepino e molho de mostarda e mel|13.50|side"
    ["Batata Rústica"]="Batatas assadas com ervas finas e toque de alho|23.90|side"
    ["Molho Especial"]="Combinação de maionese, ketchup, mostarda e especiarias, ideal para mergulhar|4.90|side"
    ["Chips de Batata Doce"]="Fatias finas de batata doce fritas ou assadas, levemente salgadas|16.80|side"
    ["Queijo Coalho Grelhado"]="Espetinhos de queijo coalho grelhado, com toque de orégano|23.90|side"
    ["Palitinhos de Cenoura e Pepino"]="Alternativa saudável, acompanhada de molho de iogurte ou homus|20.90|side"

    ["Refrigerante em Lata"]="Clássicos como Coca-Cola, Guaraná, Sprite e Fanta|5.00|drink"
    ["Suco Natural"]="Opções como laranja, limão, maracujá e acerola, feitos na hora|8.00|drink"
    ["Água com Gás"]="Refrescante e ideal para quem prefere uma opção sem açúcar|4.00|drink"
    ["Água Mineral"]="Pura e sem gás, para hidratar|3.00|drink"
    ["Chá Gelado"]="Sabores como limão, pêssego e mate, servidos bem gelados|6.00|drink"
    ["Água Tônica"]="Ideal para quem gosta de um sabor amargo e sofisticado|5.50|drink"
    ["Energético"]="Para quem precisa de um impulso extra|12.00|drink"
    ["Limonada Suíça"]="Saborosa e levemente azeda, com toque de leite condensado|10.00|drink"
    ["Milkshake"]="Opções clássicas como chocolate, morango e baunilha|15.00|drink"
    ["Refrigerante Zero"]="Alternativa sem açúcar para os clássicos|5.00|drink"

    ["Brigadeiro Gourmet"]="Tradicional doce brasileiro com sabor intenso de chocolate|4.00|dessert"
    ["Brownie de Chocolate"]="Massinha densa e saborosa, com ou sem nozes|8.00|dessert"
    ["Cheesecake de Frutas Vermelhas"]="Sobremesa cremosa com cobertura de frutas|12.00|dessert"
    ["Pudim de Leite Condensado"]="Clássico doce brasileiro, cremoso e caramelizado|10.00|dessert"
    ["Mousse de Maracujá"]="Cremosa e com aquele toque azedinho característico|7.00|dessert"
    ["Petit Gâteau"]="Bolinho quente de chocolate com recheio cremoso, servido com sorvete|20.00|dessert"
    ["Cookies Caseiros"]="Crocanes por fora, macios por dentro, com gotas de chocolate|6.00|dessert"
    ["Sorvete Artesanal"]="Diversos sabores como chocolate belga, pistache e frutas|10.00|dessert"
    ["Torta de Limão"]="Massa crocante com recheio azedinho e cobertura de merengue|9.00|dessert"
    ["Banoffee Pie"]="Camadas de banana, doce de leite, chantilly e biscoito crocante|12.00|dessert"

)

for NOME in "${!ITENS[@]}"; do
    IFS='|' read -r -a array <<< "${ITENS[$NOME]}"

    DESCRICAO=${array[0]}
    PRECO=${array[1]}
    CATEGORIA=${array[2]}
    
    curl -X POST "$API_URL" -H "$HEADER" \
    -d "{
        \"name\": \"$NOME\",
        \"description\": \"$DESCRICAO\",
        \"price\": \"$PRECO\",
        \"category\": \"$CATEGORIA\"
    }"
    
    echo -e "\nEnviado: $NOME - $DESCRICAO - $PRECO - $CATEGORIA" 
done
