# Mono
### The Go framework for monolithic microservices

## Concept
This is a concept I've been working on that sounds bizarre, but I think makes plenty of sense in practice. I'm developing this framework alongside a personal project to determine if I'm insane or not.

The idea is this: deploy different services using the same binary. Mono accomplishes this by loading a different router depending on which service is being run. Let's explain with an example:
 
 You want to have 3 services: auth, users, and posts. These three services serve wildly different traffic throughput, so you want to scale them individually to maximize resources, however you're a young company, or you already have a traditional monolith running in production so you don't have the resources to set up the infrastructure needed for inter-service communication (circuit breaking, service discovery, etc.) You can use Mono to seperate these services in code, while still sharing code between them.
 
 When you write your application with Mono, you can then run the same binary in three different ways:
  - myapp auth
  - myapp users
  - myapp posts
  
 This will launch the same binary, but the Mono framework will load the router you defined in your code for that particular service. Now you can deploy 5 instances of auth and posts, but only 2 instances of users, or... anything you want. Docker makes this especially easy. Each service can easily import code from the others, negating the need for inter-service communication. You get the deployability of (micro)services, and the simplicity of a monolith. Cool!
 
## Details
Mono is compatible with Go v1.8+, and is currently very simple. I will be adding support for SSL/TLS very soon, along with an improved set of runtime options. I won't be writing docs just yet, as the framework will surely change rapidly. If you think I'm crazy, please open an issue and let me know why! If you choose to try it out, open an issue and tell me why I'm not! Pull requests are always welcome, and I'm always happy to chat on Twitter (@cohix)

## Benifits
  - Teams that own particular services will be forced to fix compatibility with new changes, as the "Monoserver" will not compile until code across the entire application is updated
  - Services can be developed in seperate repos with their own "test" server for development, and then imported into the "Monoserver" with 0 changes.

## Drawbacks
I'm fully aware that this will not be the golden egg that will fix the industry's issues with microservices, and I have an idea on a few reasons why:
- The ability to have different teams working independantly on different services and deploying their own code is somewhat lessened; builds of the "Monoserver" must be done more carefully to avoid conflicts and subtle implementation changes.
- The argument could be made that this defeats the purpose of service oriented architecture, but I see Mono as a middle ground between monoliths and microservices, not a replacement for either.
