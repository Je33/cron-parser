# Cron Schedule Parser

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/Je33/currency-converter)
[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/Je33/currency-converter/test.yml?style=flat-square)](https://github.com/Je33/currency-converter/actions/workflows/test.yml)

This tool provide the parser of cron schedule common format

## Usage

Just run bin file with 3 arguments: amount (float), base currency (string) and target currency (string)

 ```sh
 ./parser "*/15 1-12,15,20-23 1-4 * * /bin/command -with -parameters"
 ```

And you will get the result:

```sh
minutes:       0 15 30 45
hours:         1 2 3 4 5 6 7 8 9 10 11 12 15 20 21 22 23
days of month: 1 2 3 4
months:        1 2 3 4 5 6 7 8 9 10 11 12
days of week:  0 1 2 3 4 5 6
command:       /bin/command -with -parameters

```


## Common cron schedule format

```
* * * * * /command
^ ^ ^ ^ ^     ^
| | | | |     |
1 2 3 4 5     6
```

| # | Field name   | Mandatory? | Allowed values | Special characters   |
|---|--------------|------------|----------------|----------------------|
| 1 | Minute       | Yes        | 0-59           | * / , -              |
| 2 | Hour         | Yes        | 0-23           | * / , -              |
| 3 | Day of month | Yes        | 1-31           | * / , -              |
| 4 | Month        | Yes        | 1-12           | * / , -              |
| 5 | Day of week  | Yes        | 0-6            | * / , -              |
| 6 | Command      | Yes        | Script path    | Path safe characters |

