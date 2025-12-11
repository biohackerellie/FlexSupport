package middleware

import (
	"context"
	"net/http"
)

type ctxKey string

const (
	ctxTenantID   ctxKey = "tenant_id"
	ctxTenantSlug ctxKey = "tenant_slug"
)

type Tenant struct {
	ID   string
	Slug string
	Name string
}

type TenantResolver interface {
	ResolveByHost(ctx context.Context, host string) (*Tenant, bool, error)
	ResolveBySlug(ctx context.Context, slug string) (*Tenant, bool, error)
}

func TenantMiddleware(resolver TenantResolver, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		tenant, ok, err := resolver.ResolveByHost(r.Context(), host)
		if err != nil {
			http.Error(w, "tenant lookup failed", http.StatusInternalServerError)
			return
		}

		// // Optional: fallback to /t/{slug}/... (dev mode)
		// if !ok {
		// 	slug, ok2 := tenantSlugFromPath(r.URL.Path) // you implement
		// 	if ok2 {
		// 		tenant, ok, err = resolver.ResolveBySlug(r.Context(), slug)
		// 		if err != nil {
		// 			http.Error(w, "tenant lookup failed", http.StatusInternalServerError)
		// 			return
		// 		}
		// 	}
		// }

		if !ok {
			http.NotFound(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), ctxTenantID, tenant.ID)
		ctx = context.WithValue(ctx, ctxTenantSlug, tenant.Slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
