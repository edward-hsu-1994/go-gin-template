package domain

import "github.com/joomcode/errorx"

var namespace = errorx.NewNamespace("go-fiber-template")

var (
	ErrorRoot              = errorx.NewType(namespace, "error")
	ErrorServerSideRoot    = ErrorRoot.NewSubtype("server_side_error")
	ErrorClientSideRoot    = ErrorRoot.NewSubtype("client_side_error")
	ErrorBusinessLogicRoot = ErrorRoot.NewSubtype("business_logic_error")
	ErrorDependencyRoot    = ErrorRoot.NewSubtype("dependency_error")
	ErrorSystemRoot        = ErrorRoot.NewSubtype("system_error")
	ErrorSecurityRoot      = ErrorRoot.NewSubtype("security_error")
)

var ErrorHttpStatusProperty = errorx.RegisterProperty("http_status")

var (
	ErrorInvalidPostStatus = ErrorClientSideRoot.NewSubtype("invalid_post_status")
	ErrorPostNotFound      = ErrorClientSideRoot.NewSubtype("post_not_found")
)
