// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

// tmplCtx template context for generate project
type tmplCtx struct {
	PkgName string
}

// tmplInfo template data info
type tmplInfo struct {
	name   string
	isTmpl bool
	isExec bool
}

// tmplFiles map of generated file's assets info
var tmplFiles = map[string]map[string]tmplInfo{
	"grpc": {
		"cmd/zimctl/.gitignore":                                          {"grpc/cmd/zimctl/.gitignore.tmpl", false, false},
		"cmd/zimctl/root.go":                                             {"grpc/cmd/zimctl/root.tmpl", true, false},
		"cmd/zimctl/grpc.go":                                             {"grpc/cmd/zimctl/grpc.tmpl", true, false},
		"cmd/zimctl/admin.go":                                            {"grpc/cmd/zimctl/admin.tmpl", true, false},
		"cmd/zimctl/version.go":                                          {"grpc/cmd/zimctl/version.tmpl", true, false},
		"cmd/zimlet.go":                                                  {"grpc/cmd/zimlet.tmpl", true, false},
		"cmd/core/core.go":                                               {"grpc/cmd/core/core.tmpl", true, false},
		"cmd/zimlet/version.go":                                          {"grpc/cmd/zimlet/version.tmpl", true, false},
		"cmd/zimlet/agent.go":                                            {"grpc/cmd/zimlet/agent.tmpl", true, false},
		"cmd/zimlet/.gitignore":                                          {"grpc/cmd/zimlet/.gitignore.tmpl", true, false},
		"cmd/zimlet/admin.go":                                            {"grpc/cmd/zimlet/admin.tmpl", true, false},
		"cmd/zimlet/affairs.go":                                          {"grpc/cmd/zimlet/affairs.tmpl", true, false},
		"cmd/zimlet/root.go":                                             {"grpc/cmd/zimlet/root.tmpl", true, false},
		"cmd/README.md":                                                  {"grpc/cmd/readme_md.tmpl", false, false},
		"cmd/zimctl.go":                                                  {"grpc/cmd/zimctl.tmpl", true, false},
		"servants/dataware/dw_metrics.go":                                {"grpc/servants/dataware/dw_metrics.tmpl", true, false},
		"servants/dataware/dataware.go":                                  {"grpc/servants/dataware/dataware.tmpl", true, false},
		"servants/dataware/dw_metrics_fake.go":                           {"grpc/servants/dataware/dw_metrics_fake.tmpl", true, false},
		"servants/dataware/dw_push_msg.go":                               {"grpc/servants/dataware/dw_push_msg.tmpl", true, false},
		"servants/business/affairs_handler.go":                           {"grpc/servants/business/affairs_handler.tmpl", true, false},
		"servants/business/internal_admin.go":                            {"grpc/servants/business/internal_admin.tmpl", true, false},
		"servants/business/admin_handler.go":                             {"grpc/servants/business/admin_handler.tmpl", true, false},
		"servants/business/business.go":                                  {"grpc/servants/business/business.tmpl", true, false},
		"servants/business/external_affairs.go":                          {"grpc/servants/business/external_affairs.tmpl", true, false},
		"servants/agent/agent_ws_svc.go":                                 {"grpc/servants/agent/agent_ws_svc.tmpl", true, false},
		"servants/agent/agent_wards_svc.go":                              {"grpc/servants/agent/agent_wards_svc.tmpl", true, false},
		"servants/agent/agent.go":                                        {"grpc/servants/agent/agent.tmpl", true, false},
		"servants/agent/agent_svc.go":                                    {"grpc/servants/agent/agent_svc.tmpl", true, false},
		"servants/agent/agent_wards.go":                                  {"grpc/servants/agent/agent_wards.tmpl", true, false},
		"servants/agent/agent_ws.go":                                     {"grpc/servants/agent/agent_ws.tmpl", true, false},
		"servants/agent/agent_wards_handler.go":                          {"grpc/servants/agent/agent_wards_handler.tmpl", true, false},
		"servants/servants.go":                                           {"grpc/servants/servants.tmpl", true, false},
		"servants/share/share.go":                                        {"grpc/servants/share/share.tmpl", true, false},
		"go.mod":                                                         {"grpc/go_mod.tmpl", true, false},
		"core/proto/admin/v1/internal_admin.proto":                       {"grpc/core/proto/admin/v1/internal_admin_proto.tmpl", true, false},
		"core/proto/admin/v1/internal_admin_api.proto":                   {"grpc/core/proto/admin/v1/internal_admin_api_proto.tmpl", true, false},
		"core/proto/affairs/v1/external_affairs_api.proto":               {"grpc/core/proto/affairs/v1/external_affairs_api_proto.tmpl", true, false},
		"core/proto/affairs/v1/external_affairs.proto":                   {"grpc/core/proto/affairs/v1/external_affairs_proto.tmpl", true, false},
		"core/proto/agent/v1/agent_wards.proto":                          {"grpc/core/proto/agent/v1/agent_wards_proto.tmpl", true, false},
		"core/proto/agent/v1/agent_wards_api.proto":                      {"grpc/core/proto/agent/v1/agent_wards_api_proto.tmpl", true, false},
		"core/proto/gen/admin/v1/internal_admin.validator.pb.go":         {"grpc/core/proto/gen/admin/v1/internal_admin_validator_pb.tmpl", true, false},
		"core/proto/gen/admin/v1/internal_admin_api.validator.pb.go":     {"grpc/core/proto/gen/admin/v1/internal_admin_api_validator_pb.tmpl", true, false},
		"core/proto/gen/admin/v1/internal_admin.pb.go":                   {"grpc/core/proto/gen/admin/v1/internal_admin_pb.tmpl", true, false},
		"core/proto/gen/admin/v1/internal_admin_api.pb.go":               {"grpc/core/proto/gen/admin/v1/internal_admin_api_pb.tmpl", true, false},
		"core/proto/gen/affairs/v1/external_affairs_api.pb.go":           {"grpc/core/proto/gen/affairs/v1/external_affairs_api_pb.tmpl", true, false},
		"core/proto/gen/affairs/v1/external_affairs.validator.pb.go":     {"grpc/core/proto/gen/affairs/v1/external_affairs_validator_pb.tmpl", true, false},
		"core/proto/gen/affairs/v1/external_affairs.pb.go":               {"grpc/core/proto/gen/affairs/v1/external_affairs_pb.tmpl", true, false},
		"core/proto/gen/affairs/v1/external_affairs_api.validator.pb.go": {"grpc/core/proto/gen/affairs/v1/external_affairs_api_validator_pb.tmpl", true, false},
		"core/proto/gen/agent/v1/agent_wards_api.validator.pb.go":        {"grpc/core/proto/gen/agent/v1/agent_wards_api_validator_pb.tmpl", true, false},
		"core/proto/gen/agent/v1/agent_wards.validator.pb.go":            {"grpc/core/proto/gen/agent/v1/agent_wards_validator_pb.tmpl", true, false},
		"core/proto/gen/agent/v1/agent_wards.pb.go":                      {"grpc/core/proto/gen/agent/v1/agent_wards_pb.tmpl", true, false},
		"core/proto/gen/agent/v1/agent_wards_api.pb.go":                  {"grpc/core/proto/gen/agent/v1/agent_wards_api_pb.tmpl", true, false},
		"core/proto/gen/common/v1/metrics.pb.go":                         {"grpc/core/proto/gen/common/v1/metrics_pb.tmpl", true, false},
		"core/proto/gen/common/v1/metrics.validator.pb.go":               {"grpc/core/proto/gen/common/v1/metrics_validator_pb.tmpl", true, false},
		"core/proto/common/v1/metrics.proto":                             {"grpc/core/proto/common/v1/metrics_proto.tmpl", true, false},
		"core/proto/proto.go":                                            {"grpc/core/proto/proto.tmpl", true, false},
		"core/core.go":                                                   {"grpc/core/core.tmpl", true, false},
		"core/models/models.go":                                          {"grpc/core/models/models.tmpl", true, false},
		"core/models/service.go":                                         {"grpc/core/models/service.tmpl", true, false},
		"core/models/user.go":                                            {"grpc/core/models/user.tmpl", true, false},
		"core/models/metrics.go":                                         {"grpc/core/models/metrics.tmpl", true, false},
		"core/models/context.go":                                         {"grpc/core/models/context.tmpl", true, false},
		"core/servant/dataware.go":                                       {"grpc/core/servant/dataware.tmpl", true, false},
		"core/servant/agent.go":                                          {"grpc/core/servant/agent.tmpl", true, false},
		"core/servant/business.go":                                       {"grpc/core/servant/business.tmpl", true, false},
		"Makefile":                                                       {"grpc/makefile.tmpl", true, false},
		"internal/insecure/server.go":                                    {"grpc/internal/insecure/server.tmpl", true, false},
		"internal/insecure/insecure.go":                                  {"grpc/internal/insecure/insecure.tmpl", true, false},
		"internal/insecure/client.go":                                    {"grpc/internal/insecure/client.tmpl", true, false},
		"internal/insecure/README.md":                                    {"grpc/internal/insecure/readme_md.tmpl", false, false},
		"internal/config/config.go":                                      {"grpc/internal/config/config.tmpl", true, false},
		"internal/config/app_conf.go":                                    {"grpc/internal/config/app_conf.tmpl", true, false},
		"internal/logus/logus.go":                                        {"grpc/internal/logus/logus.tmpl", true, false},
		"internal/rpcx/core.go":                                          {"grpc/internal/rpcx/core.tmpl", true, false},
		"internal/rpcx/rpcx.go":                                          {"grpc/internal/rpcx/rpcx.tmpl", true, false},
		"internal/rpcx/rx_affairs.go":                                    {"grpc/internal/rpcx/rx_affairs.tmpl", true, false},
		"internal/assets/assets.go":                                      {"grpc/internal/assets/assets.tmpl", true, false},
		"internal/assets/assets_gen.go":                                  {"grpc/internal/assets/assets_gen.tmpl", true, false},
		"hack/systemd/config/app.toml":                                   {"grpc/hack/systemd/config/app_toml.tmpl", false, false},
		"hack/systemd/README.md":                                         {"grpc/hack/systemd/readme_md.tmpl", false, false},
		"hack/systemd/zimlet-affairs-ws.service":                         {"grpc/hack/systemd/zimlet-affairs-ws_service.tmpl", false, false},
		"hack/systemd/zimlet-affairs.service":                            {"grpc/hack/systemd/zimlet-affairs_service.tmpl", false, false},
		"hack/systemd/zimlet-agent.service":                              {"grpc/hack/systemd/zimlet-agent_service.tmpl", false, false},
		"hack/systemd/zimlet-agent-ws.service":                           {"grpc/hack/systemd/zimlet-agent-ws_service.tmpl", false, false},
		"hack/goprotoc.sh":                                               {"grpc/hack/goprotoc_sh.tmpl", false, true},
		"hack/README.md":                                                 {"grpc/hack/readme_md.tmpl", false, false},
		"docs/README.md":                                                 {"grpc/docs/readme_md.tmpl", false, false},
		"README.md":                                                      {"grpc/readme_md.tmpl", false, false},
		".gitignore":                                                     {"grpc/.gitignore.tmpl", false, false},
		"version/version.go":                                             {"grpc/version/version.tmpl", true, false},
		"assets/config/app.toml":                                         {"grpc/assets/config/app_toml.tmpl", false, false},
		"assets/README.md":                                               {"grpc/assets/readme_md.tmpl", false, false},
		"assets/certs/server.pem":                                        {"grpc/assets/certs/server_pem.tmpl", false, false},
		"assets/certs/server.key":                                        {"grpc/assets/certs/server_key.tmpl", false, false},
	},
}