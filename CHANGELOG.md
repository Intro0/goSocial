# Changelog

## [1.1.0](https://github.com/Intro0/goSocial/compare/v1.0.0...v1.1.0) (2026-10-06)


### Features

* add API rate limiting ([d862d5b](https://github.com/Intro0/goSocial/commit/d862d5baf600d4111cdb87f0819e005ca44f2c82))
* add API rate limiting ([6b4a7f3](https://github.com/Intro0/goSocial/commit/6b4a7f30ace7f8865a62283af72591cae0d5cefe))
* add basic auth middleware ([85fdb1c](https://github.com/Intro0/goSocial/commit/85fdb1c105e5dfea57ebf8dc8c9874ddc39a37a4))
* add comments to post responses ([110fa42](https://github.com/Intro0/goSocial/commit/110fa42b86c38a04d6276230de49061142a85302))
* add configurable CORS policy ([92f8220](https://github.com/Intro0/goSocial/commit/92f8220507b8ba5480e9653c25bea51980140941))
* add configurable CORS policy ([fbea6dc](https://github.com/Intro0/goSocial/commit/fbea6dc2e6b8bca404fc9c92e169bef4e11817fc))
* add database connection pool ([dd78a4b](https://github.com/Intro0/goSocial/commit/dd78a4b77da195011d3621dd786bc44e7dcba09d))
* add database seeding command ([82a5b5c](https://github.com/Intro0/goSocial/commit/82a5b5c4649a25e3a8032c4544aef04dc307d3d0))
* add environment configuration helpers ([9c51fa1](https://github.com/Intro0/goSocial/commit/9c51fa1775b17bb26ffa9203370a2e62e6aca95b))
* add feed search and filtering ([1c10438](https://github.com/Intro0/goSocial/commit/1c1043808c6da379645ad9f61c5867ce83b4852f))
* add fixed window rate limiter ([eca3104](https://github.com/Intro0/goSocial/commit/eca3104515094f2103a53f695b8bae9bec1a0afe))
* add generated Swagger documentation ([440dbb6](https://github.com/Intro0/goSocial/commit/440dbb6a7c3f2e6ab260f9bb876dbb60d63e4df3))
* add generated Swagger documentation ([435867b](https://github.com/Intro0/goSocial/commit/435867b76ebff05b43986bc143337efc6b70cc0b))
* add graceful server shutdown ([fa4723a](https://github.com/Intro0/goSocial/commit/fa4723aa9ffc27f2bbe473fe9bfb170bc9ee2d9d))
* add graceful server shutdown ([01a55d2](https://github.com/Intro0/goSocial/commit/01a55d2b71bb1722c7767030879584a6327271e1))
* add HTTP API with health endpoint ([0ad5091](https://github.com/Intro0/goSocial/commit/0ad5091059334b2b54f9b19c0d4126e5c726aeb5))
* add invitation-based user registration ([cafd076](https://github.com/Intro0/goSocial/commit/cafd076991ef8b4c8c85f9c782da4cfe96a000d9))
* add JSON response and error helpers ([6a8e8fd](https://github.com/Intro0/goSocial/commit/6a8e8fd63c7815179c33eda3907022e40f26a988))
* add JWT authentication ([52656e6](https://github.com/Intro0/goSocial/commit/52656e6746180edbec9ddd98177e880db7dde1f1))
* add optimistic concurrency control to post updates ([53a93b1](https://github.com/Intro0/goSocial/commit/53a93b178007e47a5bc16ee82d6c5452d4bdfc47))
* add paginated user feed ([f112389](https://github.com/Intro0/goSocial/commit/f11238944d167d052506050493613c211bef145e))
* add post creation endpoint with validation ([33e15b8](https://github.com/Intro0/goSocial/commit/33e15b89216005dd068f09105ac4a502cd5b494b))
* add post retrieval endpoint ([490007c](https://github.com/Intro0/goSocial/commit/490007c34a0ca61e5633b63a36378ad92357ea5d))
* add post update and delete endpoints ([e877897](https://github.com/Intro0/goSocial/commit/e8778975561bfbd165e888d6b3d63c64bfed9a38))
* add PostgreSQL post and user stores ([c65e79c](https://github.com/Intro0/goSocial/commit/c65e79c25a48541bb53d51125600460dfb27aaaf))
* add Redis configuration ([da04be5](https://github.com/Intro0/goSocial/commit/da04be5b136f8f12362403b2f8a199f47a816964))
* add Redis user cache storage ([1123d3f](https://github.com/Intro0/goSocial/commit/1123d3fcd4dfa9eb2a9527f35c979d7c2d1c43db))
* add Redis user caching ([d469850](https://github.com/Intro0/goSocial/commit/d469850312f86a73759deb41daf6bd7163dc69f4))
* add role-based authorization ([316edb5](https://github.com/Intro0/goSocial/commit/316edb5c44218a43bfad606221da574dc6afd474))
* add SendGrid invitation email integration ([bdd3d78](https://github.com/Intro0/goSocial/commit/bdd3d78845286d2a324182d2fed14cdee1a3f09c))
* add SQL migrations ([7211dbc](https://github.com/Intro0/goSocial/commit/7211dbccbd015103f614ee609398add1eeae22d4))
* add structured logging ([bba9e13](https://github.com/Intro0/goSocial/commit/bba9e13dd5fc4032faacd77bde75fa75c3e07d96))
* add timeouts to database queries ([82ef4fb](https://github.com/Intro0/goSocial/commit/82ef4fba18663ce3608553016b88f054a3099925))
* add user feed endpoint ([8f0d2d2](https://github.com/Intro0/goSocial/commit/8f0d2d2d4af70af9ec645ff33a605a0b1b0cd703))
* add user follow and unfollow endpoints ([6ae464b](https://github.com/Intro0/goSocial/commit/6ae464b2db59cc8bc2479fdd7fd12f71a92adaa0))
* add user retrieval endpoint ([f8c5d11](https://github.com/Intro0/goSocial/commit/f8c5d11d257ae33e090f15a233488414b848d31c))
* cache user lookups ([f9fe092](https://github.com/Intro0/goSocial/commit/f9fe092af3573469ca7c85d68e783114819d4990))
* expose protected server metrics ([0e27770](https://github.com/Intro0/goSocial/commit/0e277700c25e503aa8894169fe3d6749b1d4918d))
* expose protected server metrics ([8b4ebdc](https://github.com/Intro0/goSocial/commit/8b4ebdcb74bc3e6f0bcfa5abe72642b24f61ec05))
* standardize JSON response ([550e636](https://github.com/Intro0/goSocial/commit/550e6364c5b36a71c667a94c000c3869ea2381ca))
* wire Redis cache client ([297380d](https://github.com/Intro0/goSocial/commit/297380dc5486dbd161733b37a294dbe0e9a9e5da))


### Bug Fixes

* allow PATCH requests through CORS ([a0c7a76](https://github.com/Intro0/goSocial/commit/a0c7a765ee271dc0d1339444e52237502e024de2))
* harden Redis user caching ([ec848a6](https://github.com/Intro0/goSocial/commit/ec848a60ffbe2f8fff16360f4a547214199eac0a))
* make metrics registration idempotent ([bd51e2f](https://github.com/Intro0/goSocial/commit/bd51e2fffd0ea15f4b35277297bed32bb121d9a6))
* remove unused user context middleware ([d1a0f0c](https://github.com/Intro0/goSocial/commit/d1a0f0caf0a3f3f68908e23983bf5373424fc511))


### Performance Improvements

* add database indexes for post, user, and comment queries ([60f9938](https://github.com/Intro0/goSocial/commit/60f99386369d76b4fe545dc7a237347874d2a446))

## 1.0.0 (2026-10-06)


### Features

* add API rate limiting ([d862d5b](https://github.com/Intro0/goSocial/commit/d862d5baf600d4111cdb87f0819e005ca44f2c82))
* add API rate limiting ([6b4a7f3](https://github.com/Intro0/goSocial/commit/6b4a7f30ace7f8865a62283af72591cae0d5cefe))
* add basic auth middleware ([85fdb1c](https://github.com/Intro0/goSocial/commit/85fdb1c105e5dfea57ebf8dc8c9874ddc39a37a4))
* add comments to post responses ([110fa42](https://github.com/Intro0/goSocial/commit/110fa42b86c38a04d6276230de49061142a85302))
* add configurable CORS policy ([92f8220](https://github.com/Intro0/goSocial/commit/92f8220507b8ba5480e9653c25bea51980140941))
* add configurable CORS policy ([fbea6dc](https://github.com/Intro0/goSocial/commit/fbea6dc2e6b8bca404fc9c92e169bef4e11817fc))
* add database connection pool ([dd78a4b](https://github.com/Intro0/goSocial/commit/dd78a4b77da195011d3621dd786bc44e7dcba09d))
* add database seeding command ([82a5b5c](https://github.com/Intro0/goSocial/commit/82a5b5c4649a25e3a8032c4544aef04dc307d3d0))
* add environment configuration helpers ([9c51fa1](https://github.com/Intro0/goSocial/commit/9c51fa1775b17bb26ffa9203370a2e62e6aca95b))
* add feed search and filtering ([1c10438](https://github.com/Intro0/goSocial/commit/1c1043808c6da379645ad9f61c5867ce83b4852f))
* add fixed window rate limiter ([eca3104](https://github.com/Intro0/goSocial/commit/eca3104515094f2103a53f695b8bae9bec1a0afe))
* add generated Swagger documentation ([440dbb6](https://github.com/Intro0/goSocial/commit/440dbb6a7c3f2e6ab260f9bb876dbb60d63e4df3))
* add generated Swagger documentation ([435867b](https://github.com/Intro0/goSocial/commit/435867b76ebff05b43986bc143337efc6b70cc0b))
* add graceful server shutdown ([fa4723a](https://github.com/Intro0/goSocial/commit/fa4723aa9ffc27f2bbe473fe9bfb170bc9ee2d9d))
* add graceful server shutdown ([01a55d2](https://github.com/Intro0/goSocial/commit/01a55d2b71bb1722c7767030879584a6327271e1))
* add HTTP API with health endpoint ([0ad5091](https://github.com/Intro0/goSocial/commit/0ad5091059334b2b54f9b19c0d4126e5c726aeb5))
* add invitation-based user registration ([cafd076](https://github.com/Intro0/goSocial/commit/cafd076991ef8b4c8c85f9c782da4cfe96a000d9))
* add JSON response and error helpers ([6a8e8fd](https://github.com/Intro0/goSocial/commit/6a8e8fd63c7815179c33eda3907022e40f26a988))
* add JWT authentication ([52656e6](https://github.com/Intro0/goSocial/commit/52656e6746180edbec9ddd98177e880db7dde1f1))
* add optimistic concurrency control to post updates ([53a93b1](https://github.com/Intro0/goSocial/commit/53a93b178007e47a5bc16ee82d6c5452d4bdfc47))
* add paginated user feed ([f112389](https://github.com/Intro0/goSocial/commit/f11238944d167d052506050493613c211bef145e))
* add post creation endpoint with validation ([33e15b8](https://github.com/Intro0/goSocial/commit/33e15b89216005dd068f09105ac4a502cd5b494b))
* add post retrieval endpoint ([490007c](https://github.com/Intro0/goSocial/commit/490007c34a0ca61e5633b63a36378ad92357ea5d))
* add post update and delete endpoints ([e877897](https://github.com/Intro0/goSocial/commit/e8778975561bfbd165e888d6b3d63c64bfed9a38))
* add PostgreSQL post and user stores ([c65e79c](https://github.com/Intro0/goSocial/commit/c65e79c25a48541bb53d51125600460dfb27aaaf))
* add Redis configuration ([da04be5](https://github.com/Intro0/goSocial/commit/da04be5b136f8f12362403b2f8a199f47a816964))
* add Redis user cache storage ([1123d3f](https://github.com/Intro0/goSocial/commit/1123d3fcd4dfa9eb2a9527f35c979d7c2d1c43db))
* add Redis user caching ([d469850](https://github.com/Intro0/goSocial/commit/d469850312f86a73759deb41daf6bd7163dc69f4))
* add role-based authorization ([316edb5](https://github.com/Intro0/goSocial/commit/316edb5c44218a43bfad606221da574dc6afd474))
* add SendGrid invitation email integration ([bdd3d78](https://github.com/Intro0/goSocial/commit/bdd3d78845286d2a324182d2fed14cdee1a3f09c))
* add SQL migrations ([7211dbc](https://github.com/Intro0/goSocial/commit/7211dbccbd015103f614ee609398add1eeae22d4))
* add structured logging ([bba9e13](https://github.com/Intro0/goSocial/commit/bba9e13dd5fc4032faacd77bde75fa75c3e07d96))
* add timeouts to database queries ([82ef4fb](https://github.com/Intro0/goSocial/commit/82ef4fba18663ce3608553016b88f054a3099925))
* add user feed endpoint ([8f0d2d2](https://github.com/Intro0/goSocial/commit/8f0d2d2d4af70af9ec645ff33a605a0b1b0cd703))
* add user follow and unfollow endpoints ([6ae464b](https://github.com/Intro0/goSocial/commit/6ae464b2db59cc8bc2479fdd7fd12f71a92adaa0))
* add user retrieval endpoint ([f8c5d11](https://github.com/Intro0/goSocial/commit/f8c5d11d257ae33e090f15a233488414b848d31c))
* cache user lookups ([f9fe092](https://github.com/Intro0/goSocial/commit/f9fe092af3573469ca7c85d68e783114819d4990))
* expose protected server metrics ([0e27770](https://github.com/Intro0/goSocial/commit/0e277700c25e503aa8894169fe3d6749b1d4918d))
* expose protected server metrics ([8b4ebdc](https://github.com/Intro0/goSocial/commit/8b4ebdcb74bc3e6f0bcfa5abe72642b24f61ec05))
* standardize JSON response ([550e636](https://github.com/Intro0/goSocial/commit/550e6364c5b36a71c667a94c000c3869ea2381ca))
* wire Redis cache client ([297380d](https://github.com/Intro0/goSocial/commit/297380dc5486dbd161733b37a294dbe0e9a9e5da))


### Bug Fixes

* allow PATCH requests through CORS ([a0c7a76](https://github.com/Intro0/goSocial/commit/a0c7a765ee271dc0d1339444e52237502e024de2))
* harden Redis user caching ([ec848a6](https://github.com/Intro0/goSocial/commit/ec848a60ffbe2f8fff16360f4a547214199eac0a))
* make metrics registration idempotent ([bd51e2f](https://github.com/Intro0/goSocial/commit/bd51e2fffd0ea15f4b35277297bed32bb121d9a6))
* remove unused user context middleware ([d1a0f0c](https://github.com/Intro0/goSocial/commit/d1a0f0caf0a3f3f68908e23983bf5373424fc511))


### Performance Improvements

* add database indexes for post, user, and comment queries ([60f9938](https://github.com/Intro0/goSocial/commit/60f99386369d76b4fe545dc7a237347874d2a446))
