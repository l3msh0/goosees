Goosees is a wrapper of [goose](https://bitbucket.org/liamstask/goose/) to apply migration for multiple databases.

# Usage

    goosees <conf> <group> <subcommand> [subcommand options]


## create


    $ goosees dev companyDB create CreateCompanies sql
    goosees: companyDB/20160128222309_CreateCompanies.sql

## up

    $ goosees dev employeeDB up
    goose: migrating db environment 'employeeDB[0]', current version: 0, target: 20160128221551
    OK    20160128221551_CreateEmployees.sql
    goose: migrating db environment 'employeeDB[1]', current version: 0, target: 20160128221551
    OK    20160128221551_CreateEmployees.sql    $ goosees dev employeeDB up

## down

    $ goosees dev employeeDB down
    goose: migrating db environment 'employeeDB[0]', current version: 20160128221551, target: 0
    OK    20160128221551_CreateEmployees.sql
    goose: migrating db environment 'employeeDB[1]', current version: 20160128221551, target: 0
    OK    20160128221551_CreateEmployees.sql

## status

## dbversion

    $ goosees dev companyDB dbversion
    goosees: [companyDB:0] dbversion 20160128221155
    
    $ goosees dev employeeDB dbversion
    goosees: [employeeDB:0] dbversion 20160128221551
    goosees: [employeeDB:1] dbversion 20160128221551
