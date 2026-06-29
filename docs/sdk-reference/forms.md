# `forms` Golden Namespace

Go client access: `client.Forms`


These methods are Golden because they come from bundled Stoplight controller contracts.

## Methods

### `submit_form`

- Go wrapper: `client.Forms.SubmitForm(ctx, opts, out)`
- Dynamic helper: `client.RequestGolden(ctx, "forms", "submit_form", opts, out)`
- HTTP route: `POST /forms/submit`
- Source controller: `Forms`

Submit new record to a form

#### Supports form submissions to Incident IQ using pre-configured forms that trigger actions upon submission
#### Sample request:
```
POST /api/v1.0/forms/submit
{
  "FormKey": "test-form-1",
  "SiteId": "ac6cece8-e4f4-e511-a789-005056bb000e",
  "Data": [ {
    "CustomFieldTypeId": "ac6cece8-e4f4-e511-a789-005056bb000e",
    "Value": "Test form submission",
  } ]
}
```

#### Parameters

| RequestOptions Field | API Name | In | Required | Type | Schema / Model | Description |
| --- | --- | --- | --- | --- | --- | --- |
| `JSON` | `Model` | `body` | `yes` | `FormSubmission` | `FormSubmission` | Form submission infomation / parameters |

#### Returns

- Go wrapper return: `error`; decoded `ActionResponse` responses are written into `out`.
- Decoded response: caller-provided `out` receives `map[string]any | []any | nil` when the route returns JSON.
- Response model: `ActionResponse`
- Pagination helper: No paging query parameters detected; the generated wrapper returns one decoded response through `out`.

---
