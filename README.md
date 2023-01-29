# Costs

A simple Golang CLI tool for calculating shared expenses. It takes a list of incomes and a total amount of expenses and calculates how much each person should pay.

It calculates two different ratios, one based on the per-person income and one calculates a 50/50 ratio.

### Installation

    $ go get github.com/bad33ndj3/costs

### Usage

    $ costs -person1=2000 -person2=2500 -expenses=1200

### Example output
```shell
TO RATIO
┏━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━┓
┃     # ┃ INCOME ┃ INC. RATIO ┃ PAYS   ┃
┣━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━╋━━━━━━━━┫
┃     1 ┃ 2000   ┃ 44.44      ┃ 666.67 ┃
┃     2 ┃ 2500   ┃ 55.56      ┃ 833.33 ┃
┣━━━━━━━╋━━━━━━━━╋━━━━━━━━━━━━╋━━━━━━━━┫
┃ TOTAL ┃ 4500   ┃            ┃ 1500   ┃
┗━━━━━━━┻━━━━━━━━┻━━━━━━━━━━━━┻━━━━━━━━┛
50/50
┏━━━━━━━┳━━━━━━━━┳━━━━━━┓
┃     # ┃ INCOME ┃ PAYS ┃
┣━━━━━━━╋━━━━━━━━╋━━━━━━┫
┃     1 ┃ 2000   ┃ 750  ┃
┃     2 ┃ 2500   ┃ 750  ┃
┣━━━━━━━╋━━━━━━━━╋━━━━━━┫
┃ TOTAL ┃ 4500   ┃ 1500 ┃
┗━━━━━━━┻━━━━━━━━┻━━━━━━┛


```
