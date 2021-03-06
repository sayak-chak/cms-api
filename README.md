# CMS (Content Management System) API

## Features

- Lightweight, purely in golang
- Modular monolith CMS API - modules can be easily extended into microservices if needed
- Highly customizable & extensible
- Built to scale - [performace benchmarks](https://github.com/sayak-chak/cms-api/blob/master/load_tests/test_results.md)

## Quickstart

###### Bring up the service

`make start-service` and post its completion, `make start-server`

```
PS: The DB is seeded with some dummy data when the service is started locally
```

###### Bring down the service

`make stop-service`

## Overview

| Directory | Description                                                                                                                                                                                   |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Config    | All the necessary configurations                                                                                                                                                              |
| Database  | Generic package defining the database layer as an interface. Can be implemented by any database of choice in plug and play fashion (by default, has Postgres implementing the database layer) |
| Modules   | Service modules. Each module here can be potentially extrapolated out into a micro service if needed                                                                                          |
| Models    | All the request, response & database models                                                                                                                                                   |
| Server    | Starting point for server                                                                                                                                                                     |
| Utils     | Some utility features - example, logging db queries                                                                                                                                           |

## Endpoints

| Endpoint            | Description                                                                                           |
| ------------------- | ----------------------------------------------------------------------------------------------------- |
| /register           | Register (as an author) once an otp has been sent to the mobile number that would be used to register |
| /login              | Login (as an author) with mobile number & password                                                    |
| /add-content        | Add new content                                                                                       |
| /top-contents       | Get list of top few interesting contents sorted by votes                                              |
| /top-contents/tag   | Get list of top few interesting content by the provided tag, sorted by votes                          |
| /content/content-id | Get the content with the specified content id                                                         |
| /upvote             | Upvote a specific content                                                                             |
| /subscribe          | Subscribe to get updates in email                                                                     |
| /unsubscribe        | Unsubscribe from the service - should be sent alongwith email updates                                 |
| /authors            | Get list of top few content creators                                                                  |

---

### /register

###### Request

```

{
    otp: string,
    password: string,
    mobile: number,
    name: string,
}

```

###### Response

_Success status - 201_

---

### /login

###### Request

```

{
    mobile: number,
    password: string,
}

```

###### Response

_Success status - 201_

```

{
    token: string,
    id: int, // author id
}

PS: This is a jwt & should be sent when a new post is to be added, not doing so would result in an authorization failure

```

---

### /add-content

###### Request

_Authorization bearer token should be sent with headers_

```

{
    tags: [string],
    title: string,
    summary: string,
    body: string, // main content body
    imageSrc: string, // url to image used
    authorId: int
}

```

###### Response

_Success status - 200_

---

### /top-contents (& /top-contents/tag)

###### Response

```

[
    {
        author: string, // name
        contentId: string,
        authorId: string,
        body: string, // content body
        imageSrc: string, // url to image used
        title: string,
        summary: string,
        votes: number
    }
]

```

`PS: Ideally this should be cached on client side to reduce load on server when a specific content is requested`

---

### /content/content-id

###### Response

```

{
    author: string, // name
    contentId: string,
    authorId: string,
    body: string, // content body
    imageSrc: string, // url to image used
    title: string,
    summary: string,
    votes: number
}

```

---

### /upvote

###### Request

```

{
    contentId: int
}

```

###### Response

_Success status - 200_

---

### /subscribe

###### Request

```

{
    email: string
}

```

###### Response

_Success status - 204_

---

### /unsubscribe

###### Request

```

{
    email: string
}

```

###### Response

_Success status - 204_

---

### /authors

###### Response

```

{
    [
        {
            name: string
        }
    ]

}

```

---

## Customization

- By default, Postgres is used, but any database can work with this by implementing the `Database` interface. The new database directory should be placed in the same level as `postgres` directory. Also, the `database model` might have be updated depending on the database. No other change is needed.

- Mobile number (wherein otp has to be sent pre-registering) is used as the identifier for creating a new account. This can be changed to email/username/anything else by updating the `AuthorCreds` & `AuthorTempCreds` tables.

- Right now the subscribers can only view the contents and the authors can post/view them, the subscibers can be removed if aiming for a model without distinct author/subscriber enitites. That way, everyone can post and view.

- By default, there are some predefined tags based on which filtering of content can be done. To add a new tag, there needs to be another tag table containing the relevant content ids & `tagNamesMap` attribute has to be updated.

- To have an even more optimized filtering, tags can also be created on the go by doing some content analysis.

- For optimized performance, the server caches frequent requests for some time, the cache time should be set to 0 to remove this.

- By default, otp is valid for 6 hours, the `otpExpiryPeriod` needs to be updated to change this.

- JWT authetication is used, but the api is flexible enough to work with any auth system - just the appropriate auth middleware needs to be updated.

```

```
