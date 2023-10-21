# Ecom Layered Architecture:

## Problem to solve:
* Monolithic Code becomes complicated and complex since everything is packed into one
* Adding new changes becomes difficult since the code is interconnected. New Changes can break existing working code
even if the change is minor.
* Going through the code and maintaining it will become difficult since everything is highly coupled and to change/update one minor thing, we have to
understand whole codebase

## Requirements

### Entities/Domains
* Products
* Users
* Orders


### Functional Requirements
* Show products to users
* Allow users to login and register
* Allow users to add product to cart
* Allow users to checkout cart and place order
* Allow users to see old orders

### Non Functional Requirements
* Support small number of concurrent users
* Code should be maintainable
* Available



## Solution:
### Architecture 
Layered Architecture

### Principles
* KISS
* YAGNI
* DRY
* Separation of Concern
* SOLID

## Explanation about the project
There are a lot of ways to organise a project


