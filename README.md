# go-webapp-template
===========================

A template for Go Web Applications using: 
 * [Gorilla Toolkit](http://www.gorillatoolkit.org/)
 * [Negroni](https://github.com/urfave/negroni)
 * [nnmware / DevOOPS Bootstrap 3 Admin theme](https://github.com/nnmware/devoops)

At this time, consider the project TOTALLY WIP
You can login with user:admin pass:123456

===========================

- To be added
  - Add libs license in README (show all the license neeeded in the right way)
  - [Sanitize input](https://github.com/kennygrant/sanitize)
  - Local Authentication with MongoDB
  - Enable text compression
  - Unit test

- To FIX
  - Fail to serve favicon.ico(legit, doesn't exist), server return 307 Temporary Redirect (wrong, must return 404)
  - Examine Chrome warning: [Deprecation] Synchronous XMLHttpRequest on the main thread is deprecated because of its detrimental effects to the end user's experience.

- Controversial choices
  - Token of CookieStore is regenerated randomly at each start

- Ideas
  - Add pages to site using Go Plugin (is it possible ??)
---

Have a request, suggestion , critic or question? Open an Issue!

---
