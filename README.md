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

# Usage

The application has an API to create classes (```/classes```) and an API to book a class for a member (```/bookings```)

Application supports standard REST operations ```GET, POST, PUT, DELETE```

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

### API Objects

#### **`class.json`**
```json
{
    "name": "yoga",
    "capacity": "10",
    "startDate": "2022-06-10",
    "endDate": "2022-07-10"
}
```

#### **`booking.json`**
```json
{
    "name": "david farrelly",
    "className": "yoga",
    "date": "2022-06-20"
}

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