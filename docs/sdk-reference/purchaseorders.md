# `purchaseorders` Golden Namespace

Go client access: `client.Purchaseorders`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Aliases

| Alias | Canonical Method | Route |
| --- | --- | --- |
| `delete` | `delete_purchase_order` | `DELETE /purchaseorders/{PurchaseOrderId}` |
| `get` | `get_purchase_order` | `GET /purchaseorders/{PurchaseOrderId}` |
| `list` | `get_purchase_orders` | `GET /purchaseorders` |
| `update` | `update_purchase_order` | `POST /purchaseorders/{PurchaseOrderId}` |

## Methods

### `delete_purchase_order`

- Go wrapper: `client.Purchaseorders.DeletePurchaseOrder(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "purchaseorders", "delete_purchase_order", opts, out)`
- HTTP route: `DELETE /purchaseorders/{PurchaseOrderId}`
- Source controller: `Parts`
- Aliases: `delete`

Delete a Purchase Order

#### Delete a specific purchase order.
#### Sample request:
```
DELETE /api/v1.0/purchaseorders/b199f092-e1a9-418b-8eef-f6e44e273539
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PurchaseOrderId"]` | `PurchaseOrderId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemDeleteResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemDeleteResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_purchase_order`

- Go wrapper: `client.Purchaseorders.GetPurchaseOrder(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "purchaseorders", "get_purchase_order", opts, out)`
- HTTP route: `GET /purchaseorders/{PurchaseOrderId}`
- Source controller: `Parts`
- Aliases: `get`

Get Purchase Order

#### Retrieve a specific purchase order with a specific PurchaseOrderId
#### Sample request:
```
GET /api/v1.0/purchaseorders/b199f092-e1a9-418b-8eef-f6e44e273539
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PurchaseOrderId"]` | `PurchaseOrderId` | `path` | `yes` | `string` | `-` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemGetResponseOfPurchaseOrder` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemGetResponseOfPurchaseOrder`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `get_purchase_orders`

- Go wrapper: `client.Purchaseorders.GetPurchaseOrders(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "purchaseorders", "get_purchase_orders", opts, out)`
- HTTP route: `GET /purchaseorders`
- Source controller: `Parts`
- Aliases: `list`

Get Purchase Orders

#### Retrieves a list of purchase orders. A specific purchase order can be retrieved via GET `api/v1.0/purchaseorders/{PurchaseOrderId:guid}`.
#### Sample request:
```
GET /api/v1.0/purchaseorders
```

#### Parameters

This operation does not define request parameters.

#### Returns

- Go wrapper return: `error`; decoded `ListGetResponseOfPurchaseOrder` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ListGetResponseOfPurchaseOrder`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---

### `update_purchase_order`

- Go wrapper: `client.Purchaseorders.UpdatePurchaseOrder(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "purchaseorders", "update_purchase_order", opts, out)`
- HTTP route: `POST /purchaseorders/{PurchaseOrderId}`
- Source controller: `Parts`
- Aliases: `update`

No contract summary provided.

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `PathParams["PurchaseOrderId"]` | `PurchaseOrderId` | `path` | `yes` | `string` | `-` | - |
| `JSON` | `PurchaseOrder` | `body` | `yes` | `PurchaseOrder` | `PurchaseOrder` | - |

#### Returns

- Go wrapper return: `error`; decoded `ItemUpdateResponseOfPurchaseOrder` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ItemUpdateResponseOfPurchaseOrder`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
