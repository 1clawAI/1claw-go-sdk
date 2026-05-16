package oneclaw

// Platform API operations are available via the PlatformService.
// Regenerate the Go OpenAPI client (make generate) after updating
// packages/openapi-spec/openapi.yaml to get typed request/response structs.
//
// Endpoints:
//   POST   /v1/platform/apps                          — CreateApp
//   GET    /v1/platform/apps                          — ListApps
//   GET    /v1/platform/apps/{appId}                  — GetApp
//   PATCH  /v1/platform/apps/{appId}                  — UpdateApp
//   DELETE /v1/platform/apps/{appId}                  — DeleteApp
//   POST   /v1/platform/apps/{appId}/templates        — CreateTemplate
//   GET    /v1/platform/apps/{appId}/templates        — ListTemplates
//   POST   /v1/platform/users/upsert                  — UpsertUser
//   GET    /v1/platform/apps/{appId}/users            — ListUsers
//   POST   /v1/platform/connections/{connId}/bootstrap — BootstrapUser
//   GET    /v1/platform/claim/{token}                 — ClaimPreview (public, no auth)
//   POST   /v1/platform/claim/{token}                 — ClaimRedeem (public, no auth)
//   GET    /v1/platform/connected-apps                — ListConnectedApps (user-side)
//   DELETE /v1/platform/connected-apps/{connId}       — DisconnectApp (user-side)
