package dataverse_config

import (
	"os"
	"time"
	
	"github.com/spf13/viper"
)

const (
	AuthEndpoint      = "dataverse-auth-endpoint"
	AuthScope         = "dataverse-auth-scope"
	AuthRenewInterval = "dataverse-auth-renew-interval"
	Endpoint          = "dataverse-endpoint"
	ClientId          = "dataverse-client-id"
	ClientSecret      = "dataverse-client-secret"
	TenantId          = "dataverse-tenant-id"
)

func Read(cfg *viper.Viper) {
	cfg.Set(AuthEndpoint, "https://login.microsoftonline.com/c139a2ef-3d33-4c87-bfb7-53fda8b7749c/oauth2/v2.0/token")
	cfg.Set(AuthScope, "https://catch-demo-xy.api.crm4.dynamics.com/.default")
	cfg.Set(AuthRenewInterval, 15*time.Minute)
	cfg.Set(Endpoint, "https://catch-demo-xy.api.crm4.dynamics.com/api/data/v9.2")
	cfg.Set(ClientId, os.Getenv("DATAVERSE_CLIENT_ID"))
	cfg.Set(ClientSecret, os.Getenv("DATAVERSE_CLIENT_SECRET"))
	cfg.Set(TenantId, os.Getenv("DATAVERSE_TENANT_ID"))
}
