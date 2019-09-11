# Login
Clearly, we can't have a web system without a login!

I really hate managing sessions, and I think they're kind of an outdated concept, same with usernames and passwords.

"Sessions" are actually instead authorization credentials, and the process of login is the process of getting an authorization credential for an authentication credential.

```mermaid
sequenceDiagram
    participant User
    participant Server
    participant Client

    User ->> Client: /login
    Client ->> Server: valid authn for this user?
    Note right of Server: Server replies with <br/> redirect uris
    Server -->> Client: [Google, Facebook, ...]
    Client -->> User: sign in via google, facebook please!
    User ->> Client: Click: "login via google"
    Client ->> Client: Perform google auth flow 
    Client ->> Server: Authn(ID Token, scopes: [...])
    Server ->> Client: Authorization Token ([scopes])
    Client ->> User: [auhorized page]
```

This authorization token is then essentially the session.