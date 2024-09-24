# sys

* epoll
    - `epoll` is high-concurrency and I/O-bound applications where you need to handle many connections or file descriptors. Instead of continuously polling or logging, epoll allows you to be event-driven, responding only when needed. This improves scalability and performance in networking applications.
    - Used in webservers like in `nginx` and `nodejs`. 