<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title>{{ .Title }}</title>
        <link
          rel="stylesheet"
          type="text/css"
          href="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui.css"
        />
        <link rel="icon" type="image/png" href="https://github.com/swagger-api/swagger-ui/blob/master/dist/favicon-16x16.png?raw=true" />
    </head>
    <body>
        <div id="swagger-ui"></div>

        <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui-bundle.js"></script>
        <script>
            // init Swagger for faucet's openapi.yml.
            window.onload = function() {
              window.ui = SwaggerUIBundle({
                url: {{ .URL }},
                dom_id: "#swagger-ui",
                deepLinking: true,
                layout: "BaseLayout",
              });
            }
        </script>
    </body>
</html>