# Go BINLIST

Go BINLIST is a Go library for retrieving credit card information based on the bank identification number (BIN). The library uses the [binlist.net](https://binlist.net/) public API to retrieve credit card information.



### Installation

```go
go get github.com/alomia/go-binlist/card
```



### Usage

You can obtain information about a credit card by using the BIN number as follows:

```go
cardInfo, err := GetInfo("123456")
if err != nil {
    // handle error
}

fmt.Println(cardInfo.Type) // print credit card type

```



### CardInfo struct

The CardInfo struct contains the following fields:

- *Number*: A *number* struct containing the length of the card number and whether or not it passes the Luhn algorithm.
- *Scheme*: The payment scheme of the card (e.g. Visa, Mastercard).
- *Type*: The type of the card (e.g. debit, credit).
- *Brand*: The brand of the card (e.g. Chase, Wells Fargo).
- *Prepaid*: A boolean indicating if the card is prepaid.
- *Country*: A *country* struct containing information about the country where the card was issued.
- *Bank*: A *bank* struct containing information about the bank that issued the card.



### GetInfo function

The *GetInfo* function retrieves card information from a bin number.

```go
func GetInfo(bin string) (*CardInfo, error)
```

If the card number is not valid, the function returns an error.



### isValidCardNumber function

The *isValidCardNumber* function checks if a card number is valid.

```go
func isValidCardNumber(bin string) bool
```

The function returns true if the card number is valid, and false otherwise.



### Contribute

If you wish to contribute to the library, fork the repository, make your changes and submit a pull request. Be sure to follow the project guidelines and include unit tests for any changes you make.
