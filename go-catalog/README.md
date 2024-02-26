## Running catalog service

Enter in catalog service

```bash
  cd go-catalog
```

Start mysql database

```bash
   docker-compose up -d
```

Install dependencies

```bash
  go mod download
```

Init server

```bash
  go run cmd/catalog/main.go
```

#### Get all categories

```http
  http://localhost:8080
```

```http
  GET /category
```

#### Get one category

```http
  GET /category/${id}
```

| Param | Type     | Description                            |
| :---- | :------- | :------------------------------------- |
| `id`  | `string` | **Required**. The category id you want |

#### Add category

```http
  POST /category
```

| Body                      | Type   | Description                                                         |
| :------------------------ | :----- | :------------------------------------------------------------------ |
| `{"name": "Category 01"}` | `json` | **Required**. A JSON with the name of the category to be registered |
