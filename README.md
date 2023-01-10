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

```bash
short-circuit
```

## Base URL

* /api/files

## Endpoints

* /api/files
* /api/files/{fileId}
* /api/files/{fileId}/size
* /api/files/{fileId}/bars
* /api/files/{fileId}/type
* /api/files/{fileId}/type/{typeId}
* /api/files/{fileId}/type/{typeId}/elements
* /api/files/{fileId}/type/{typeId}/elements/{elementId}

## Historic

|Version|Description|Requester|Date|
|:------|:----------|:---|:---|
|2.0.0|Alterado framework de API REST e alterado os endpoints|Luis Avelino|16-12-2022|
|1.0.0|Desenvolver API dos elementos do sistema|Luis Avelino|16-12-2022|
