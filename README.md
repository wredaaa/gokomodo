
## API Reference Buyer

#### Login

```http
  POST /api/auth/buyer
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### View List of Products

```http
  GET /api/products
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Token`      | `string` | **Required**. add Berear Token |


#### Order product

```http
  POST /api/orders
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `buyer_id` | `int` | **Required**. |
| `seller_id` | `int` | **Required**. |
| `delivery_source` | `string` | **Required**. |
| `delivery_destination` | `string` | **Required**. |
| `product_id` | `int` | **Required**. |
| `items` | `string` | **Required**. |
| `price` | `decimal` | **Required**. |
| `quantity` | `int` | **Required**. |
| `total_price` | `decimal` | **Required**. |
| `status` | `string` | **Required**. |
| `Token`      | `string` | **Required**. add Token in Headers |


## API Reference Seller

#### Login

```http
  POST /api/auth/seller
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### View product list

```http
  GET /api/products
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Token`      | `string` | **Required**. add Berear Token |


#### Add new product

```http
  POST /api/products
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. |
| `description` | `string` | **Required**. |
| `price` | `decimal` | **Required**. |
| `seller_id` | `int` | **Required**. |
| `Token`      | `string` | **Required**. add Token in Headers |


#### View order list

```http
  GET /api/orders
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Token`      | `string` | **Required**. add Berear Token |

#### Accept Order

```http
  PUT /api/orders/update/status/{order_id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `status`     | `string` | **Required**. |
| `Token`      | `string` | **Required**. add Berear Token |
