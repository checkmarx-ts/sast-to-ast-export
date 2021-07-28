# How to export data

1. Create an admin user in SAST
2. Assign credentials in environment:
```
   export SAST_EXPORT_USERNAME='myusername'
   export SAST_EXPORT_PASSWORD='mypassw0rd'
```
3. Run export
```
    ./sast-export http://127.0.0.1 -o .
```