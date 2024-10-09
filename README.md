# Word of Wisdom

### About

The repository contains PoC of a server with DDOS protection. The protection is implemented using proof-of-work concept
and in particular [Hashcash](https://ru.wikipedia.org/wiki/Hashcash) check. The algorithm is quite famous, and its
difficulty can be easily adjusted with just one parameter, i.e. the number of leading zeroes.

### Launch

Simply run `docker-compose up` in the project's root directory. \
This launches the Word of Wisdom server and a few clients.

### Sample output

```log
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:25 [INFO] initialize config
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:25 [INFO] connect to server:9999
word-of-wisdom-client-hard-1  | 2024/10/09 14:46:25 [INFO] initialize config
word-of-wisdom-client-hard-1  | 2024/10/09 14:46:25 [INFO] connect to server:9999
word-of-wisdom-server-1       | 2024/10/09 14:46:25 [INFO] initialize config
word-of-wisdom-server-1       | 2024/10/09 14:46:25 [INFO] initialize repositories
word-of-wisdom-server-1       | 2024/10/09 14:46:25 [INFO] initialize services
word-of-wisdom-server-1       | 2024/10/09 14:46:25 [INFO] start server on :9999
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] got challenge of difficulty: 20
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-client-hard-1  | 2024/10/09 14:46:26 [INFO] got challenge of difficulty: 24
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] challenge client: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::nNMWm6mS:4b958
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] remember carefully: На покляпое дерево и козы скачут.
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] got challenge of difficulty: 20
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::pKAbHcaH:4471b
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] remember carefully: Пишет, словно разводы разводит (крупно и медленно).
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] got challenge of difficulty: 20
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] remember carefully: Не черт тебя нес на худой на мост.
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:26 [INFO] got challenge of difficulty: 20
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::GnKC0wOO:2d0
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:26 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:28 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::AYn1ohTO:1e553a
word-of-wisdom-server-1       | 2024/10/09 14:46:28 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:28 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:28 [INFO] remember carefully: Никола в путь, Христос по дорожке (отплывающим на судах).
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:28 [INFO] got challenge of difficulty: 20
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:29 [INFO] remember carefully: Пошел черных кобелей набело перемывать.
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:29 [INFO] got challenge of difficulty: 20
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::nvik8NBH:12fbc8
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] check solution: 1:20:241009:c442bad7-02ec-4637-a25a-1e21b1651351::TE9gUTgy:26170
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] send word to: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:29 [INFO] remember carefully: Пришли казаки с Дону, да прогнали ляхов до (к) дому.
word-of-wisdom-client-easy-1  | 2024/10/09 14:46:29 [INFO] got challenge of difficulty: 20
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] challenge client: c442bad7-02ec-4637-a25a-1e21b1651351
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] check solution: 1:24:241009:2b1d273c-aca3-4b79-b44a-6221cf60c6af::xD6a4Ni2:41acd9
word-of-wisdom-server-1       | 2024/10/09 14:46:29 [INFO] send word to: 2b1d273c-aca3-4b79-b44a-6221cf60c6af
word-of-wisdom-client-hard-1  | 2024/10/09 14:46:29 [INFO] remember carefully: И не хочет коза на базар, да ведут за рога.
word-of-wisdom-client-hard-1  | 2024/10/09 14:46:29 [INFO] got challenge of difficulty: 24
```