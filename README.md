grep-csv
===

This tool works like grep but it understands csv files.

Introduction
---

Here is how it works:

    $ wget http://www.ferc.gov/docs-filing/eqr/soft-tools/sample-csv/contract.txt
    
    # get all contracts where seller_company_name is The Electric Company and prints out contract_execution_date
    $ cat contract.txt | grep-csv  -f "seller_company_name" -v "The Electric Company" -p "contract_execution_date"

    # get contracts where contract_id is C78 and prints out seller_company_name,customer_company_name,customer_duns_number,contract_affiliate
    $ cat contract.txt | grep-csv  -f "contract_id" -v "C78" -p "seller_company_name,customer_company_name,customer_duns_number,contract_affiliate"

Install
--

Make sure you have go installed. Get it on golang

    $ go get github.com/htruong/grep-csv

That's all, folks!

Questions: htruong@tnhh.net