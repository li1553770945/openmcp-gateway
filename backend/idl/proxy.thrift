namespace go proxy

struct ProxyRequest {}

struct ProxyResponse {}

service ProxyService {
    ProxyResponse Forward(1: ProxyRequest req) (api.any="/proxy/*path");
}
