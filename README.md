# short-circuit-analysis-elements

Api do projeto de análise de curto-circuito

## Objectives

Projeto responsável por retornar todos os elementos do tipo 1, 2 e 3, do sistema de potência.

## Functionalities

* Todos os arquivos disponíveis a serem analisados
* Todos os elementos do sistema de potência
* Todos os elementos de determinado tipo
* Elemento específico selecionado de determinado tipo

## Environment

```
short-circuit
```

## Base URL

* /api/files

## Endpoints

* /api/files
* /api/files/{fileId}
* /api/files/{fileId}/elements
* /api/files/{fileId}/elements/type/{typeId}
* /api/files/{fileId}/elements/type/{typeId}/element/{elementId}

## Historic

|Version|Description|Requester|Date|
|:------|:----------|:---|:---|
|1.0.0|Desenvolver API dos elementos do sistema|Luis Avelino|16-12-2022|
