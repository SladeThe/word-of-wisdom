# Word of Wisdom
A test task for Server Engineer

### About

The repository contains PoC of a server with DDOS protection. The protection is implemented using proof-of-work concept
and in particular [Hashcash](https://ru.wikipedia.org/wiki/Hashcash) check. The algorithm is quite famous, and its
difficulty can be easily adjusted with just one parameter, i.e. the number of leading zeroes.

### Launch

Simply run `docker-compose up` in the project's root directory. \
This launches the Word of Wisdom server and a few clients.

### Sample output

```log
word-of-wisdom-client-2  | 2024/10/07 00:24:59 [INFO] initialize config
word-of-wisdom-client-2  | 2024/10/07 00:24:59 [INFO] connect to server:9999
word-of-wisdom-client-1  | 2024/10/07 00:24:59 [INFO] initialize config
word-of-wisdom-client-1  | 2024/10/07 00:24:59 [INFO] connect to server:9999
word-of-wisdom-server-1  | 2024/10/07 00:24:59 [INFO] initialize config
word-of-wisdom-server-1  | 2024/10/07 00:24:59 [INFO] initialize repositories
word-of-wisdom-server-1  | 2024/10/07 00:24:59 [INFO] initialize services
word-of-wisdom-server-1  | 2024/10/07 00:24:59 [INFO] start server on :9999
word-of-wisdom-server-1  | 2024/10/07 00:25:00 [INFO] challenge client: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-server-1  | 2024/10/07 00:25:00 [INFO] challenge client: 3842e086-6e14-404b-a4d8-b60ca88368d8
word-of-wisdom-server-1  | 2024/10/07 00:25:00 [INFO] check solution: 1:20:241007:3842e086-6e14-404b-a4d8-b60ca88368d8::z0KyfrLm:5ebab
word-of-wisdom-server-1  | 2024/10/07 00:25:00 [INFO] send word to: 3842e086-6e14-404b-a4d8-b60ca88368d8
word-of-wisdom-server-1  | 2024/10/07 00:25:00 [INFO] challenge client: 3842e086-6e14-404b-a4d8-b60ca88368d8
word-of-wisdom-client-1  | 2024/10/07 00:25:00 [INFO] remember carefully: Каковы гости, таков и пир.
word-of-wisdom-server-1  | 2024/10/07 00:25:01 [INFO] check solution: 1:20:241007:2b1d273c-aca3-4b79-b44a-6221cf60c6af::sX16BWH0:123bfe
word-of-wisdom-server-1  | 2024/10/07 00:25:01 [INFO] send word to: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-server-1  | 2024/10/07 00:25:01 [INFO] challenge client: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-client-2  | 2024/10/07 00:25:01 [INFO] remember carefully: Ехал к Фоме, а заехал к куме.
word-of-wisdom-server-1  | 2024/10/07 00:25:03 [INFO] check solution: 1:20:241007:2b1d273c-aca3-4b79-b44a-6221cf60c6af::ptvr1owv:29aa6c
word-of-wisdom-client-2  | 2024/10/07 00:25:03 [INFO] remember carefully: На балалайку станет, и на кабак станет, а на свечку не станет.
word-of-wisdom-server-1  | 2024/10/07 00:25:03 [INFO] send word to: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-server-1  | 2024/10/07 00:25:03 [INFO] challenge client: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-server-1  | 2024/10/07 00:25:04 [INFO] check solution: 1:20:241007:3842e086-6e14-404b-a4d8-b60ca88368d8::Mt4OnNO0:4a7b0a
word-of-wisdom-client-1  | 2024/10/07 00:25:04 [INFO] remember carefully: У бабы семьдесят две увертки в день.
```