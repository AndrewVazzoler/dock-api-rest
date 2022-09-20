# Cenário

A Dock está crescendo e expandindo seus negócios, gerando novas oportunidades de revolucionar o mercado financeiro e criar produtos diferenciados.
Nossa próxima missão é construir uma nova conta digital Dock para nossos clientes utilizarem através de endpoints, onde receberemos requisições em um novo backend que deverá gerenciar as contas e seus portadores (os donos das contas digitais).

# Requisitos

- Deve ser possível criar e remover **portadores**
  - Um **portador** deve conter apenas seu _nome completo_ e _CPF_
  - O _CPF_ deve ser válido e único no cadastro de **portadores**
- As **contas digital Dock** devem conter as seguintes funcionalidades:
  - A conta deve ser criada utilizando o _CPF_ do **portador**
  - Uma conta deve ter seu _saldo_, _número_ e _agência_ disponíveis para consulta
  - Necessário ter funcionalidade para fazer a _consulta de extrato_ da conta _por período_
  - Um **portador** pode fechar a **conta digital Dock** a qualquer instante
  - Executar as operações de _saque_ e _depósito_
    - _Depósito_ é liberado para todas as _contas ativas_ e _desbloqueadas_
    - _Saque_ é permitido para todas as _contas ativas_ e _desbloqueadas_ desde que haja _saldo disponível_ e não ultrapasse o limite diário de _2 mil reais_

## Regulação obrigatória

- Precisamos _bloquear_ e _desbloquear_ a **conta digital Dock** a qualquer momento
- A **conta digital Dock** nunca poderá ter o _saldo negativo_

# Orientações

Utilize qualquer uma das linguagens de programação:

- Java
- Javascript
- Typescript
- Python
- Kotlin
- Golang

Desenvolva o case seguindo as melhores práticas que julgar necessário, aplique todos os conceitos, se atente a qualidade, utilize toda e qualquer forma de governança de código válido. Vamos considerar toda e qualquer implementação, trecho de código, documentação e/ou intenção compartilhada conosco. Esperamos também que o desafio seja feito dentro do tempo disponibilizado e que esteja condizente com a posição pretendida.

É necessário ter o desafio 100% funcional contendo informações e detalhes sobre: como iniciar a aplicação, interagir com as funcionalidades disponíveis e qualquer outro ponto adicional.

## Diferenciais

- Práticas, padrões e conceitos de microservices será considerado um diferencial para nós por existir uma variedade de produtos e serviços dentro da Dock.
- Temos 100% das nossas aplicações e infraestrutura na nuvem, consideramos um diferencial, caso o desafio seja projeto para ser executado na nuvem.
- Nossos times são autônomos e têm liberdade para definir arquiteturas e soluções. Por este motivo será considerado diferencial toda: arquitetura, design, paradigma e documentação detalhando a sua abordagem.

### Instruções

      1. Faça o fork do desafio;
      2. Crie um repositório privado no seu github para o projeto e adicione como colaborador, os usuários informados no email pelo time de recrutameto ;
      3. Após concluir seu trabalho faça um push;
      4. Envie um e-mail à pessoa que está mantendo o contato com você durante o processo notificando a finalização do desafio para validação.
