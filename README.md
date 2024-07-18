# Projeto de Rate Limit

Este projeto implementa um sistema de limitação de taxa de requisições (rate limit) configurável, projetado para prevenir o uso excessivo de APIs ou endpoints. A configuração do sistema é baseada em variáveis de ambiente, permitindo uma fácil adaptação a diferentes cenários de uso.

## Características

- **Configuração via `.env`**: Todas as configurações principais, incluindo limites de taxa e conexões ao banco de dados, são gerenciadas através de um arquivo `.env`. Isso facilita a personalização e a configuração em diferentes ambientes sem a necessidade de alterar o código.
- **Execução via Docker Compose**: Para simplificar a implantação e garantir a consistência entre os ambientes de desenvolvimento e produção, o projeto é empacotado e executado utilizando o Docker Compose. Isso encapsula o serviço da aplicação e suas dependências (como o Redis) em contêineres.
- **Uso do Redis**: O controle de rate limit é realizado com o auxílio do Redis, um armazenamento de estrutura de dados em memória, que proporciona alta performance e latência baixa para operações de leitura e escrita necessárias para o controle de taxa.

## Personalização do Banco de Dados

Para trocar o tipo de banco de dados utilizado pelo sistema, é necessário implementar um novo `DBClient` que siga a interface determinada pelo projeto. Isso permite a integração com diferentes sistemas de gerenciamento de banco de dados (SGBDs) de acordo com as necessidades específicas de cada cenário de uso.

## Pré-requisitos

Antes de iniciar, certifique-se de que você tem o Docker e o Docker Compose instalados em sua máquina. Essas ferramentas são necessárias para construir e executar o ambiente de contêineres.

## Configuração

1. Clone o repositório para sua máquina local.
2. Crie um arquivo `.env` na pasta [`cmd/server`] do projeto, seguindo o exemplo fornecido em `.env.example`. Adapte as configurações de acordo com suas necessidades.
3. Execute o comando abaixo para iniciar os serviços via Docker Compose:

```bash
docker-compose -f docker-compose.dev.yml up