# Gin Gorm Firebase Minimum Template
## Usages
### Run app
```bash
$ make setup
```

### Test app
```bash
$ make test
```

### Migration
#### create migration
```bash
# ex
$ make new-migration FILE=create_users
```

#### apply migration
```
$ make migration
```

## About .env
You must need .env file to run app appropriately.
Before writing .env file, you must execute following commands in order to convert the json file of firebase service key into string because the app does not read service key as json but strings for some future development benefits.  
```bash
$ export FIREBASE_KEYFILE_JSON="$(< ./firebase-adminsdk.json)"
$ echo $FIREBASE_KEYFILE_JSON
```

Then, create .env file as below
```
ENV=development|staging|production
API_PORT=put the port whatever you want to use
DB_PORT=put the port whatever you want to use
FIREBASE_JSON=paste strings that echo shows
```