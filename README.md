# An exercise to build a succinct golang starting template for micro-services, cli and web-assembly

## Read more thoughts and discussion

To make the README shorter to discuss and document, I respectively created the [gogo-ikimasu board](https://github.com/arnos/gogo-ikimasu/discussions) and the [gogo-ikimasu wiki](https://github.com/arnos/gogo-ikimasu/wiki)

##  The Goal

The Goal is to provide a simple template, inspired by clean Architecture and screaming Architecture.
This comes back to a principle of convention over configuration as well, writing/refactoring code or reviewing code should not be seen as shore or rubber stamp approve.
It should create a common framework which then leads discussion on software quality.

1. You shouldn't need to open file to understand what their function is.
2. Files should have a standard naming format.   
    as such I propose the following standard:  
    > 
    > [nano-domain]_[purpose].[language extension]
    >
    In the golang context I propose the following purposes:
    - _main (where everything starts and ends)
    - _config (where we setup everything).
    - _logic (how everything works)
    - _interfaces (how to abstract implementation from the logic) 
    - _implementation (where everything happens)
    - _test (because it inspired the template).
3. Having the above purposes allows one to quickly find the proper file for a given piece of code 
4. Then comes the domain, use a domain that draws the attention of what your code is doing
    I.E. the file domain should tell what the code purpose is, here I've done a rōmaji Hello (Harō) World (Wārudo) creating 2 "nano-domains" Harō for the execution and Wārudo for the implementation (NB. utilities_ is a special nano-domains similar to the "util" packages in java)

    Here we see 4 Harō "nano-domain" files
    - **Hello_main.go**: is where the execution starts and we create our main struct for our code
    - **Hello_entities.go**: host the interfaces, structs and constants that are required by the logic/execution (in this template only a simple "hello" interface)
    - **Hello_logic.go**: is home of the logic, it relies on **Hello_entities.go** to provide it with interface, constants and structs (mainly if you need methods vs functions) to seperate logic from implementation
    - **Hello_test.go**: while empty here, it would be you would have your test against the logic code "aka bdd testing" 

    For the implementation we see 3 Wārudo "nano-domain" files
    - **World_entities.go**: host the interfaces, structs and constants that are required by the implementation (in this template only a simple "world" struct)
    - **World_implementation.go**: methods that implement the interfaces in **Hello_entities.go** and functions to support these
    - **World_test.go**: the core purpose is for unit tests for this "nano-domain"

    NB. each "nano-domain" can contain any combination of [nano-domain]_[purpose].go files as required by your design, after all this is just a Hello World :)
5. The domains should be "nano-domain" and only be used when needed to seperate code into smaller more managable junks.

With the above file name format I submit that we are using screaming architecture

Clean Architecture is about seperating your core entities (interfaces) and logic from the messy implementations 

I submit that using the purpose in the file names provides the seperation.