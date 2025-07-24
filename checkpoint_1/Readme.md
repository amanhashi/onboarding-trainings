As part of consolidating my Go learnings, I built a small command-line Contact Manager in Go. The project demonstrates my understanding of core Go concepts including:

Structs & Custom Types for contact modeling

Functions, Arrays, Slices, and Maps for operations

Interfaces and Generics for extensibility

Use of Pointers for modifying contact data

Package and modular code design

Usage of standard Go syntax and CLI input/output

This helped reinforce both foundational and advanced topics in a practical context.

Mini Project Idea: Go Contact Manager (CLI)
A basic command-line Contact Manager that allows you to:

Add new contacts

List all contacts

Search for a contact by name

This small project includes:

Structs for contact data

Functions for handling logic

Maps & slices for data storage

Pointers to modify data

Interfaces & generics

Packages (split into files/folders if needed)

Arrays (for static demo)

Mutex

Go basics & syntax usage


-----------------------------------------------------------------


Terraform 

Its an open source IAC tool that allows to defined ,provision and manage infra across multiple cloud platforms .

Technical Concepts

1- Terraform Plugins and Providers
Providers are terraform plugins 
that act as middleman between Terraform and cloud api's(azure ,azws etc).
Providers will define resources like EC2,VPC and DB's or data sources

2- How terraform works with Plugins?
Terraform install providers from the TerraformRegistry ,also there are private registry from where it can import

3- Terraform State and Provider Configuration 
It stores a state file that records the actual infra it manages,this is how terraform know what it created
We can define multiple configurations
Provider aliases support multple provider instances
We can lock and upgrade provider versions as well

4- Terraform's value to Practitioners
Consistency - Define IAAC and use the same config across all environments
Multi-Cloud Ready - Manage AWS,GCP,AZURE ,with a single code language
Automation- Apply changes through pipelines without tension of manual provisioning 
Auditability - Track who changed what when  and why version control system 