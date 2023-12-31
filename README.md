# Gurl
A cURL imitation built from scratch in go

### Four method requests are currently implemented:
 | Method | Command |
| ----------- | ----------- |
| GET | gurl *exampleURL* |
| POST | gurl **-post** *exampleURL* |
| DELETE | gurl **-delete** *exampleURL* |
| PUT | gurl **-put** *exampleURL* |
---
## Important note!
- Always prepend your url with `http://` or `https://`, the tool can't find your url otherwise
