# Recipe

## API

### Requirement
- golang version 1.7 or latest
- sql-migrate (https://github.com/rubenv/sql-migrate)

### How to Use
    git clone https://github.com/mmuflih/recipe.git
    cd recipe/api
    cp env.json.example env.json
    cp example.dbconfig.yml dbconfig.yml
    
### How to Run
    - sql-migrate up
    - go run main.go

### Api Documentation

    https://documenter.getpostman.com/view/145345/U16gQnhn

----------------------------
## Frontend
    cd recipe/web
    npm install or yarn install
    yarn serve

----------------------------
## Video Demo
https://drive.google.com/drive/folders/1DnGn_kVTVtvPkhKt93jaZbBDkFWIpQ-9?usp=sharing