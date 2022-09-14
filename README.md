# proxyHeaders

proxyHeaders is a middleware plugin for traefik. It forwards original header.

## Configuration

### Static

```yaml
experimental:
  plugins:
    proxyHeaders:
      modulename: "github.com/Hvid/proxyHeaders"
      version: "v0.0.7"
```

### Dynamic

```yaml
http:
    middlewares:
        proxyHeaders:
            enable: true
```