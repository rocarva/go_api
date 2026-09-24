# Frontend — Personalidades

Frontend em React (Create React App) que lista as personalidades consumindo a
API Go.

## Como rodar

```bash
npm install
npm start
```

Abra `http://localhost:3000`.

## Configuração

A URL da API é definida pela variável `REACT_APP_API_URL` (padrão
`http://localhost:8000`):

```bash
REACT_APP_API_URL=http://localhost:8000 npm start
```

## Scripts disponíveis

| Script           | Descrição                                        |
| ---------------- | ------------------------------------------------ |
| `npm start`      | Sobe o app em modo desenvolvimento                |
| `npm run build`  | Gera o build de produção na pasta `build/`        |
| `npm test`       | Roda os testes em modo watch                      |
| `npm run eject`  | Remove a abstração do CRA (operação irreversível) |

## Estrutura

```
src/
├── components/
│   ├── Personalidades.js   # Busca e lista as personalidades
│   └── Personalidades.css
├── App.js                  # Cabeçalho
└── index.js                # Ponto de entrada
```
