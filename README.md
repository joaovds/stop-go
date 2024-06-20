## Entidades

- Player
- Room
- ...

## Requisitos Funcionais

### 1. Manipulação de Salas (rooms)

#### 1.1 Criar Sala
- **Descrição:** Um jogador poderá criar uma nova sala de jogo
- **Detalhes:**
    - O criador da sala terá automaticamente o cargo de anfitrião
    - O criador da sala poderá adicionar um nome para a sala
    - Caso o nome não seja dado, um nome aleatório deve ser gerado
    - O criador da sala poderá configurar a quantidade máxima e mínima de jogadores
    - A quantidade mínima de jogadores deve ser no mínimo 2
    - O criador da sala poderá restringir o acesso a sala com alguma senha
    - O criador da sala poderá permitir ou não o ingresso de novas pessoas caso o jogo já tenha iniciado
    - A sala poderá ter n partidas, configuradas para ter um vencedor por quantidade de pontos, jogadas ou tempo
    - Deve ser gerado um código de 6 caracteres para a identificação e acesso da sala

#### 1.2 Entrar em Sala
- **Descrição:** Jogadores poderão entrar em uma sala existente
- **Detalhes:**
    - Deve ser informado o código da sala e verificar se o mesmo é válido
    - Caso seja uma sala restrita, o jogador deverá informar a senha junto ao código da sala
    - A senha deverá ser validada
    - Deve-se verificar se a sala já não atingiu a quantidade máxima de jogadores
    - O jogador deve ser adicionado na sala do código informado, e ser notificado do ingresso
    - Em qualquer tipo de falha, deve-se informar o usuário

#### 1.3 Listar salas
- **Descrição:** Jogadores poderão ver uma lista de salas disponíveis no momento
- **Detalhes:**
    - A lista deve mostrar o nome da sala, a quantidade atual de jogadores e o máximo permitido (ex: 2/6)
    - A listagem não deve mostrar a identificação dos jogadores da sala para jogares de fora dela

#### 1.4 Atualizar configurações da sala
- **Descrição:** A sala poderá ter suas configurações alteradas
- **Detalhes:**
    - O anfitrião ou jogadores com cargo de adm poderão alterar as configurações da sala
    - As configurações só poderão ser alteradas quando não houver nenhuma partida em andamento
    - O anfitrião poderá alterar os cargos dos jogadores da sala

#### 1.5 Sair de sala
- **Descrição:** Um jogador poderá sair da sala
- **Detalhes:**
    - Caso o ***anfitrião*** da sala saia, o cargo será passado para o user mais antigo que entrou na sala
    - Quando o último jogador sair, a sala deve ser excluída

### 2. Jogadores

#### 2.1 Criar jogador (sessão)
- **Descrição:**  Uma sessão poderá ser iniciada para jogar partidas
- **Detalhes:**
    - Qualquer pessoa poderá criar uma nova sessão para jogar
    - Não é necessário login
    - Deve ser informado um ***nickname*** para identificação do jogador
