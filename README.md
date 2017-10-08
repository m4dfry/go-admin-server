# go-webapp-template

A template for Go Web Applications using: 
 * [Gorilla Toolkit](http://www.gorillatoolkit.org/)
 * [Negroni](https://github.com/urfave/negroni)
 * [nnmware / DevOOPS Bootstrap 3 Admin theme](https://github.com/nnmware/devoops)

At this time, consider the project **TOTALLY WIP**

You can login with user:admin pass:123456 loading the standard template.

---

The idea is to work with plug-in that users can compile and share with pages to include on the template.

**NB**
Unfortunatly in version 1.8 Go plugin only works on Linux.

---
- To do ASAP
  - some API must check the login first
  - FIX info in index template
  - ADD option on config to disable login
  - ADD option OnlyLocal for server

- To be added
  - [Sanitize input](https://github.com/kennygrant/sanitize)
  - Local Authentication with MongoDB
  - Enable text compression
  - Unit test

- To FIX
  - Fail to serve favicon.ico(legit, doesn't exist), server return 307 Temporary Redirect (wrong, must return 404)
  - Examine Chrome warning: [Deprecation] Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience.

- Controversial choices
  - Token of CookieStore is regenerated randomly at each start


## Demo

You can see the theme at his full potential with a checkout of demo branch.
On the main branch I will strip all the unecessary, in order to get the best reduced version of all elementary functionality

**Why a branch and not a tag ?**

Beacause I see future improvement (and a lot of bugs) even on demo


## License

Used software (all software is licensed under appropriate license of authors)
* DevOOPS Bootstrap 3 Admin theme [https://github.com/nnmware/devoops](https://github.com/nnmware/devoops)  v2 MIT
* Gorilla web toolkit [http://www.gorillatoolkit.org/](http://www.gorillatoolkit.org/) BSD 3-clause "New" or "Revised" License
* Negroni [https://github.com/urfave/negroni](https://github.com/urfave/negroni) 0.2.0 - 2016-05-10 MIT 

---

Have a request, suggestion , critic or question? Open an Issue!

---
