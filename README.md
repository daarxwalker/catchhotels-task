# CatchHotels Task

## Dev start

```bash
docker compose build --build-arg DATAVERSE_CLIENT_ID=<client-id> --build-arg DATAVERSE_CLIENT_SECRET=<client-secret>  --build-arg DATAVERSE_TENANT_ID=<tenant_id>
docker compose up
```

App is running on port `:8000`, but uses Caddy to proxy to `:80`, so you can access it on `http://localhost`

## Docs path
``http://localhost/docs``
