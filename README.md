### Uyga Vazifa: Go'da gRPC Gateway foydalanadigan kutubxona boshqaruv tizimi

#### Maqsad
`Go`, mikroservislar arxitekturasi va `gRPC` yordamida oddiy kutubxona boshqaruv tizimini yarating. Kitoblarni boshqarish va kreditlarni boshqarish (ijara) uchun mikroservislarni joriy qiling va ularni `gRPC Gateway` orqali namoyish qiling.

### Talablar
- **Book Service**: Kitoblar katalogini boshqarish uchun
- **Loan Service**: Kitob ijaraga berishni boshqarish uchun
- **API Gateway**: `HTTP` endpointlarini ochib beradi va so'rovlarni tegishli `gRPC` mikroservislariga yo'naltiradi.


**book.proto**
```proto
syntax = "proto3";

service BookService {
    rpc AddBook(AddBookRequest) returns (AddBookResponse);
    rpc GetBook(GetBookRequest) returns (Book);
}

message AddBookRequest {
    string title = 1;
    string author = 2;
}

message AddBookResponse {
    string book_id = 1;
}

message GetBookRequest {
    string book_id = 1;
}

message Book {
    string book_id = 1;
    string title = 2;
    string author = 3;
}

```

**loan.proto**
```proto
syntax = "proto3";

service LoanService {
    rpc IssueLoan(IssueLoanRequest) returns (IssueLoanResponse);
    rpc ReturnLoan(ReturnLoanRequest) returns (ReturnLoanResponse);
}

message IssueLoanRequest {
    string book_id = 1;
    string user_id = 2;
}

message IssueLoanResponse {
    string loan_id = 1;
}

message ReturnLoanRequest {
    string loan_id = 1;
}

message ReturnLoanResponse {
    bool success = 1;
}

```



   


