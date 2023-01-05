# Syndio Backend App


## Software requirements
Please make sure that these softwares are installed on your machine.
- Golang should be installed on your system
- Sqlite3 needs to be installed on your system

Reference:

[How to install SQLite](https://www.servermania.com/kb/articles/install-sqlite/)

[How to install Golang](https://go.dev/doc/install)

<br>

# Setup Instruction

* Set the environment variable "PORT"

    you can set the environment variable by creating .env file.
    
    On MacOS or Linux, you can run the following command.

    ```cp .env.example  .env```

    The application will load .env file and it will set the environment "PORT"
    You can change the PORT number by editing the .env file.

<br>

* Start the application
    
    please go inside this project folder and run the following command:

    ```go run .```

<br>

* Access to the endpoint

    ```curl localhost:$PORT/employees```
    
    OR

    ```curl http://localhost:$PORT/employees```
    
<br>



# Result

The application loads data from SQLite database and it should show the data like this:

```
[
    {
        "gender": "male",
        "id": 1
    },
    {
        "gender": "male",
        "id": 2
    },
    {
        "gender": "male",
        "id": 3
    },
    {
        "gender": "female",
        "id": 4
    },
    {
        "gender": "female",
        "id": 5
    },
    {
        "gender": "female",
        "id": 6
    }
]
```