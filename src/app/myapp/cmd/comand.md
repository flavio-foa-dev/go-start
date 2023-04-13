- deletar os arquivos go.mod e go.sum,
-  gerar novamente com o comando "go mod init"
-  e depois executei o comando "go mod tidy".
-  go get

$ cd $GOPATH
bin pkg src

- bin == arquivos binarios os executaveis
- pkg == referente a todos os pacotes da aplicacao
- src == referente a todos os armazenamente de codigo fonte de mim e de terceiros

Modularizamos o código para tornar a manutenção e/ou atualização mais clara, criando as pastas controllers, db, models, routes e templates;

Criamos uma página para criar novos produtos e uma rota capaz de atender essa requisição http.HandleFunc("/new", controllers.New);

Buscamos os dados da página new com o código r.FormValue() para cada input (nome, descrição, preço e quantidade) no controller de produtos;

Salvamos o produto através do modelo de produto criando a função CriaNovoProduto().

