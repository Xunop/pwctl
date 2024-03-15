---
title: pwatch3
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.22"
---

# pwatch3

Base URLs:

- <a href="http://localhost:8080">test env: http://localhost:8080</a>

# Authentication

# Auth

## POST Login

POST /login

> Body Parameters

```json
{
  "user": "admin",
  "password": "admin"
}
```

### Params

| Name       | Location | Type   | Required | Description |
| ---------- | -------- | ------ | -------- | ----------- |
| body       | body     | object | no       | none        |
| » user     | body     | string | yes      | none        |
| » password | body     | string | yes      | none        |

> Response Examples

> Success

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Db

## GET Get Monitore Db

GET /db

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |

> Response Examples

> Success

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## DELETE Delete Monitore Db

DELETE /db

### Params

| Name  | Location | Type   | Required | Description  |
| ----- | -------- | ------ | -------- | ------------ |
| id    | query    | string | no       | Db Unique ID |
| TOKEN | header   | string | no       | none         |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## POST Add Monitored DB

POST /db

> Body Parameters

```json
{
  "md_name": "string",
  "md_connstr": "string",
  "md_preset_config_name": "minimal",
  "md_config": null,
  "md_is_superuser": true,
  "md_is_enabled": true
}
```

### Params

| Name                    | Location | Type    | Required | Description                    |
| ----------------------- | -------- | ------- | -------- | ------------------------------ |
| TOKEN                   | header   | string  | no       | none                           |
| body                    | body     | object  | no       | none                           |
| » md_name               | body     | string  | yes      | db name                        |
| » md_connstr            | body     | string  | yes      | connection string              |
| » md_preset_config_name | body     | string  | yes      | default is 'basic' in database |
| » md_config             | body     | null    | yes      | none                           |
| » md_is_superuser       | body     | boolean | no       | default is false in database   |
| » md_is_enabled         | body     | boolean | no       | default is true in database    |

#### Description

**» md_preset_config_name**: default is 'basic' in database

this value must in `preset_cofig` table

#### Enum

| Name                    | Value               |
| ----------------------- | ------------------- |
| » md_preset_config_name | minimal             |
| » md_preset_config_name | basic               |
| » md_preset_config_name | standard            |
| » md_preset_config_name | pgbouncer           |
| » md_preset_config_name | pgpool              |
| » md_preset_config_name | exhaustive          |
| » md_preset_config_name | full                |
| » md_preset_config_name | full_influx         |
| » md_preset_config_name | unprivileged        |
| » md_preset_config_name | aurora              |
| » md_preset_config_name | azure               |
| » md_preset_config_name | rds                 |
| » md_preset_config_name | gce                 |
| » md_preset_config_name | prometheus          |
| » md_preset_config_name | prometheus-async    |
| » md_preset_config_name | superuser_no_python |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## PATCH Edit Enable Db

PATCH /db

> Body Parameters

```json
{
  "md_is_enabled": true
}
```

### Params

| Name            | Location | Type    | Required | Description  |
| --------------- | -------- | ------- | -------- | ------------ |
| id              | query    | string  | no       | Db Unique ID |
| TOKEN           | header   | string  | no       | none         |
| body            | body     | object  | no       | none         |
| » md_is_enabled | body     | boolean | yes      | none         |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## POST Test Connection

POST /test-connect

> Body Parameters

```
postgresql://pgwatch3:pgwatch3admin@postgres/pgwatch3

```

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |
| body  | body     | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Stats

## GET Stats Summary

GET /stats

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Metrics

## GET Get stored metric

GET /metric

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## DELETE Delete metric

DELETE /metric

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| id    | query    | string | no       | none        |
| TOKEN | header   | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## POST Add metric

POST /metric

> Body Parameters

```json
{}
```

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |
| body  | body     | object | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Preset

## GET Get presets

GET /preset

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| TOKEN | header   | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## DELETE Delete presets

DELETE /preset

### Params

| Name  | Location | Type   | Required | Description |
| ----- | -------- | ------ | -------- | ----------- |
| id    | query    | string | no       | pc_name     |
| TOKEN | header   | string | no       | none        |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## POST Add presets

POST /preset

> Body Parameters

```json
{
  "pc_config": {
    "backends": 10,
    "cpu_load": 10
  },
  "pc_name": "test1",
  "pc_description": "test1"
}
```

### Params

| Name             | Location | Type   | Required | Description                     |
| ---------------- | -------- | ------ | -------- | ------------------------------- |
| TOKEN            | header   | string | no       | none                            |
| body             | body     | object | no       | none                            |
| » pc_config      | body     | object | yes      | metric_name and update interval |
| » pc_name        | body     | string | yes      | none                            |
| » pc_description | body     | string | yes      | none                            |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

## PATCH Edit presets

PATCH /preset

> Body Parameters

```json
{
  "pc_config": {
    "backends": 10,
    "cpu_load": 10
  },
  "pc_name": "test1",
  "pc_description": "test1"
}
```

### Params

| Name             | Location | Type   | Required | Description                     |
| ---------------- | -------- | ------ | -------- | ------------------------------- |
| id               | query    | string | no       | pc_name                         |
| TOKEN            | header   | string | no       | none                            |
| body             | body     | object | no       | none                            |
| » pc_config      | body     | object | yes      | metric_name and update interval |
| » pc_name        | body     | string | yes      | none                            |
| » pc_description | body     | string | yes      | none                            |

> Response Examples

> 200 Response

```json
{}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

# Data Schema
