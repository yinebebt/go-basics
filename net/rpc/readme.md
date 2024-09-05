# net

## rpc
* JSON-RPC makes it easy to communicate with a service written in RPC using any programming language.

* You can test server response by sending a request using curl:

    ```bash
    Copy code
    curl -X POST \
    http://localhost:9000/rpc \
    -H 'Cache-Control: no-cache' \
    -H 'Content-Type: application/json' \
    -d '{
        "method": "JSONServer.TwitterProfileDetail",
        "params": [
            {
                "id": 1
            }
        ],
        "id": 1
    }'
    ```

* In JSON-RPC, the request typically includes the fields method, params, and id:
    - **method**: The name of the method to be invoked on the server.
    - **params**: An array or object containing the parameters to be passed to the method.
    - **id**: A unique identifier used to match the response with the request, which is particularly useful when multiple requests are sent simultaneously, as the order of responses may vary.