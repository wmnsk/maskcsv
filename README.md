# mask_csv.go

Masks the last X letters in specified fields in a CSV.

## Example

```shell-session
# cat test.csv
name,age,telephone,password
Ichiro,42,818012345678,ichiro1234
Matsui,42,819012345678,hideki99
Darvish,30,818098765432,yu19860816
#
# bin/mask_csv_linux_amd64 -i test.csv -o testx.csv -f "telephone,password" -m "x" -l 4
Masking successfully finished!
#
# cat testx.csv
name,age,telephone,password
Ichiro,42,81801234xxxx,ichiroxxxx
Matsui,42,81901234xxxx,hidexxxx
Darvish,30,81809876xxxx,yu1986xxxx
```

## Usage

```shell-session
Usage of mask_csv:
  -d string
        Delimiter of input file. (default ",")
  -f string
        Path to the original CSV file to be masked. (default "no,header,val")
  -i string
        File to be masked.
  -l int
        Number of the letters to mask. (default 2)
  -m string
        Character to be used as the mask. (default "X")
  -o string
        File to be masked.
  -s string
        Delimiter of output file. (default ",")
```
