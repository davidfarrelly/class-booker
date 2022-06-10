# GoLang Class Booking
This project contains a simple class booking API, built in GoLang. 

# Release

- [v1.0.0](https://github.com/davidfarrelly/class-booker/releases/tag/v.1.0.0)

# Building

The following command can be used to build an executable for Windows:

```bash
set GOOS=windows
set GOARCH=amd64

cd src/
go build -o class-booker.exe
```

Or on Linux:

```bash
export GOOS=linux
export GOARCH=amd64

cd src/
go build -o class-booker
```

# API Usage

The application has an API to create classes (```/classes```) and an API to book a class for a member (```/bookings```)

Application supports standard REST operations ```GET, POST, PUT, DELETE```

## Classes

#### **GET all classes**
```http://localhost:8080/classes```

#### **GET specific class**
```http://localhost:8080/classes/{className}```

#### **POST a new class**
```http://localhost:8080/classes```
```
{
    "name": "yoga",
    "capacity": "10",
    "startDate": "2022-06-10",
    "endDate": "2022-07-10"
}
```

#### **PUT an updated class**
```http://localhost:8080/classes```
```
{
    "name": "yoga",
    "capacity": "20",
    "startDate": "2022-06-10",
    "endDate": "2022-07-10"
}
```

#### **DELETE a class**
```http://localhost:8080/classes/{className}```


## Bookings

#### **GET all bookings**
```http://localhost:8080/bookings```

#### **GET specific class**
```http://localhost:8080/bookings/{memberName}?class={className}```

#### **POST a new class**
```http://localhost:8080/bookings```
```
{
    "name": "david farrelly",
    "className": "yoga",
    "date": "2022-06-20"
}
```

#### **PUT an updated class**
```http://localhost:8080/bookings```
```
{
    "name": "david farrelly",
    "className": "yoga",
    "date": "2022-06-22"
}
```

#### **DELETE a class**
```http://localhost:8080/bookings/{memberName}?class={className}```

# Running

## Pre-requisites

- Golang installed on machine

## Running main.go

```bash
./main.go
```

## Running Binary

```bash
./class-booker.exe
```

# Testing

Unit tests can be run wih:

```bash
cd test/
go test -v
```

# Limitations & Potential Improvements

 - Could be more input validation, do not currently handle misformated date fields in class and booking objects
 - Actual DB such as mongodb could be used instead of slice (in-memory) [ISSUE-2](https://github.com/davidfarrelly/class-booker/issues/2)
 - Do not currently handle overbooked classes, could add a check on the capacity when creating a new booking [ISSUE-1](https://github.com/davidfarrelly/class-booker/issues/1)
 - Some logic could be moved into a seperate package, outside of the controllers.
 - Some positive and negative unit tests were added, but not every scenario is covered yet.