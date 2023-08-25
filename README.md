# solgate

Caddy based plugin to manage (path based) access to services on K8S

## Building

You first need to build a new caddy executable with this plugin. The easiest way is to do this with xcaddy.

Install xcaddy :

```shell
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
```

After xcaddy installation you can build caddy with this plugin by executing:

```shell
xcaddy build v2.6.4 --with github.com/loafoe/solgate
```

## Configuration

### Helm

Use the included `Helm` chart to deploy. Example `values.yaml`

```yaml
issuer: https://dex.hsp.hostedzonehere.com/

loki:
  url: loki-gateway.observability.svc
  
ingress:
  enabled: true
  className: "nginx"
  hosts:
    - host: solgate.test.hostedzonehere.com
      paths:
        - path: /
          pathType: ImplementationSpecific
```

Then deploy:

```shell
helm template solgate charts/solgate --skip-tests --values values.yaml|kubectl apply -f - -n solgate
```

Once deployed you can configure your Grafana Data source to point to `https://solgate.test.hostedzonehere.com`.

## License

Apache 2.0
