package server

import (
	"net/http"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []route

func (r *Router) initRoutes() {
	r.routes = routes{
		route{
			"Ping", "GET", "/ping", r.ping,
		},
		route{
			"EndpointCreate", "POST", "/endpoint/{endpointID}", r.endpointCreate,
		},
		route{
			"EndpointDelete", "DELETE", "/endpoint/{endpointID}", r.endpointDelete,
		},
		route{
			"EndpointGet", "GET", "/endpoint/{endpointID}", r.endpointGet,
		},
		route{
			"EndpointsGet", "GET", "/endpoints", r.endpointsGet,
		},
		route{
			"EndpointUpdate", "POST", "/endpoint/update/{endpointID}", r.endpointUpdate,
		},
		route{
			"AllocateIPv6", "POST", "/allocator/ipam-allocate/{ipam-type}", r.allocateIPv6,
		},
		route{
			"ReleaseIPv6", "POST", "/allocator/ipam-release/{ipam-type}", r.releaseIPv6,
		},
		route{
			"GetLabels", "GET", "/labels/by-uuid/{uuid}", r.getLabels,
		},
		route{
			"GetLabelsBySHA256", "GET", "/labels/by-sha256sum/{sha256sum}", r.getLabelsBySHA256,
		},
		route{
			"PutLabels", "POST", "/labels", r.putLabels,
		},
		route{
			"DeleteLabelsBySHA256", "DELETE", "/labels/by-sha256sum/{sha256sum}", r.deleteLabelsBySHA256,
		},
		route{
			"DeleteLabelsByUUID", "DELETE", "/labels/by-uuid/{uuid}", r.deleteLabelsByUUID,
		},
		route{
			"GetMaxID", "GET", "/labels/status/maxUUID", r.getMaxUUID,
		},
		route{
			"PolicyAdd", "POST", "/policy/{path}", r.policyAdd,
		},
		route{
			"PolicyDelete", "DELETE", "/policy/{path}", r.policyDelete,
		},
		route{
			"PolicyGet", "GET", "/policy/{path}", r.policyGet,
		},
		route{
			"PolicyCanConsume", "POST", "/policy-consume-decision", r.policyCanConsume,
		},
	}
}
