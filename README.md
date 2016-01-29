Goosees is a wrapper of [goose](https://bitbucket.org/liamstask/goose/) for applying migration to multiple databases.

# Config exmaple

```yaml
# "companyDB" group has one database.
companyDB:
  - driver: mysql
    open: root@tcp(127.0.0.1:3306)/company?parseTime=true

# "employeeDB" group has two databases.
employeeDB:
  - driver: mysql
    open: root@tcp(127.0.0.1:3306)/employee1?parseTime=true
  - driver: mysql
    open: root@tcp(127.0.0.1:3306)/employee2?parseTime=true

```

* `?parseTime=true` **MUST** be included in `open` string to avoid error at status command. See [goose issue #22](https://bitbucket.org/liamstask/goose/issues/22/scan-error-on-column-index-0-unsupported#comment-15419169)

* So far only `mysql` is available as a value of `driver`.

# Usage

    goosees <conf> <group> <subcommand> [subcommand options]

* `conf` is path to config file. For convenience, extension(`.yml`,`.yaml`) can be omitted.

## create

Migration file is created under a directory named group name.

    $ goosees dev companyDB create CreateCompanies sql
    goosees: companyDB/20160128222309_CreateCompanies.sql

## up

Migration is applied to all databases in a group.

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

    $ goosees dev companyDB status
    goosees: status for environment 'companyDB[0]'
        Applied At                  Migration
        =======================================
        Thu Jan 28 22:10:56 2016 -- 20160128220727_CreateCompanies.sql
        Thu Jan 28 22:13:44 2016 -- 20160128221155_AddCompanyAddress.sql

## dbversion

    $ goosees dev companyDB dbversion
    goosees: [companyDB:0] dbversion 20160128221155
    
    $ goosees dev employeeDB dbversion
    goosees: [employeeDB:0] dbversion 20160128221551
    goosees: [employeeDB:1] dbversion 20160128221551
