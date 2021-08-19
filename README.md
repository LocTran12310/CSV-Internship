# GolangAPI

### Ngày 19/08/2021
1. Tạo table Users: https://docs.google.com/spreadsheets/d/18QPaqFgQ9U_CWuVjtc8N00UGKvWLB0ib5_pDlR_8L3o/edit#gid=1251229191Preview
2. Code API Login.
3. Update lại các API GET, Add, Update profile sau khi login thành công mới có quyền truy cập.
4. Note:
    1. Mã hóa password: Tùy chọn cách mã hóa (https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/09.5.html)
    1. Quản session: dùng lib github.com/gorilla/sessions đề quản. Quản lý phiên bằng Cookie trả về cho client.


### Ngày 17/08/2021
Code API GET profiles với 2 params: numPage, numRow.

Kết quả trả về:
```
{
    "maxPage":
    "result": [
        {},
        {},
        ...
    ]
}
```

### Ngày 16/08/2021
Tái cấu trúc lại project.
1. Thao tác với database (select,update,delete) viết trong folder reponsitory -> profile_repository.go => return về data profile dùng model để hứng data
2. Tạo 1 folder Model viết các file model trên trong đó
3. Tạo 1 folder View viết các file model return view json trong đó.
4. Tạo 1 file server.go viết các router trong đó.
5. Kết hợp các folder với nhau. Để thành 1 project API.

```
.
├──database/
|   └──connect.go
├──model/
|   └──profile.go
├──repository/
|   └──profile_repository.go
├──view/
|   └──profile.go
├──main.go
├──server.go
├──...
```

### Ngày 13/08/2021
link cấu trúc table profiles tạo database bằng mysql (Dùng Xampp)
https://docs.google.com/spreadsheets/d/18QPaqFgQ9U_CWuVjtc8N00UGKvWLB0ib5_pDlR_8L3o/edit#gid=744585978Preview

- Code bằng Golang
- Code Get List profile
- Code API Add profile
- Code API Update profile
- Dùng postman để thao tác với API




<!-- Object mẫu
//Profile
{
    "id": 9,
    "employee_id": "999999",
    "name": "Tran Phuoc Loc",
    "email": "email@mail.com",
    "birthday": "2021-08-13 12:13:14",
    "position_id": 1,
    "department_id": 2,
    "status": 1,
    "address": "abca acb, dhc, Vietnam",
    "telephone": "0809123456",
    "mobile": "0123456789",
    "official_date": "2021-08-09",
    "probation_date": "2021-08-09",
    "gender": 1,
    "image": "",
    "del_flag": 0
}

//User
{
    "id": 1,
    "login_id": "user",
    "contract_type_id": 1,
    "profile_id": 10,
    "password": "123456",
    "del_flag": 0,
    "updated_time": "2021-08-18 10:20:49",
    "updated_user": "123",
    "created_time": "2021-08-12 10:20:49",
    "created_user": "123"
}
-->